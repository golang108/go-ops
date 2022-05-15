package script

import (
	"context"
	"go-ops/internal/model"
	"go-ops/pkg/agent/cmdrunner"
	ospsys "go-ops/pkg/system"
	"os"
	"strings"

	log "go-ops/pkg/logger"
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
		path = ScriptPath + "/" + scriptJob.Jobid
	}
	runer := cmdrunner.NewScriptCmdRunner(ospsys.NewExecCmdRunner(ctx), ScriptPath)
	p := JobScriptProvider{
		cmdRunner: runer,
	}

	// 如果文件得扩展名不是以.结尾
	if scriptJob.Script.Ext != "" && !strings.HasPrefix(scriptJob.Script.Ext, ".") {
		scriptJob.Script.Ext = "." + scriptJob.Script.Ext
	}

	if scriptJob.Script.Ext == "" {
		scriptJob.Script.Ext = ScriptExt
	}

	log.Infof("[%s] new script job, execway: %v, waitime:%d", scriptJob.Jobid, scriptJob.Script.ExecWay, scriptJob.Script.Timeout)

	switch scriptJob.Script.ExecWay {
	case model.ExecCmd:
		p.scripter = NewScript(runer, scriptJob.Jobid, path, scriptJob.Script.Content, scriptJob.Script.Env, scriptJob.Script.Timeout, scriptJob.Script.User, scriptJob.Script.Args)
	case model.ExecContent:
		p.scripter = NewContentScript(runer, scriptJob.Jobid, path, scriptJob.Script.Cmd, scriptJob.Script.Content, scriptJob.Script.Env,
			scriptJob.Script.Timeout, scriptJob.Script.Input, scriptJob.Script.User, scriptJob.Script.Args, scriptJob.Script.Ext)
	case model.ExecScriptName:
		p.scripter = NewNameScript(runer, scriptJob.Jobid, path, scriptJob.Script.Cmd, scriptJob.Script.Content, scriptJob.Script.Env, scriptJob.Script.Timeout, scriptJob.Script.Input, scriptJob.Script.User, scriptJob.Script.Args)
	case model.ExecURL:
		p.scripter = NewUrlScript(runer, scriptJob.Jobid, path, scriptJob.Script.Cmd, scriptJob.Script.Content, scriptJob.Script.Env, scriptJob.Script.Timeout, scriptJob.Script.User, scriptJob.Script.Args, scriptJob.Script.Input)

	}
	return p
}

func (p JobScriptProvider) Run() model.ResCmd {
	return p.scripter.Run()
}

func getCmdArgs(s string) (cmd string, args []string) {
	s = os.ExpandEnv(s)
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
