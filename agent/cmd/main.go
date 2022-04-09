package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"osp/agent"
	"osp/peer"
	"syscall"
	"time"

	"github.com/luxingwen/pnet/config"
)

func InitSignal() {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for {
			s := <-sig
			fmt.Println("Got signal:", s)
			os.Exit(1)
		}
	}()
}

func main() {

	fg := flag.NewFlagSet("agent", flag.ContinueOnError)

	var addr string
	fg.StringVar(&addr, "addr", "tcp://127.0.0.1:9000", "-addr")

	err := fg.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

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

	remoteNode, err := p.Join(addr)
	if err != nil {
		log.Fatal("join:", err)
	}

	InitSignal()

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

}
