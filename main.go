package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"osp/internal/cmd"
	_ "osp/internal/packed"
	"osp/peer"

	"github.com/luxingwen/pnet/config"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/uuid"
)

func main() {
	fg := flag.NewFlagSet("osp", flag.ContinueOnError)

	var addr string
	var id string
	var port uint64
	var h bool
	fg.StringVar(&addr, "addr", "", "-addr")
	fg.StringVar(&id, "id", "", "-id")
	fg.Uint64Var(&port, "port", 9999, "-port port")
	fg.BoolVar(&h, "h", false, "-h help")

	err := fg.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if h {
		fg.Usage()
		return
	}

	if id == "" {
		id = uuid.New().String()
	}

	id += "@osp"
	conf := &config.Config{}
	conf.Port = uint16(port)
	conf.Name = "osp"

	err = peer.InitOspPeer(id, conf)
	if err != nil {
		log.Fatal(err)
	}

	if addr != "" {
		fmt.Println("join addr:", addr)
		_, err = peer.GetOspPeer().Join(addr)
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd.Main.Run(gctx.New())
}
