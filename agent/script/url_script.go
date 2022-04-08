package script

import (
	"io"
	"net/http"
	"os"
	"osp/agent/cmdrunner"
	"osp/agent/script/cmd"
	"osp/internal/model"
	"path"
)

type UrlScript struct {
	GenericScript
	cmd string
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
) UrlScript {

	if cmd == "" {
		cmd = Cmder
	}

	s := UrlScript{cmd: cmd}
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
	command.Args = append(command.Args, s.content)
	command.Args = append(command.Args, s.args...)
	command.Timeout = s.timeout
	command.User = s.user

	for key, val := range s.env {
		command.Env[key] = val
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
