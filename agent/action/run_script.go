package action

import (
	"osp/agent/model"
	"osp/agent/script"
)

type RunScriptAction struct {
}

func (a RunScriptAction) Run(req model.ScriptJob) model.ResCmd {
	return script.NewJobScriptProvider(req).Run()
}
