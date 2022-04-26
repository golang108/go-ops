package controller

import (
	"context"
	v1 "go-ops/api/v1"
	"go-ops/service"
)

var Task *task = new(task)

type task struct{}

func (self *task) Query(ctx context.Context, req *v1.TaskQueryReq) (res *v1.TaskInfoRes, err error) {
	return service.Task().QueryTask(ctx, req)
}

func (self *task) Get(ctx context.Context, req *v1.TaskInfoReq) (res v1.TaskDetailRes, err error) {
	taskinfo, err := service.Task().GetTaskInfo(ctx, req)
	if err != nil {
		return
	}
	res = v1.TaskDetailRes(*taskinfo)
	return
}
