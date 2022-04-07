package task

type asyncTaskService struct {
	currentTasks map[string]Task
	taskChan     chan Task
	taskSem      chan func()
}

func NewAsyncTaskService() (service Service) {
	s := asyncTaskService{
		currentTasks: make(map[string]Task),
		taskChan:     make(chan Task),
		taskSem:      make(chan func()),
	}

	go s.processTasks()
	go s.processSemFuncs()

	return s
}

func (service asyncTaskService) CreateTask(
	id string,
	taskFunc Func,
	cancelFunc CancelFunc,
	endFunc EndFunc,
) Task {
	return Task{
		ID:         id,
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

		// Nil to prevent to memory leaks in case these are closures.
		task.Func = nil
		task.CancelFunc = nil
		task.EndFunc = nil

		service.taskSem <- func() {
			service.currentTasks[task.ID] = task
		}
	}
}
