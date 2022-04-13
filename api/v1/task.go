package v1

import (
	"osp/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type TaskQueryReq struct {
	g.Meta  `path:"/v1/m/task/query" tags:"Task管理" method:"post" summary:"任务查询"`
	Name    string `json:"name" dc:"任务名"`
	Creater string `json:"creater" dc:"创建人"`
	TaskID  string `json:"taskid" dc:"任务id"`
}

type TaskInfo struct {
	Task    *entity.Task `json:"task" dc:"任务"`
	Sublist []*TaskInfo  `json:"sublist" dc:"子任务详情"`
}

type TaskInfoRes struct {
	List []*TaskInfo `json:"list"`
}
