package v1

import (
	"osp/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DownloadFileReq struct {
	g.Meta  `path:"/peer/downloadfile" tags:"文件分发" method:"post" summary:"创建文件分发任务"`
	Name    string `json:"name" dc:"任务名称"`
	Creater string `json:"creater" dc:"创建人"`
	DownloadFileInfo
}

type DownloadFileInfo struct {
	Peers []string                  `json:"peers" dc:"节点列表"`
	Files []*model.DownloadFileInfo `json:"files" dc:""`
}

type DownloadfileRes struct {
	Taskid string              `json:"taskid" dc:"任务id"`
	Status string              `json:"status" dc:"状态（doing, failed, done）"`
	List   []*DownloadfileItem `json:"list" dc:"任务列表详情"`
}

type DownloadfileItem struct {
	Status string `json:"status" dc:"状态（doing, failed, done）`
	model.DownloadFileJobRes
}

type DownloadFileDetailsReq struct {
	g.Meta `path:"/peer/downloadfile/details" tags:"文件分发" method:"post" summary:"文件分发任务详情"`
	Taskid string `json:"taskid"`
}
