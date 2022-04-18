package script

import (
	"go-ops/agent/cmdrunner"
	"go-ops/agent/script/cmd"
	"go-ops/internal/model"
	"strings"
)

type NameScript struct {
	GenericScript
	cmd   string
	input string
}

func NewNameScript(
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
) NameScript {

	if cmd == "" {
		cmd = Cmder
	}

	s := NameScript{cmd: cmd, input: input}
	s.GenericScript.runner = runner
	s.GenericScript.path = path
	s.GenericScript.content = content
	s.GenericScript.env = env
	s.GenericScript.jobid = jobid
	s.GenericScript.timeout = timeout
	s.GenericScript.args = args
	return s
}

func (s NameScript) Run() (r model.ResCmd) {

	cmdstr, args := getCmdArgs(s.cmd)
	command := cmd.BuildCommand(cmdstr)
	command.Args = append(command.Args, args...)
	command.Args = append(command.Args, s.content)
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
