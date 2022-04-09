package script

import (
	"context"
	"osp/agent/cmdrunner"
	"osp/internal/model"
	ospsys "osp/pkg/system"
	"strings"
)

type JobScriptProvider struct {
	cmdRunner cmdrunner.CmdRunner
	scripter  Script
}

func NewJobScriptProvider(
	ctx context.Context,
	scriptJob model.ScriptJob,
) JobScriptProvider {
	path := scriptJob.Script.Path
	if path == "" {
		path = ScriptPath
	}
	runer := cmdrunner.NewScriptCmdRunner(ospsys.NewExecCmdRunner(ctx), path)
	p := JobScriptProvider{
		cmdRunner: runer,
	}

	switch scriptJob.Script.ExecWay {
	case model.ExecCmd:
		p.scripter = NewScript(runer, scriptJob.Jobid, path, scriptJob.Script.Content, scriptJob.Script.Env, scriptJob.Script.Timeout, scriptJob.Script.User, scriptJob.Script.Args)
	case model.ExecContent:
		p.scripter = NewContentScript(runer, scriptJob.Jobid, path, scriptJob.Script.Cmd, scriptJob.Script.Content, scriptJob.Script.Env, scriptJob.Script.Timeout, scriptJob.Script.Input, scriptJob.Script.User, scriptJob.Script.Args)

	}
	return p
}

func (p JobScriptProvider) Run() model.ResCmd {
	return p.scripter.Run()
}

func getCmdArgs(s string) (cmd string, args []string) {
	list := strings.Fields(s)
	if len(list) <= 0 {
		return
	}
	if len(list) == 1 {
		cmd = list[0]
		return
	}

	cmd = list[0]
	args = list[1:]
	return

}
