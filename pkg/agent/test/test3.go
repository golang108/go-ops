package main

import (
	"context"
	"fmt"
	"go-ops/internal/model"
	"go-ops/pkg/agent/script"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := model.Script{
		Content: "ping baidu.com",
	}
	s1 := model.ScriptJob{
		Jobid:  "1111",
		Script: s,
	}
	fmt.Println("tttt->")

	ctx, cancel := context.WithCancel(context.Background())
	go handlesig(func() {
		cancel()
	})
	res := script.NewJobScriptProvider(ctx, s1).Run()

	fmt.Println(res)

}

func handlesig(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT,
	)
	sig := <-c
	fmt.Println("sig:", sig)
	f()

}
