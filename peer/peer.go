package peer

import (
	"go-ops/pkg/message"
	"time"

	"github.com/luxingwen/pnet/log"
	"github.com/luxingwen/pnet/protos"

	"github.com/luxingwen/pnet/config"

	"github.com/luxingwen/pnet"
	"github.com/luxingwen/pnet/node"
)

var defaultPN *pnet.PNet

type HandlerPNetMsg func(msg interface{}, msgID []byte, srcID, rpath string, pn *pnet.PNet)

func NewPnet(id string, conf *config.Config, h HandlerPNetMsg) (pn *pnet.PNet, err error) {
	pn, err = pnet.NewPNet(id, conf)
	if err != nil {
		return
	}

	pn.ApplyMiddleware(node.BytesReceived{func(msg, msgID []byte, srcID, rpath string, remoteNode *node.RemoteNode) ([]byte, bool) {
		v, err := message.JSONCodec.Decode(msg)
		if err != nil {
			log.Error("json decode err:", err)
			return nil, true
		}

		h(v, msgID, srcID, rpath, pn)
		return nil, true
	}, 0})

	pn.ApplyMiddleware(node.LocalNodeStarted{func(lc *node.LocalNode) bool {
		lc.SetReady(true)
		return true
	}, 0})

	return
}

// 发送异步消息
func SendMsgAsync(pn *pnet.PNet, msg interface{}, srcID, rpath string) (err error) {
	data, err := message.JSONCodec.Encode(msg)
	if err != nil {
		log.Error("json encode err:", err)
		return
	}
	_, err = pn.SendBytesRelayAsync(data, srcID)
	return
}

// 发送同步消息
func SendMsgSync(pn *pnet.PNet, msg interface{}, srcID, rpath string) (r []byte, err error) {
	return SendMsgSyncWithTimeout(pn, msg, srcID, rpath, 0)
}

// 发送同步消息
func SendMsgSyncWithTimeout(pn *pnet.PNet, msg interface{}, srcID, rpath string, timeout time.Duration) (r []byte, err error) {
	data, err := message.JSONCodec.Encode(msg)
	if err != nil {
		log.Error("json encode err:", err)
		return
	}
	r, _, err = pn.SendBytesRelaySyncWithTimeout(data, srcID, timeout)
	return
}

// 回复消息
func SendMsgReplay(pn *pnet.PNet, msg interface{}, msgID []byte, srcID, rpath string) (err error) {
	data, err := message.JSONCodec.Encode(msg)
	if err != nil {
		log.Error("json encode err:", err)
		return
	}
	_, err = pn.SendBytesRelayReply(msgID, data, srcID)
	return
}

func SendMsgBroadCast(pn *pnet.PNet, msg interface{}) (err error) {
	data, err := message.JSONCodec.Encode(msg)
	if err != nil {
		log.Error("json encode err:", err)
		return
	}

	_, err = pn.SendMessageAsync(pn.GetLocalNode().NewBraodcastMessage(data), protos.BROADCAST)
	return
}
