package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	rp, _ := io.Pipe()

	data := make([]byte, 1024)
	for {
		_, err := rp.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("data:", string(data))
	}
}
