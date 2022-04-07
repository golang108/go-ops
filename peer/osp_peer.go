package peer

import (
	"encoding/json"

	"github.com/luxingwen/pnet"
	"github.com/luxingwen/pnet/config"
	"github.com/luxingwen/pnet/log"
)

type OspPeer struct {
	*pnet.PNet
}

func (self *OspPeer) HandlerFunc(msg interface{}, msgID []byte, srcID, rpath string, pn *pnet.PNet) {
	switch v := msg.(type) {
	default:
		b, _ := json.Marshal(v)
		log.Error("msg handler not found,msg: %s", string(b))
	}
}

var ospPeer *OspPeer

func InitOspPeer(id string, conf *config.Config) (err error) {

	ospPeer = &OspPeer{}

	pn, err := NewPnet(id, conf, ospPeer.HandlerFunc)
	if err != nil {
		return
	}
	ospPeer.PNet = pn
	pn.Start()
	return
}

func GetOspPeer() *OspPeer {
	if ospPeer == nil {
		panic("osp peer is nil")
	}
	return ospPeer
}
