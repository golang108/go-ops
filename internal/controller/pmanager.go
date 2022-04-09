package controller

import (
	"context"
	"errors"
	v1 "osp/api/v1"
	"osp/peer"
	"time"

	"github.com/luxingwen/pnet/log"

	"github.com/gogo/protobuf/proto"
	"github.com/luxingwen/pnet/protos"
)

var PeerManagaer = new(peerManager)

type peerManager struct {
}

func (self *peerManager) GetNodes(ctx context.Context, nodeReq *v1.NodeReq) (res *v1.NodeRes, err error) {
	res = new(v1.NodeRes)
	if nodeReq.NodeId == "" {
		rns, err := peer.GetOspPeer().GetLocalNode().GetNeighbors(nil)
		if err != nil {
			return nil, err
		}
		for _, item := range rns {
			res.Nodes = append(res.Nodes, item.Node.Node)
		}
		return res, nil
	}

	msg := peer.GetOspPeer().GetLocalNode().NewGetNeighborsMessage(nodeReq.NodeId)

	resMsg, ok, err := peer.GetOspPeer().SendMessageSync(msg, protos.RELAY, time.Minute)
	if err != nil {
		return
	}

	if !ok {
		err = errors.New("获取节点信息失败")
		return
	}

	nodes := &protos.Neighbors{}

	err = proto.Unmarshal(resMsg.Message, nodes)
	if err != nil {
		log.Error("unmarshal err:", err)
		return
	}

	res.Nodes = nodes.Nodes
	return
}

func (self *peerManager) NodeConnect(ctx context.Context, req *v1.NodeConnectReq) (res *v1.NodeOpRes, err error) {
	if req.NodeId == "" {
		req.NodeId = peer.GetOspPeer().GetLocalNode().GetId()
	}

	msg, err := peer.GetOspPeer().GetLocalNode().NewConnnetNodeMessage(req.NodeId, &protos.Node{Addr: req.RemoteAddr})
	if err != nil {
		return
	}

	resMsg, ok, err := peer.GetOspPeer().SendMessageSync(msg, protos.RELAY, time.Minute)
	if err != nil {
		return
	}

	if !ok {
		err = errors.New("链接节点失败")
		return
	}
	res = new(v1.NodeOpRes)

	res.Msg = string(resMsg.Message)
	return

}

func (self *peerManager) StopConnect(ctx context.Context, req *v1.NodeStopReq) (res *v1.NodeOpRes, err error) {
	if req.NodeId == "" {
		req.NodeId = peer.GetOspPeer().GetLocalNode().GetId()
	}

	msg, err := peer.GetOspPeer().GetLocalNode().NewStopConnnetNodeMessage(req.NodeId, &protos.Node{Id: req.RemoteId})
	if err != nil {
		return
	}

	resMsg, ok, err := peer.GetOspPeer().SendMessageSync(msg, protos.RELAY, time.Minute)
	if err != nil {
		return
	}

	if !ok {
		err = errors.New("停止链接节点失败")
		return
	}
	res = new(v1.NodeOpRes)

	res.Msg = string(resMsg.Message)
	return
}
