package main

import (
	"fmt"
	"go-ops/pkg/daemon"
	"os"
	"time"
)

func main() {
	dae := daemon.NewDaemon()
	dae.Start()

	fmt.Println("pid : ", os.Getegid())

	time.Sleep(time.Second)
	fmt.Println("end->")
}
