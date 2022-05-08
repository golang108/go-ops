package task

type Service interface {
	CreateTask(string, interface{}, Func, CancelFunc, EndFunc) Task
	StartTask(Task)
	FindTaskWithID(string) (Task, bool)
}
