package action

import (
	"context"
	"path/filepath"

	"go-ops/internal/model"
	"go-ops/pkg/agent/script/cmd"
	"go-ops/pkg/errors"
	"go-ops/pkg/system"

	"github.com/toolkits/file"
)

func control(workdir, args string, timeout int) (stdout string, err error) {
	controlScript := getControlName()
	cmdRunner := system.NewExecCmdRunner(context.Background())

	command := cmd.BuildCommand("." + string(filepath.Separator) + controlScript)
	command.Args = append(command.Args, args)

	command.Timeout = timeout
	command.WorkingDir = workdir

	stdout, stderr, exitCode, _, err := cmdRunner.RunComplexCommand(command)
	if stderr != "" {
		err = errors.WrapErrorf(err, "stderr:%s \nexitCode:%d", stderr, exitCode)
		return
	}

	return
}

func (self *agentManager) workDir(agentName string) (r string) {
	return self.agentDir + string(filepath.Separator) + agentName
}

func (self *agentManager) Start(a *model.AgentInfo) (err error) {
	_, err = control(self.workDir(a.Name), "start", a.Timeout)
	return
}

func (self *agentManager) Stop(a *model.AgentInfo) (err error) {
	_, err = control(self.workDir(a.Name), "stop", a.Timeout)
	return
}

func (self *agentManager) Status(a *model.AgentInfo) (s string, err error) {
	s, err = control(self.workDir(a.Name), "status", a.Timeout)
	if err != nil {
		return
	}
	return
}

func (self *agentManager) Version(a *model.AgentInfo) (s string, err error) {
	s, err = control(self.workDir(a.Name), "status", a.Timeout)
	if err != nil {
		return
	}
	return
}

func (self *agentManager) Delete(a *model.AgentInfo) (err error) {
	_, err = control(self.workDir(a.Name), "delete", a.Timeout)
	if err != nil {
		return
	}

	return
}

func (self *agentManager) ControlScriptCheck(a *model.AgentInfo) bool {
	if file.IsExist(self.workDir(a.Name) + string(filepath.Separator) + getControlName()) {
		return true
	}
	return false
}
