package v1

import (
	"go-ops/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/luxingwen/pnet/protos"
	"github.com/luxingwen/pnet/stat"
)

type NodeReq struct {
	g.Meta `path:"/peer/nodes" tags:"节点管理" method:"post" summary:"获取节点连接信息"`
	NodeId string `json:"nodeid" dc:"节点id,空表示当前节点"`
}

type NodeRes struct {
	Nodes []*protos.Node `json:"" dc:"节点列表"`
}

type NodeConnectReq struct {
	g.Meta     `path:"/peer/node/connect" tags:"节点管理" method:"post" summary:"连接节点"`
	NodeId     string `json:"nodeid" dc:"节点id"`
	RemoteAddr string `json:"remoteAddr" dc:"远程节点连接地址"`
}

type NodeOpRes struct {
	Msg string `json:"msg"`
}

type NodeStopReq struct {
	g.Meta   `path:"/peer/node/stop" tags:"节点管理" method:"post" summary:"停止节点连接"`
	NodeId   string `json:"nodeid"`
	RemoteId string `json:"remoteId"`
}

type NodeStatReq struct {
	g.Meta `path:"/peer/node/stat" tags:"节点管理" method:"post" summary:"获取节点状态"`
	NodeId string `json:"nodeid"`
}

type NodeStatRes struct {
	*stat.NodeStat
}

type NodeFileListReq struct {
	g.Meta `path:"/peer/node/files" tags:"节点管理" method:"post" summary:"获取节点文件夹信息"`
	NodeId string `json:"nodeid"`
	Path   string `json:"path" dc:"路径"`
}

type NodeFileListRes struct {
	Files []*model.FileInfo `json:"files"`
}

type NodeFileCreateDirReq struct {
	g.Meta `path:"/peer/node/files/createDir" tags:"节点管理" method:"post" summary:"创建节点文件夹"`
	NodeId string `json:"nodeid"`
	Path   string `json:"path" dc:"文件夹路径"`
}

type NodeFileDeleteReq struct {
	g.Meta `path:"/peer/node/files/delete" tags:"节点管理" method:"post" summary:"删除节点文件"`
	NodeId string `json:"nodeid"`
	Path   string `json:"path" dc:"文件夹(文件)路径"`
}

type NodeFileMoveReq struct {
	g.Meta `path:"/peer/node/files/move" tags:"节点管理" method:"post" summary:"移动节点文件"`
	NodeId string `json:"nodeid"`
	Src    string `json:"src" dc:"源文件夹(文件)路径"`
	Dst    string `json:"dst" dc:"目标文件夹(文件)路径"`
}
