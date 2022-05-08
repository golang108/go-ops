package main

import (
	"flag"
	"fmt"
	"go-ops/internal/peer"
	"go-ops/pkg/agent"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/luxingwen/pnet/config"
	"github.com/shirou/gopsutil/v3/host"
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
	var port uint64
	var name string
	var id string
	var h bool
	fg.StringVar(&addr, "addr", "tcp://127.0.0.1:9000", "-addr")
	fg.Uint64Var(&port, "port", 13333, "-port port")
	fg.StringVar(&name, "name", "", "-name name")
	fg.StringVar(&id, "id", "", "-id name")
	fg.BoolVar(&h, "h", false, "-h help")

	err := fg.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if h {
		fg.Usage()
		return
	}

	ospAgent := agent.NewOspAgent("./")

	conf := &config.Config{}

	if name == "" {
		name, err = os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
	}

	if id == "" {
		hid, err := host.HostID()
		if err != nil {
			id = uuid.New().String()
		} else {
			id = hid
		}
	}

	id = id + "@agent"

	conf.Name = name
	conf.Port = uint16(port)

	p, err := peer.NewPnet(id, conf, ospAgent.HandlerFunc)
	if err != nil {
		fmt.Println("new peer err:")
		return
	}

	//p.GetLocalNode().SetType("client")

	err = p.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("join add:", addr)

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
