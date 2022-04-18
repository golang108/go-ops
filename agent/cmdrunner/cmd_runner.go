package cmdrunner

import (
	"go-ops/internal/model"
	ospsys "go-ops/pkg/system"
)

type CmdResult struct {
	Stdout []byte
	Stderr []byte

	ExitStatus int
	ResCode    model.ResCode
}

type CmdRunner interface {
	RunCommand(jobId string, cmd ospsys.Command) (*CmdResult, error)
}
