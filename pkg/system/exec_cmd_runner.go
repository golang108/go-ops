package system

import (
	"context"
	"go-ops/pkg/errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

type execCmdRunner struct {
	ctx context.Context
}

func NewExecCmdRunner(ctx context.Context) CmdRunner {
	return execCmdRunner{ctx: ctx}
}

func (r execCmdRunner) RunComplexCommand(cmd Command) (string, string, int, ExitWay, error) {
	execCmd, err := r.buildComplexCommand(cmd)
	if err != nil {
		return "", "", -1, 0, err
	}
	process := NewExecProcess(execCmd, cmd.KeepAttached, cmd.Timeout)

	err = process.Start()
	if err != nil {
		return "", "", -1, 0, err
	}

	result := process.Wait(r.ctx)

	return result.Stdout, result.Stderr, result.ExitStatus, result.ExitWay, result.Error
}

func (r execCmdRunner) RunComplexCommandAsync(cmd Command) (Process, error) {
	execCmd, err := r.buildComplexCommand(cmd)
	if err != nil {
		return nil, err
	}
	process := NewExecProcess(execCmd, cmd.KeepAttached, cmd.Timeout)

	err = process.Start()
	if err != nil {
		return nil, err
	}

	return process, nil
}

func (r execCmdRunner) RunCommand(cmdName string, args ...string) (string, string, int, ExitWay, error) {
	return r.RunComplexCommand(Command{Name: cmdName, Args: args})
}

func (r execCmdRunner) RunCommandWithInput(input, cmdName string, args ...string) (string, string, int, ExitWay, error) {
	cmd := Command{
		Name:  cmdName,
		Args:  args,
		Stdin: strings.NewReader(input),
	}
	return r.RunComplexCommand(cmd)
}

func (r execCmdRunner) CommandExists(cmdName string) bool {
	_, err := exec.LookPath(cmdName)
	return err == nil
}

func (r execCmdRunner) buildComplexCommand(cmd Command) (execCmd *exec.Cmd, err error) {
	execCmd = newExecCmd(cmd.Name, cmd.Args...)

	if cmd.Stdin != nil {
		execCmd.Stdin = cmd.Stdin
	}

	if cmd.Stdout != nil {
		execCmd.Stdout = cmd.Stdout
	}

	if cmd.Stderr != nil {
		execCmd.Stderr = cmd.Stderr
	}

	if cmd.User != "" {
		// 检测用户是否存在
		user, err := user.Lookup(cmd.User)
		if err != nil {
			return nil, errors.WrapErrorf(err, "invalid user %s", cmd.User)
		}

		// set process attr
		// 获取用户 id
		uid, err := strconv.ParseUint(user.Uid, 10, 32)
		if err != nil {
			return nil, err
		}
		// 获取用户组 id
		gid, err := strconv.ParseUint(user.Gid, 10, 32)
		if err != nil {
			return nil, err
		}
		// 修复execCmd.SysProcAttr为nil的bug
		var attr *syscall.SysProcAttr
		if execCmd.SysProcAttr == nil {
			attr = new(syscall.SysProcAttr)
		} else {
			attr = execCmd.SysProcAttr
		}

		//设置进程执行用户
		attr.Credential = &syscall.Credential{
			Uid: uint32(uid),
			Gid: uint32(gid),
		}

		execCmd.SysProcAttr = attr
	}

	execCmd.Dir = cmd.WorkingDir

	var env []string
	if !cmd.UseIsolatedEnv {
		env = os.Environ()
	}
	if cmd.UseIsolatedEnv && runtime.GOOS == "windows" {
		panic("UseIsolatedEnv is not supported on Windows")
	}

	execCmd.Env = mergeEnv(env, cmd.Env)

	return execCmd, nil
}
