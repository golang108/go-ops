package system

import (
	"context"
	"io"
)

type Command struct {
	User           string
	Name           string
	Args           []string
	Env            map[string]string
	UseIsolatedEnv bool

	WorkingDir string

	// 只有在linux生效
	KeepAttached bool

	Timeout int

	Stdin io.Reader

	Stdout io.Writer
	Stderr io.Writer
}

type Process interface {
	Wait(context.Context) Result
}

type ExitWay int

const (
	EXITWAY_TIMEOUT ExitWay = 1
	EXITWAY_CANCEL  ExitWay = 2
)

type Result struct {
	Stdout string
	Stderr string

	ExitStatus int
	Error      error
	ExitWay    ExitWay
}

type CmdRunner interface {
	RunComplexCommand(cmd Command) (stdout, stderr string, exitStatus int, exitWay ExitWay, err error)

	RunComplexCommandAsync(cmd Command) (Process, error)

	RunCommand(cmdName string, args ...string) (stdout, stderr string, exitStatus int, exitWay ExitWay, err error)

	RunCommandWithInput(input, cmdName string, args ...string) (stdout, stderr string, exitStatus int, exitWay ExitWay, err error)

	CommandExists(cmdName string) (exists bool)
}
