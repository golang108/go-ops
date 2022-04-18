// +build !windows

package system

import (
	"strings"
	"syscall"

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
	//p.logger.Debug(execProcessLogTag, "Running command '%s'", cmdString)

	if !p.keepAttached {
		p.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}

	err := p.cmd.Start()
	if err != nil {
		return errors.WrapErrorf(err, "Starting command '%s'", cmdString)
	}

	if !p.keepAttached {
		p.pgid = p.cmd.Process.Pid
	} else {
		p.pgid, err = syscall.Getpgid(p.pid)
		if err != nil {

			//p.logger.Error(execProcessLogTag, "Failed to retrieve pgid for command '%s'", cmdString)
			p.pgid = -1
		}
	}

	return nil
}
