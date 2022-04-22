package model

import (
	"go-ops/pkg/message"
	"go-ops/pkg/util"
	"reflect"
	"time"
)

type FileInfo struct {
	Name    string    `json:"name" dc:"文件名称"`
	Type    string    `json:"type" dc:"文件类型, dir 表示文件夹"`
	Size    int64     `json:"size" dc:"大小"`
	Path    string    `json:"path" dc:"路径"`
	ModTime time.Time `json:"mtime" dc:"修改时间"`
}

type PeerListFileInfo struct {
	Peerid string `json:"peerid" dc:"节点id"`
	Path   string `json:"path" dc:"文件夹路径"`
}

type PeerListFileInfoRes struct {
	Peerid string      `json:"peerid" dc:"节点id"`
	Path   string      `json:"path" dc:"文件夹路径"`
	List   []*FileInfo `json:"list" dc:"文件列表"`
}

type PeerNewDir struct {
	Peerid string `json:"peerid" dc:"节点id"`
	Path   string `json:"path" dc:"文件夹路径, 如果不存在会递归创建"`
}

func init() {
	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*PeerListFileInfo)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.PeerListFileInfo")),
	})

	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*PeerListFileInfoRes)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.PeerListFileInfoRes")),
	})

	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*PeerNewDir)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.PeerNewDir")),
	})

}
