package main

import (
	"fmt"
	"log"
	"osp/agent"
	"osp/peer"

	"github.com/luxingwen/pnet/config"
)

func main() {

	ospAgent := agent.NewOspAgent("./")

	conf := &config.Config{}
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

	_, err = p.Join("tcp://127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
