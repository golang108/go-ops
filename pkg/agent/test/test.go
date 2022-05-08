package main

import (
	"context"
	"fmt"
	"go-ops/internal/model"
	"go-ops/pkg/agent"
	"go-ops/pkg/agent/script"
	"go-ops/pkg/agent/task"
	"time"
)

func main() {

	s := model.Script{
		Content: "ping baidu.com",
	}
	s1 := model.ScriptJob{
		Jobid:  "1111",
		Script: s,
	}

	ctx, cancel := context.WithCancel(context.Background())
	startFunc := func() (r interface{}, err error) {

		res := script.NewJobScriptProvider(ctx, s1).Run()
		r = &model.ResponseResCmd{
			Jobid:  s1.Jobid,
			ResCmd: res,
		}

		return
	}

	endFunc := func(t task.Task) {

		fmt.Println("end -> ")

	}

	c := func(t task.Task) error {
		fmt.Println("取消")
		cancel()
		return nil
	}

	ospAgent := agent.NewOspAgent("./")
	t := ospAgent.CreateTask(s1.Jobid, startFunc, c, endFunc)
	ospAgent.StartTask(t)
	go func() {
		time.Sleep(time.Second * 5)
		t.Cancel()
	}()
	//	ospAgent.StartTask(t)

	select {}

}
