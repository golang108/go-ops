package script

import (
	"go-ops/internal/model"
	"go-ops/pkg/agent/cmdrunner"
	"go-ops/pkg/agent/script/cmd"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type UrlScript struct {
	GenericScript
	cmd   string
	input string
}

func NewUrlScript(
	runner cmdrunner.CmdRunner,
	jobid string,
	path string,
	cmd string,
	content string,
	env map[string]string,
	timeout int,
	user string,
	args []string,
	input string,
) UrlScript {

	if cmd == "" {
		cmd = Cmder
	}

	s := UrlScript{cmd: cmd, input: input}
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

func (s UrlScript) Run() (r model.ResCmd) {

	err := s.downloadFile()
	if err != nil {
		return s.getResCmd(nil, err)
	}
	cmdstr, args := getCmdArgs(s.cmd)
	command := cmd.BuildCommand(cmdstr)
	command.Args = append(command.Args, args...)

	filename := path.Base(s.content)

	command.Args = append(command.Args, filename)
	command.Args = append(command.Args, s.args...)
	command.Timeout = s.timeout
	command.User = s.user
	command.WorkingDir = s.path

	for key, val := range s.env {
		command.Env[key] = val
	}

	if s.input != "" {
		command.Stdin = strings.NewReader(s.input)
	}

	res, err := s.runner.RunCommand(s.jobid, command)
	return s.getResCmd(res, err)
}

func (s UrlScript) downloadFile() (err error) {
	savePath := s.path
	if savePath == ScriptPath {
		savePath = path.Join(s.path, s.jobid)
	}

	filename := path.Base(s.content)

	savePath = path.Join(savePath, filename)

	err = s.ensureContainingDir(savePath)
	if err != nil {
		return
	}

	f, err := os.OpenFile(savePath, fileOpenFlag, fileOpenPerm)
	if err != nil {
		return
	}
	defer f.Close()

	res, err := http.Get(s.content)
	if err != nil {
		return
	}
	defer res.Body.Close()

	_, err = io.Copy(f, res.Body)
	return
}
