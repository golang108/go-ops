package system

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"osp/pkg/errors"
)

const (
	execProcessLogTag = "Cmd Runner"
)

type execProcess struct {
	cmd          *exec.Cmd
	stdoutWriter *bytes.Buffer
	stderrWriter *bytes.Buffer
	keepAttached bool
	pid          int
	pgid         int
	//logger       boshlog.Logger
	waitCh  chan Result
	timeout int
}

func NewExecProcess(cmd *exec.Cmd, keepAttached bool, timeout int) *execProcess {
	return &execProcess{
		cmd:          cmd,
		stdoutWriter: bytes.NewBufferString(""),
		stderrWriter: bytes.NewBufferString(""),
		keepAttached: keepAttached,
		timeout:      timeout,
	}
}

func (p *execProcess) Wait(ctx context.Context) (res Result) {
	if p.waitCh != nil {
		panic("Wait() must be called only once")
	}

	if p.timeout == 0 {
		p.timeout = 3600
	}

	done := make(chan error)
	go func() {
		done <- p.cmd.Wait()
	}()

	after := time.After(time.Duration(p.timeout) * time.Second)
	select {
	case <-after:
		syscall.Kill(-p.cmd.Process.Pid, syscall.SIGKILL)
		res = p.getResult()
		res.ExitWay = EXITWAY_TIMEOUT
		return res

	case <-ctx.Done():
		syscall.Kill(-p.cmd.Process.Pid, syscall.SIGKILL)
		res = p.getResult()
		res.ExitWay = EXITWAY_TIMEOUT
		return res
	case err := <-done:
		res = p.getResult()
		if err != nil {
			cmdString := strings.Join(p.cmd.Args, " ")
			err = errors.WrapComplexError(err, NewExecError(cmdString, res.Stdout, res.Stderr))
			res.Error = err
		}
		return res
	}

	return
}

func (p *execProcess) getResult() Result {
	stdout := string(p.stdoutWriter.Bytes())

	stderr := string(p.stderrWriter.Bytes())

	exitStatus := -1

	if p.cmd != nil && p.cmd.ProcessState != nil {
		waitStatus := p.cmd.ProcessState.Sys().(syscall.WaitStatus)

		if waitStatus.Exited() {
			exitStatus = waitStatus.ExitStatus()
		} else if waitStatus.Signaled() {
			exitStatus = 128 + int(waitStatus.Signal())
		}
	}

	return Result{
		Stdout:     stdout,
		Stderr:     stderr,
		ExitStatus: exitStatus,
	}
}
