package task

type Info struct {
	TaskID  string
	Method  string
	Payload []byte
}

type ManagerProvider interface {
	NewManager(string) Manager
}

type Manager interface {
	GetInfos() ([]Info, error)
	AddInfo(taskInfo Info) error
	RemoveInfo(taskID string) error
}
