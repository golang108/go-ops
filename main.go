package main

import (
	"flag"
	"log"
	"os"
	"osp/internal/cmd"
	_ "osp/internal/packed"
	"osp/peer"

	"github.com/luxingwen/pnet/config"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	fg := flag.NewFlagSet("osp", flag.ContinueOnError)

	var addr string
	fg.StringVar(&addr, "addr", "", "-addr")

	err := fg.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	conf := &config.Config{}
	conf.Port = 9001
	conf.Name = "osp"

	err = peer.InitOspPeer("osp-server-2", conf)
	if err != nil {
		log.Fatal(err)
	}

	peer.GetOspPeer().Join(addr)

	cmd.Main.Run(gctx.New())
}
