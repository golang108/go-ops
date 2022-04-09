package main

import (
	"fmt"
	"log"
	"os"
	"osp/agent"
	"osp/peer"
	"time"

	"github.com/luxingwen/pnet/config"
)

func main() {

	ospAgent := agent.NewOspAgent("./")

	conf := &config.Config{}
	name, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
	}

	conf.Name = name
	conf.Port = 33553

	p, err := peer.NewPnet("agent-1", conf, ospAgent.HandlerFunc)
	if err != nil {
		fmt.Println("new peer err:")
		return
	}

	err = p.Start()
	if err != nil {
		log.Fatal(err)
	}

	remoteNode, err := p.Join("tcp://127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second)
		if remoteNode.IsStopped() {
			remoteNode1, err := p.Join(remoteNode.Addr)
			if err != nil {
				fmt.Println("remote on connection err:", err)
			} else if remoteNode1 != nil {
				remoteNode = remoteNode1
			}
		}
	}

	select {}
}
