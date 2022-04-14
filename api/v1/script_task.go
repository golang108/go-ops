package v1

import (
	"osp/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ScriptTaskReq struct {
	g.Meta `path:"/script/async" tags:"脚本任务" method:"post" summary:"脚本异步执行"`
	ScriptTask
}

type ScriptTask struct {
	Name    string       `json:"name" dc:"脚本任务名"`
	Creater string       `json:"creater" dc:"创建者"`
	Peers   []string     `json:"peers" dc:"节点id列表"`
	Content model.Script `json:"content" dc:"脚本内容信息"`
}

type ScriptTaskRes struct {
	TaskId string `json:"taskid" dc:"任务id"`
}

type ScriptRes struct {
	List []*model.ResponseResCmd `json:"list" dc:"脚本任务详情列表"`
}

type ScriptTaskSyncReq struct {
	g.Meta `path:"/script/sync" tags:"脚本任务" method:"post" summary:"脚本同步执行"`
	ScriptTask
}

type ScriptTaskCancelReq struct {
	g.Meta `path:"/script/cancel" tags:"脚本任务" method:"post" summary:"取消脚本运行"`
	Tasks  []ScriptTaskCancel `json:"tasks"`
}

type ScriptTaskCancel struct {
	PeerId string `json:"peerid" dc:"节点id"`
	Jobid  string `json:"jobid" dc:"任务id"`
	Msg    string `json:"msg"`
}

type ScriptTaskCancelRes struct {
	List []*ScriptTaskCancel `json:"list"`
}

type PeerScriptTaskInfoReq struct {
	g.Meta `path:"/script/peer/taskinfo" tags:"脚本任务" method:"post" summary:"远程节点上的脚本任务信息"`
	PeerId string `json:"peerid" dc:"节点id"`
	TaskId string `json:"taskid" dc:"任务id"`
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
	TaskId string                `json:"taskid" dc:"任务id"`
	Status string                `json:"status"`
	List   []*ScriptTaskExecItem `json:"list"`
}

type ScriptTaskInfoReq struct {
	g.Meta `path:"/script/taskinfo" tags:"脚本任务" method:"post" summary:"脚本任务信息"`
	TaskId string `json:"taskid"`
}
