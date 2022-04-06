package main

import (
	_ "osp/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"osp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
