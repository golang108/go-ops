package controller

import (
	"context"
	v1 "go-ops/api/v1"
)

var TaskPreset *taskPreset = new(taskPreset)

type taskPreset struct{}

func (self *taskPreset) Create(ctx context.Context, req *v1.AddTaskPresetReq) (res *v1.TaskPresetItemRes, err error) {
	return
}

func (self *taskPreset) Update(ctx context.Context, req *v1.UpdateTaskPresetReq) (res *v1.TaskPresetItemRes, err error) {
	return
}

func (self *taskPreset) Delete(ctx context.Context, req *v1.DeleteTaskPresetReq) (res *v1.QueryTaskPresetRes, err error) {
	return
}

func (self *taskPreset) Query(ctx context.Context, req *v1.QueryTaskPresetReq) (res *v1.QueryTaskPresetRes, err error) {
	return
}
