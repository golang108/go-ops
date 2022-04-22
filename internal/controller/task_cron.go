package controller

import (
	"context"
	v1 "go-ops/api/v1"
)

var TaskCron *taskcron = new(taskcron)

type taskcron struct{}

func (self *taskcron) Create(ctx context.Context, req *v1.AddCronTaskReq) (res *v1.CronTaskItemRes, err error) {
	return
}

func (self *taskcron) Update(ctx context.Context, req *v1.UpdateCronTaskReq) (res *v1.CronTaskItemRes, err error) {
	return
}

func (self *taskcron) Delete(ctx context.Context, req *v1.DeleteCronTaskReq) (res *v1.DeleteTaskPresetRes, err error) {

	res = new(v1.DeleteTaskPresetRes)

	res.Message = "DELETED SUCCESS!"
	return
}

func (self *taskcron) Query(ctx context.Context, req *v1.QueryCronTaskReq) (res *v1.QueryCronTaskRes, err error) {
	return
}
