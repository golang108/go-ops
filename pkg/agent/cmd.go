package agent

import (
	"encoding/json"
	"errors"
	"go-ops/internal/peer"
	"io/ioutil"
	"os"
	"time"

	"github.com/luxingwen/pnet/config"
	"github.com/luxingwen/pnet/node"
	"github.com/shirou/gopsutil/v3/host"

	agentConf "go-ops/pkg/agent/config"

	"github.com/google/uuid"

	log "go-ops/pkg/logger"
)

func Main() (err error) {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Error("read config.json err:", err)
		return
	}

	var cfg agentConf.Config
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		log.Error("Unmarshal err:", err)
		return
	}

	opsAgent := NewOspAgent(&cfg)

	name, err := os.Hostname()
	if err != nil {
		log.Error("Unmarshal err:", err)
		return err
	}

	conf := &config.Config{
		Port: uint16(cfg.Port),
		Name: name,
	}

	id := ""
	hid, err := host.HostID()
	if err != nil {
		id = uuid.New().String()
	} else {
		id = hid
	}

	id = id + "@agent"

	p, err := peer.NewPnet(id, conf, opsAgent.HandlerFunc)
	if err != nil {
		log.Error("new peer err:", err)
		return
	}

	err = p.ApplyMiddleware(node.RemoteNodeReady{func(rmoteNode *node.RemoteNode) bool {
		peer.SendMsgBroadCast(p, opsAgent.GetPeerInfo(p))
		return true
	}, 0})

	if err != nil {
		log.Error("apply middle ware err:", err)
	}

	err = p.Start()
	if err != nil {
		log.Error("peer start err:", err)
		return
	}

	if len(cfg.Bootlist) == 0 {
		log.Error("need join remote peer")
		return errors.New("need join remote peer")
	}
	addr := cfg.Bootlist[0]

	var remoteNode *node.RemoteNode

	for {
		remoteNode, err = p.Join(addr)
		if err != nil {
			log.Errorf("join remote perr: %s  err:%v", addr, err)
			time.Sleep(time.Second)
			continue
		}
		break
	}

	for {
		time.Sleep(time.Second)
		if remoteNode.IsStopped() {
			remoteNode1, err := p.Join(remoteNode.Addr)
			if err != nil {
				log.Errorf("join remote perr: %s  err:%v", addr, err)
			} else if remoteNode1 != nil {
				remoteNode = remoteNode1
			}
		}
	}

}
