package main

import (
	"log"
	"osp/internal/cmd"
	_ "osp/internal/packed"
	"osp/peer"

	"github.com/luxingwen/pnet/config"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	conf := &config.Config{}
	conf.Port = 8888
	conf.Name = "osp"

	err := peer.InitOspPeer("osp-server-1", conf)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Main.Run(gctx.New())
}
