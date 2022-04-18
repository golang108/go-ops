package main

import (
	"flag"
	"fmt"
	"go-ops/agent"
	"go-ops/peer"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/luxingwen/pnet/config"
	"github.com/shirou/gopsutil/v3/host"
)

func InitSignal() {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		s := <-sig
		fmt.Println("Got signal:", s)
		os.Exit(1)
	}

}

func getClientIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}

	return "127.0.0.1"
}

func main() {

	fg := flag.NewFlagSet("relay-peer", flag.ContinueOnError)

	var addr string
	var port uint64
	var name string
	var id string
	var h bool
	fg.StringVar(&addr, "addr", "", "-addr")
	fg.Uint64Var(&port, "port", 12333, "-port port")
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
		name += "@relay"
	}

	if id == "" {
		hid, err := host.HostID()
		if err != nil {
			id = uuid.New().String()
		} else {
			id = hid
		}
	}

	id = id + "@relay"

	conf.Name = name
	conf.Port = uint16(port)

	p, err := peer.NewPnet(id, conf, ospAgent.HandlerFunc)
	if err != nil {
		fmt.Println("new peer err:")
		return
	}

	err = p.Start()
	if err != nil {
		log.Fatal(err)
	}

	localaddr := fmt.Sprintf("tcp://%s:%d", getClientIp(), port)
	fmt.Println("relay start addr:", localaddr)

	if addr != "" {
		fmt.Println("join add:", addr)
		_, err := p.Join(addr)
		if err != nil {
			log.Fatal("join:", err)
		}
	}

	InitSignal()

}
