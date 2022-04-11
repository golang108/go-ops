package system

import (
	"fmt"
)

const (
	execErrorMsgFmt = "Running command: '%s', stdout: '%s', stderr: '%s'"
)

type ExecError struct {
	Command string
	StdOut  string
	StdErr  string
}

func NewExecError(cmd, stdout, stderr string) ExecError {
	return ExecError{
		Command: cmd,
		StdOut:  stdout,
		StdErr:  stderr,
	}
}

func (e ExecError) Error() string {
	return fmt.Sprintf(execErrorMsgFmt, e.Command, e.StdOut, e.StdErr)
}
