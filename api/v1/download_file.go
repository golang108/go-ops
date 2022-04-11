package v1

import (
	"osp/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DownloadFileReq struct {
	g.Meta  `path:"/peer/downloadfile" tags:"Hello" method:"post" summary:"You first hello api"`
	Name    string `json:"name"`
	Creater string `json:"creater"`
	DownloadFileInfo
}

type DownloadFileInfo struct {
	Peers []string                  `json:"peers"`
	Files []*model.DownloadFileInfo `json:"files"`
}

type DownloadfileRes struct {
	Taskid string              `json:"taskid"`
	Status string              `json:"status"`
	List   []*DownloadfileItem `json:"list"`
}

type DownloadfileItem struct {
	Status string `json:"status"`
	model.DownloadFileJobRes
}

type DownloadFileDetailsReq struct {
	g.Meta `path:"/peer/downloadfile/details" tags:"Hello" method:"post" summary:"You first hello api"`
	Taskid string `json:"taskid"`
}
