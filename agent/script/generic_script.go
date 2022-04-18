package script

import (
	"os"
	"path/filepath"

	"github.com/charlievieth/fs"

	"go-ops/agent/cmdrunner"
	"go-ops/agent/script/cmd"
	"go-ops/internal/model"
)

const (
	fileOpenFlag int         = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	fileOpenPerm os.FileMode = os.FileMode(0777)
)

type ExecType int

const (
	ExecCmd     ExecType = iota // 直接命令执行
	ExecContent                 // 需要解释器  内容保存在一个文件脚本里面执行
	ExecScript                  // 需要解释器 脚本就在本地
	ExecURL                     // 脚本从云端下载
)

type GenericScript struct {
	runner  cmdrunner.CmdRunner
	path    string
	content string
	args    []string
	jobid   string
	env     map[string]string
	timeout int
	user    string
}

func NewScript(
	runner cmdrunner.CmdRunner,
	jobid string,
	path string,
	content string,
	env map[string]string,
	timeout int,
	user string,
	args []string,
) GenericScript {
	return GenericScript{
		runner:  runner,
		path:    path,
		content: content,
		env:     env,
		jobid:   jobid,
		timeout: timeout,
		user:    user,
		args:    args,
	}
}

func (s GenericScript) Path() string { return s.path }

func (s GenericScript) Run() (r model.ResCmd) {
	cmdstr, args := getCmdArgs(s.content)
	command := cmd.BuildCommand(cmdstr)
	for key, val := range s.env {
		command.Env[key] = val
	}

	command.Args = append(command.Args, args...)
	command.Timeout = s.timeout
	command.User = s.user
	command.WorkingDir = s.path

	res, err := s.runner.RunCommand(s.jobid, command)
	return s.getResCmd(res, err)
}

func (s GenericScript) getResCmd(res *cmdrunner.CmdResult, err error) (r model.ResCmd) {
	if err != nil {
		r.Code = model.CodeNotRun
		r.Err = err.Error()
		r.ExitCode = -1
		return
	}
	r = model.ResCmd{
		Stdout:   string(res.Stdout),
		Stderr:   string(res.Stderr),
		ExitCode: res.ExitStatus,
	}

	if res.ResCode != "" {
		r.Code = res.ResCode
		return
	}

	if r.ExitCode == 0 {
		r.Code = model.CodeSuccess
	} else {
		r.Code = model.CodeFailed
	}
	return
}

func (s GenericScript) ensureContainingDir(fullLogFilename string) error {
	dir, _ := filepath.Split(fullLogFilename)
	return fs.MkdirAll(dir, os.FileMode(0750))
}
