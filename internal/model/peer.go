package model

import (
	"go-ops/pkg/message"
	"go-ops/pkg/util"
	"reflect"
)

type PeerInfo struct {
	PeerId   string `json:"peerId" `  // 节点uuid
	HostName string `json:"hostname"` // 节点hostname
	Name     string `json:"name"`     // 节点名称
	Address  string `json:"address"`  // 节点地址
	PublicIp string `json:"publicIp"` // 节点ip
	Os       string `json:"os"`
	Arch     string `json:"arch"`
}

type GetPeerInfo struct {
	HostName string `json:"hostname"` // 节点hostname
}

func init() {
	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*PeerInfo)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.PeerInfo")),
	})

	message.RegisterMessage(&message.MessageMeta{
		Type: reflect.TypeOf((*GetPeerInfo)(nil)).Elem(),
		ID:   uint32(util.StringHash("model.GetPeerInfo")),
	})

}
