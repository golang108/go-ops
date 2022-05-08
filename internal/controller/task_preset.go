package controller

import (
	"context"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"
)

var TaskPreset *taskPreset = new(taskPreset)

type taskPreset struct{}

func (self *taskPreset) Create(ctx context.Context, req *v1.AddTaskPresetReq) (res *v1.TaskPresetItemRes, err error) {
	return service.TaskPreset().Create(ctx, req)
}

func (self *taskPreset) Update(ctx context.Context, req *v1.UpdateTaskPresetReq) (res *v1.TaskPresetItemRes, err error) {
	return service.TaskPreset().Update(ctx, req)
}

func (self *taskPreset) Delete(ctx context.Context, req *v1.DeleteTaskPresetReq) (res *v1.DeleteTaskPresetRes, err error) {
	err = service.TaskPreset().Delete(ctx, req)
	res = new(v1.DeleteTaskPresetRes)
	if err != nil {
		res.Message = err.Error()
		return
	}
	res.Message = "DELETED SUCCESS!"
	return
}

func (self *taskPreset) Query(ctx context.Context, req *v1.QueryTaskPresetReq) (res *v1.QueryTaskPresetRes, err error) {
	return service.TaskPreset().Query(ctx, req)
}
