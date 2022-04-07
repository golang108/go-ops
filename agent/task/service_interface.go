package task

type Service interface {
	CreateTask(string, Func, CancelFunc, EndFunc) Task
	StartTask(Task)
	FindTaskWithID(string) (Task, bool)
}
