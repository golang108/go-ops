package script

import (
	"context"
	"go-ops/internal/model"
	"go-ops/pkg/agent/action"
	"go-ops/pkg/agent/cmdrunner"
	"go-ops/pkg/agent/script/cmd"
	"path"
	"path/filepath"
	"strings"
)

// 压缩包脚本，脚本放在一个压缩包内
type PackScript struct {
	GenericScript
	cmd     string
	input   string
	fileMd5 string
}

func NewPackScript(
	runner cmdrunner.CmdRunner,
	jobid string,
	path string,
	cmd string,
	content string,
	env map[string]string,
	timeout int,
	input string,
	user string,
	args []string,
	filemd5 string,
) PackScript {

	if cmd == "" {
		cmd = Cmder
	}

	s := PackScript{cmd: cmd, input: input, fileMd5: filemd5}
	s.GenericScript.runner = runner
	s.GenericScript.path = path
	s.GenericScript.content = content
	s.GenericScript.env = env
	s.GenericScript.jobid = jobid
	s.GenericScript.timeout = timeout
	s.GenericScript.user = user
	s.GenericScript.args = args

	return s
}

func (s PackScript) Run() (r model.ResCmd) {

	if s.path == "" {
		s.path = ScriptPath
	}
	runpath := path.Join(s.path, s.jobid)
	err := s.ensureContainingDir(runpath)

	if err != nil {
		return s.getResCmd(nil, err)
	}

	filename := filepath.Join(runpath, filepath.Base(s.content))

	ctx := context.Background()

	err = action.Download(ctx, filename, s.content)

	if err != nil {
		return s.getResCmd(nil, err)
	}

	err = action.CheckFileMd5(filename, s.fileMd5)
	if err != nil {
		return s.getResCmd(nil, err)
	}

	err = action.Untar(filename)
	if err != nil {
		return s.getResCmd(nil, err)
	}

	cmdstr, args := getCmdArgs(s.cmd)

	command := cmd.BuildCommand(cmdstr)

	command.Args = append(command.Args, args...)
	command.Args = append(command.Args, runpath)
	command.Args = append(command.Args, s.args...)
	command.Timeout = s.timeout
	command.User = s.user
	command.WorkingDir = s.path

	if s.input != "" {
		command.Stdin = strings.NewReader(s.input)
	}

	for key, val := range s.env {
		command.Env[key] = val
	}

	res, err := s.runner.RunCommand(s.jobid, command)
	return s.getResCmd(res, err)
}
