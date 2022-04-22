package service

import (
	"context"
	v1 "go-ops/api/v1"
	"go-ops/model/entity"
	"go-ops/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

type (
	sTaskCron struct{}
)

var (
	insTaskCron = sTaskCron{}
)

func TaskCron() *sTaskCron {
	return &insTaskCron
}

func (self *sTaskCron) Create(ctx context.Context, req *v1.AddCronTaskReq) (res *v1.CronTaskItemRes, err error) {

	item := &entity.CronTask{
		CronUid:  guid.S(),
		Type:     req.Type,
		Content:  req.Content,
		CronExpr: req.CronExpr,
		Name:     req.Name,
		Creater:  req.Creater,
		Created:  gtime.Now(),
		Status:   req.Status,
	}

	_, err = dao.CronTask.Ctx(ctx).Data(item).Insert()

	if err != nil {
		return
	}

	res = &v1.CronTaskItemRes{
		CronUid:  item.CronUid,
		Type:     item.Type,
		Content:  item.Content,
		CronExpr: item.CronExpr,
		Name:     item.Name,
		Creater:  item.Creater,
		Created:  item.Created.String(),
		Status:   item.Status,
	}

	return
}

func (self *sTaskCron) Update(ctx context.Context, req *v1.UpdateCronTaskReq) (res *v1.CronTaskItemRes, err error) {

	item := &entity.CronTask{
		Type:     req.Type,
		Content:  req.Content,
		CronExpr: req.CronExpr,
		Name:     req.Name,
		Updater:  req.Updater,
		Updated:  gtime.Now(),
		Status:   req.Status,
	}

	_, err = dao.CronTask.Ctx(ctx).Data(item).Where("cron_uid = ?", req.CronUid).Update()

	if err != nil {
		return
	}

	res = &v1.CronTaskItemRes{
		CronUid:  req.CronUid,
		Type:     item.Type,
		Content:  item.Content,
		CronExpr: item.CronExpr,
		Name:     item.Name,
		Updater:  req.Updater,
		Updated:  item.Updated.String(),
		Status:   item.Status,
	}

	return
}

func (self *sTaskCron) Query(ctx context.Context, req *v1.QueryCronTaskReq) (res *v1.QueryCronTaskRes, err error) {

	m := g.Map{}

	if req.Name != "" {
		m["name"] = req.Name

	}

	if req.CronUid != "" {
		m["cron_uid"] = req.CronUid
	}

	if req.Creater != "" {
		m["creater"] = req.Creater
	}

	if req.Type != "" {
		m["type"] = req.Type
	}

	list := make([]*entity.CronTask, 0)

	err = dao.CronTask.Ctx(ctx).Where(m).Scan(&list)

	if err != nil {
		return
	}

	res = new(v1.QueryCronTaskRes)

	for _, item := range list {
		res.List = append(res.List, &v1.CronTaskItemRes{
			Name:        item.Name,
			Type:        item.Type,
			Creater:     item.Creater,
			Content:     item.Content,
			CronExpr:    item.CronExpr,
			Status:      item.Status,
			Created:     item.Created.String(),
			CronUid:     item.CronUid,
			LastRunTime: item.LastrunTime.String(),
			NextRunTime: item.NextrunTime.String(),
			Updated:     item.Updated.String(),
			Updater:     item.Updater,
		})
	}

	return
}

func (self *sTaskCron) Delete(ctx context.Context, req *v1.DeleteCronTaskReq) (err error) {

	_, err = dao.TaskPreset.Ctx(ctx).WhereIn("cron_uid", req.CronUids).Delete()
	return
}
