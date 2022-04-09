package task

import (
	"github.com/panjf2000/ants/v2"
)

type asyncTaskService struct {
	currentTasks map[string]Task
	taskChan     chan Task
	taskSem      chan func()
	taskPool     *ants.PoolWithFunc
}

func NewAsyncTaskService() (service Service) {
	s := asyncTaskService{
		currentTasks: make(map[string]Task),
		taskChan:     make(chan Task),
		taskSem:      make(chan func()),
	}

	p, _ := ants.NewPoolWithFunc(10, s.execTask)

	s.taskPool = p
	go s.processTasks()
	go s.processSemFuncs()

	return s
}

func (service asyncTaskService) CreateTask(
	id string,
	req interface{},
	taskFunc Func,
	cancelFunc CancelFunc,
	endFunc EndFunc,
) Task {
	return Task{
		ID:         id,
		Req:        req,
		State:      StateRunning,
		Func:       taskFunc,
		CancelFunc: cancelFunc,
		EndFunc:    endFunc,
	}
}

func (service asyncTaskService) StartTask(task Task) {
	taskChan := make(chan Task)

	service.taskSem <- func() {
		service.currentTasks[task.ID] = task
		taskChan <- task
	}

	recordedTask := <-taskChan
	service.taskChan <- recordedTask
}

func (service asyncTaskService) FindTaskWithID(id string) (Task, bool) {
	taskChan := make(chan Task)
	foundChan := make(chan bool)

	service.taskSem <- func() {
		task, found := service.currentTasks[id]
		taskChan <- task
		foundChan <- found
	}

	return <-taskChan, <-foundChan
}

func (service asyncTaskService) processSemFuncs() {

	for {
		do := <-service.taskSem
		do()
	}
}

func (service asyncTaskService) processTasks() {

	for {
		task := <-service.taskChan
		service.taskPool.Invoke(task)

	}
}

func (service asyncTaskService) execTask(val interface{}) {
	task := val.(Task)
	value, err := task.Func()
	if err != nil {
		task.Error = err
		task.State = StateFailed
	} else {
		task.Value = value
		task.State = StateDone
	}

	if task.EndFunc != nil {
		task.EndFunc(task)
	}

	task.Func = nil
	task.CancelFunc = nil
	task.EndFunc = nil

	service.taskSem <- func() {
		service.currentTasks[task.ID] = task
	}

}
