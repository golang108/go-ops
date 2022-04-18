package system

import (
	"strings"
	"time"

	"go-ops/pkg/errors"
)

func (p *execProcess) Start() error {
	if p.cmd.Stdout == nil {
		p.cmd.Stdout = p.stdoutWriter
	}
	if p.cmd.Stderr == nil {
		p.cmd.Stderr = p.stderrWriter
	}
	cmdString := strings.Join(p.cmd.Args, " ")
	//p.logger.Debug(execProcessLogTag, "Running command: %s", cmdString)

	err := p.cmd.Start()
	if err != nil {
		return errors.WrapErrorf(err, "Starting command %s", cmdString)
	}

	p.pid = p.cmd.Process.Pid
	return nil
}
