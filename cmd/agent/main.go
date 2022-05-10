package main

import (
	"go-ops/pkg/agent"
	"log"
)

func main() {
	err := agent.Main()
	if err != nil {
		log.Fatal(err)
	}
}
