package controller

import (
	"context"
	"fmt"

	v1 "osp/api/v1"
	"osp/internal/model"
	"osp/peer"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	s := model.Script{
		Content: "ping baidu.com",
		Timeout: 10,
	}
	s1 := &model.ScriptJob{
		Jobid:  "jobid-1",
		Script: s,
	}

	err = peer.SendMsgAsync(peer.GetOspPeer().PNet, s1, "agent-1", "")
	if err != nil {
		fmt.Println("send msg err:", err)
	}

	s1.Script.Timeout = 3
	s1.Script.Content = "ls"
	r, err := peer.SendMsgSync(peer.GetOspPeer().PNet, s1, "agent-1", "")
	if err != nil {
		fmt.Println("send msg err:", err)
	}

	fmt.Println("r:", string(r))

	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
