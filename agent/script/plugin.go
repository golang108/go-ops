package script

import (
	"context"
	"encoding/json"
	"go-ops/internal/model"
	"time"
)

type ScriptPlugin struct {
	GenericScript
	cmd   string
	input string
}

func NewScriptPlugin(
	jobid string,
	path string,
	cmd string,
	content string,
	env map[string]string,
	timeout int,
	input string,
	user string,
	args []string,
) ScriptPlugin {

	if cmd == "" {
		cmd = Cmder
	}

	s := ScriptPlugin{cmd: cmd, input: input}
	s.GenericScript.path = path
	s.GenericScript.content = content
	s.GenericScript.env = env
	s.GenericScript.jobid = jobid
	s.GenericScript.timeout = timeout
	s.GenericScript.user = user
	s.GenericScript.args = args

	return s
}

func (s ScriptPlugin) Run() (r model.ResCmd) {

	opsAgentPlugin := new(AgentPlugin)

	err := json.Unmarshal([]byte(s.content), opsAgentPlugin)
	if err != nil {
		r.Err = err.Error()
		return
	}

	opsAgentPlugin.Args = s.args
	opsAgentPlugin.Cmd = s.cmd
	opsAgentPlugin.Timeout = time.Duration(s.timeout * int(time.Second))

	res, err := opsAgentPlugin.Run(context.Background(), []byte(s.input))

	if err != nil {
		r.Err = err.Error()
		return
	}

	r.Res = res

	return
}
