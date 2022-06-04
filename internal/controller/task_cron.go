package controller

import (
	"context"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"

	"github.com/gogf/gf/v2/os/glog"
)

var TaskCron *taskcron = new(taskcron)

type taskcron struct{}

func (self *taskcron) Create(ctx context.Context, req *v1.AddCronTaskReq) (res *v1.CronTaskItemRes, err error) {
	res, err = service.TaskCron().Create(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "create task cron err:%v", err)
	}
	return
}

func (self *taskcron) Update(ctx context.Context, req *v1.UpdateCronTaskReq) (res *v1.CronTaskItemRes, err error) {
	return service.TaskCron().Update(ctx, req)
}

func (self *taskcron) Delete(ctx context.Context, req *v1.DeleteCronTaskReq) (res *v1.DeleteTaskPresetRes, err error) {

	res = new(v1.DeleteTaskPresetRes)
	err = service.TaskCron().Delete(ctx, req)
	if err != nil {
		res.Message = err.Error()
		glog.Errorf(ctx, "delete task cron err:%v", err)
		return
	}
	res.Message = "DELETED SUCCESS!"
	return
}

func (self *taskcron) Query(ctx context.Context, req *v1.QueryCronTaskReq) (res *v1.QueryCronTaskRes, err error) {
	res, err = service.TaskCron().Query(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "query task cron err:%v", err)
	}
	return
}

func (self *taskcron) Start(ctx context.Context, req *v1.StartCronTaskReq) (res v1.CronTaskOpRes, err error) {
	return
}

func (self *taskcron) Stop(ctx context.Context, req *v1.StopCronTaskReq) (res v1.CronTaskOpRes, err error) {
	return
}
