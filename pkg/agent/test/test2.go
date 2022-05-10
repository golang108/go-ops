package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	res, err := exec.Command("echo", "$HOME").CombinedOutput()
	if err != nil {
		return
	}

	s := os.ExpandEnv("echo hello")
	fmt.Println("s:", s)

	fmt.Println(string(res))
}
