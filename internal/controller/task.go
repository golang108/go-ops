package controller

import (
	"context"
	v1 "osp/api/v1"
)

var Task *task = new(task)

type task struct{}

func (self *task) Query(ctx context.Context, req *v1.TaskQueryReq) (res *v1.TaskInfoRes, err error) {
	return
}
