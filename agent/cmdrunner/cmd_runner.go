package cmdrunner

import (
	ospsys "osp/pkg/system"
)

type CmdResult struct {
	Stdout []byte
	Stderr []byte

	ExitStatus int
	IsKilled   bool
}

type CmdRunner interface {
	RunCommand(jobId string, cmd ospsys.Command) (*CmdResult, error)
}
