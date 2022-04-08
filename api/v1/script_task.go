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
	Name      string       `json:"name"`
	Creater   string       `json:"creater"`
	Hostinfos []string     `json:"hostinfos"`
	Content   model.Script `json:"content"`
}

type ScriptTaskRes struct {
	TaskId string `json:"taskid"`
}

type ScriptTaskSyncReq struct {
	g.Meta `path:"/script/sync" tags:"script" method:"post" summary:"脚本执行"`
	ScriptTask
}
