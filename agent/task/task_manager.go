package task

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"osp/pkg/errors"
	"path"
)

type taskManagerProvider struct{}

func NewManagerProvider() ManagerProvider {
	return taskManagerProvider{}
}

func (provider taskManagerProvider) NewManager(
	dir string,
) Manager {
	return NewManager(path.Join(dir, "tasks.json"))
}

type taskerManager struct {
	fsSem     chan func()
	tasksPath string

	taskInfos map[string]Info
}

func NewManager(tasksPath string) Manager {
	m := &taskerManager{

		fsSem:     make(chan func()),
		tasksPath: tasksPath,
		taskInfos: make(map[string]Info),
	}

	go m.processFsFuncs()

	return m
}

func (m *taskerManager) GetInfos() ([]Info, error) {
	taskInfosChan := make(chan map[string]Info)
	errCh := make(chan error)

	m.fsSem <- func() {
		taskInfos, err := m.readInfos()
		m.taskInfos = taskInfos
		taskInfosChan <- taskInfos
		errCh <- err
	}

	taskInfos := <-taskInfosChan
	err := <-errCh

	if err != nil {
		return nil, err
	}

	var r []Info
	for _, taskInfo := range taskInfos {
		r = append(r, taskInfo)
	}

	return r, nil
}

func (m *taskerManager) AddInfo(taskInfo Info) error {
	errCh := make(chan error)

	m.fsSem <- func() {
		m.taskInfos[taskInfo.TaskID] = taskInfo
		err := m.writeInfos(m.taskInfos)
		errCh <- err
	}
	return <-errCh
}

func (m *taskerManager) RemoveInfo(taskID string) error {
	errCh := make(chan error)

	m.fsSem <- func() {
		delete(m.taskInfos, taskID)
		err := m.writeInfos(m.taskInfos)
		errCh <- err
	}
	return <-errCh
}

func (m *taskerManager) processFsFuncs() {

	for {
		do := <-m.fsSem
		do()
	}
}

func (m *taskerManager) readInfos() (map[string]Info, error) {
	taskInfos := make(map[string]Info)

	_, err := os.Stat(m.tasksPath)
	if err != nil {
		return taskInfos, nil
	}
	if os.IsNotExist(err) {
		return taskInfos, nil
	}

	tasksJSON, err := ioutil.ReadFile(m.tasksPath)
	if err != nil {
		return nil, errors.WrapError(err, "Reading tasks json")
	}

	err = json.Unmarshal(tasksJSON, &taskInfos)
	if err != nil {
		return nil, errors.WrapError(err, "Unmarshaling tasks json")
	}

	return taskInfos, nil
}

func (m *taskerManager) writeInfos(taskInfos map[string]Info) error {
	newTasksJSON, err := json.Marshal(taskInfos)
	if err != nil {
		return errors.WrapError(err, "Marshalling tasks json")
	}

	err = ioutil.WriteFile(m.tasksPath, newTasksJSON, fs.FileMode(os.O_WRONLY|os.O_TRUNC))
	if err != nil {
		return errors.WrapError(err, "Writing tasks json")
	}

	return nil
}
