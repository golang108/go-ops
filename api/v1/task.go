package v1

import (
	"go-ops/model/entity"

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

type AddTaskPresetReq struct {
	g.Meta  `path:"/v1/m/task/preset/create" tags:"Task管理" method:"post" summary:"创建预设任务"`
	Type    string `json:"type" dc:"任务类型"`
	Name    string `json:"name" dc:"任务名"`
	Creater string `json:"creater" dc:"创建人"`
	Content string `json:"content" dc:"预设任务内容"`
}

type UpdateTaskPresetReq struct {
	g.Meta  `path:"/v1/m/task/preset/update" tags:"Task管理" method:"post" summary:"更新预设任务"`
	Uuid    string `json:"uuid" dc:"预设任务uuid"`
	Type    string `json:"type" dc:"任务类型"`
	Name    string `json:"name" dc:"任务名"`
	Creater string `json:"creater" dc:"创建人"`
	Content string `json:"content" dc:"预设任务内容"`
}

type DeleteTaskPresetReq struct {
	g.Meta `path:"/v1/m/task/preset/deleted" tags:"Task管理" method:"post" summary:"删除预设任务"`
	Uuids  []string `json:"uuids" dc:"预设任务uuid列表"`
}

type DeleteTaskPresetRes struct {
	Message string `json:"message"`
}

type QueryTaskPresetReq struct {
	g.Meta  `path:"/v1/m/task/preset/query" tags:"Task管理" method:"post" summary:"查询预设任务"`
	Uuid    string `json:"uuid" dc:"预设任务uuid"`
	Name    string `json:"name" dc:"任务名"`
	Creater string `json:"creater" dc:"创建人"`
}

type QueryTaskPresetRes struct {
	List []*TaskPresetItemRes `json:"list"`
}

type TaskPresetItemRes struct {
	Uuid    string `json:"uuid" dc:"预设任务uuid"`
	Type    string `json:"type" dc:"任务类型"`
	Name    string `json:"name" dc:"任务名"`
	Creater string `json:"creater" dc:"创建人"`
	Content string `json:"content" dc:"预设任务内容"`
	Created string `json:"created" dc:"创建时间"`
	Updated string `json:"updated" dc:"更新时间"`
}
