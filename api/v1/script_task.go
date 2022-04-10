package v1

import (
	"osp/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ScriptTaskReq struct {
	g.Meta `path:"/script/async" tags:"script" method:"post" summary:"脚本执行"`
	ScriptTask
}

type ScriptTask struct {
	Name    string       `json:"name"`
	Creater string       `json:"creater"`
	Peers   []string     `json:"peers"`
	Content model.Script `json:"content"`
}

type ScriptTaskRes struct {
	TaskId string `json:"taskid"`
}

type ScriptRes struct {
	List []*model.ResponseResCmd `json:"list"`
}

type ScriptTaskSyncReq struct {
	g.Meta `path:"/script/sync" tags:"script" method:"post" summary:"脚本执行"`
	ScriptTask
}

type ScriptTaskCancelReq struct {
	g.Meta `path:"/script/cancel" tags:"script" method:"post" summary:"取消脚本运行"`
	Tasks  []ScriptTaskCancel `json:"tasks"`
}

type ScriptTaskCancel struct {
	PeerId string `json:"peerid"`
	Jobid  string `json:"jobid"`
	Msg    string `json:"msg"`
}

type ScriptTaskCancelRes struct {
	List []*ScriptTaskCancel `json:"list"`
}

type PeerScriptTaskInfoReq struct {
	g.Meta `path:"/script/peer/taskinfo" tags:"script" method:"post" summary:"脚本信息"`
	PeerId string `json:"peerid"`
	TaskId string `json:"taskid"`
}

type ScriptTaskInfoRes struct {
	PeerId string `json:"peerid"`
	*model.TaskInfo
}

type ScriptTaskExecItem struct {
	*model.ResponseResCmd
	Status string `json:"status"`
}

type ScriptTaskExecRes struct {
	TaskId string                `json:"taskid"`
	Status string                `json:"status"`
	List   []*ScriptTaskExecItem `json:"list"`
}

type ScriptTaskInfoReq struct {
	g.Meta `path:"/script/taskinfo" tags:"script" method:"post" summary:"脚本信息"`
	TaskId string `json:"taskid"`
}
