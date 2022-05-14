package main

import (
	"go-ops/internal/cmd"
	_ "go-ops/internal/packed"
	"go-ops/internal/peer"

	"github.com/gogf/gf/v2/os/glog"

	"github.com/luxingwen/pnet/config"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/uuid"
)

func main() {

	log := glog.DefaultLogger()
	var ctx = gctx.New()
	var port uint16 = 9999

	portvar, err := g.Cfg().Get(ctx, "peer.port")
	if err != nil {
		log.Warning(ctx, "没有配置节点端口,将使用默认端口：", port)
	} else {
		port = portvar.Uint16()
	}

	var addrs []string

	boots, err := g.Cfg().Get(ctx, "peer.boots")
	if err != nil {
		log.Warning(ctx, "获取引导节点失败,err:", err)
	} else {
		for _, item := range boots.Array() {
			dataval := item.(map[string]interface{})
			addrs = append(addrs, dataval["addr"].(string))
		}
	}

	id := uuid.New().String()

	id += "@ops-apiserver"
	conf := &config.Config{}
	conf.Port = uint16(port)
	conf.Name = "ops-apiserver"

	err = peer.InitOspPeer(id, conf)
	if err != nil {
		log.Fatal(ctx, err)
	}

	for _, addr := range addrs {
		_, err = peer.GetOspPeer().Join(addr)
		if err != nil {
			log.Warning(ctx, "join addr:", addr, " err: ", err)
		}
	}

	cmd.Main.Run(gctx.New())
}
