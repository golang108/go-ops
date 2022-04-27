package action

import (
	"go-ops/internal/model"
	"path/filepath"

	"github.com/toolkits/file"
)

func (self *agentManager) Start(a *model.AgentInfo) (err error) {
	workdir := self.agentDir + string(filepath.Separator) + a.Name
	err = control(workdir, "start")
	return
}

func (self *agentManager) Stop(a *model.AgentInfo) (err error) {
	workdir := self.agentDir + string(filepath.Separator) + a.Name
	err = control(workdir, "stop")
	return
}

func (self *agentManager) Status(a *model.AgentInfo) (s string, err error) {
	workdir := self.agentDir + string(filepath.Separator) + a.Name
	err = control(workdir, "status")
	if err != nil {
		return
	}
	return
}

func (self *agentManager) Version(a *model.AgentInfo) (s string, err error) {
	return
}

func (self *agentManager) Delete(a *model.AgentInfo) (err error) {
	workdir := self.agentDir + string(filepath.Separator) + a.Name
	err = control(workdir, "delete")
	return
}

func (self *agentManager) ControlScriptCheck(a *model.AgentInfo) bool {
	workDir := self.agentDir + string(filepath.Separator) + a.Name
	if file.IsExist(workDir + string(filepath.Separator) + getControlName()) {
		return true
	}
	return false
}
