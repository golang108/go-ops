package model

import (
	"go-ops/pkg/message"
	"go-ops/pkg/util"
	"reflect"
)

type DownloadFileInfo struct {
	Filename       string `json:"filename" dc:"文件名称"`              // 文件名称
	Address        string `json:"address" dc:"文件的下载url地址"`         // 文件下载地址
	Path           string `json:"path" dc:"文件的保存路径"`               // 文件保存路径
	AutoCreatePath bool   `json:"auto_create_path" dc:"是否自动创建文件夹"` // 自动创建文件夹
	Replace        bool   `json:"replace" dc:"如果文件已经存在,是否替换文件"`    // 是否替换
}

type DownloadFileRes struct {
	Filename string  `json:"filename" dc:"文件名"` //
	Code     ResCode `json:"code" dc:"代码"`
	Msg      string  `json:"msg"`
}

type DownloadFileJobRes struct {
	PeerId string `json:"peerid" dc:"节点id"`
	Jobid  string `json:"jobid" dc:"任务id"` // 任务id
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
