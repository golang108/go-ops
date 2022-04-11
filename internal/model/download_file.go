package model

import (
	"osp/pkg/message"
	"osp/pkg/util"
	"reflect"
)

type DownloadFileInfo struct {
	Filename       string `json:"filename"`         // 文件名称
	Address        string `json:"address"`          // 文件下载地址
	Path           string `json:"path"`             // 文件保存路径
	AutoCreatePath bool   `json:"auto_create_path"` // 自动创建文件夹
	Replace        bool   `json:"replace"`          // 是否替换
}

type DownloadFileRes struct {
	Filename string  `json:"filename"` //
	Code     ResCode `json:"code"`
	Msg      string  `json:"msg"`
}

type DownloadFileJobRes struct {
	PeerId string `json:"peerid"`
	Jobid  string `json:"jobid"` // 任务id
	*DownloadFileRes
}

type DownloadFileJob struct {
	Jobid string `json:"jobid"` // 任务id
	*DownloadFileInfo
}

func init() {
	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*DownloadFileJob)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.DownloadFileJob")),
	})
	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*DownloadFileJobRes)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.DownloadFileJobRes")),
	})
}
