package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/luxingwen/pnet/protos"
)

type NodeReq struct {
	g.Meta `path:"/peer/nodes" tags:"Hello" method:"post" summary:"You first hello api"`
	NodeId string `json:"nodeid"`
}

type NodeRes struct {
	Nodes []*protos.Node `json:""`
}

type NodeConnectReq struct {
	g.Meta     `path:"/peer/node/connect" tags:"Hello" method:"post" summary:"You first hello api"`
	NodeId     string `json:"nodeid"`
	RemoteAddr string `json:"remoteAddr"`
}

type NodeOpRes struct {
	Msg string `json:"msg"`
}

type NodeStopReq struct {
	g.Meta   `path:"/peer/node/stop" tags:"Hello" method:"post" summary:"You first hello api"`
	NodeId   string `json:"nodeid"`
	RemoteId string `json:"remoteId"`
}
