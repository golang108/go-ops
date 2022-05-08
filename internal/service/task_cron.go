package service

import (
	"context"
	"errors"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"
	v1 "go-ops/pkg/api/v1"

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

	var item *entity.CronTask

	err = dao.CronTask.Ctx(ctx).Where("cron_uid = ?", req.CronUid).Scan(&item)

	if item == nil {

		err = errors.New("not found:" + req.CronUid)
		return
	}

	if req.Type != "" {
		item.Type = req.Type
	}

	if req.Content != "" {
		item.Content = req.Content
	}

	if req.CronExpr != "" {
		item.CronExpr = req.CronExpr
	}

	if req.Name != "" {
		item.Name = req.Name
	}

	if req.Updater != "" {
		item.Updater = req.Updater
	}

	if req.Status != "" {
		item.Status = req.Status
	}

	item.Updated = gtime.Now()

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

	err = dao.CronTask.Ctx(ctx).Where(m).Page(req.PageNum, req.PageSize).Scan(&list)

	if err != nil {
		return
	}

	res = new(v1.QueryCronTaskRes)

	count, err := dao.CronTask.Ctx(ctx).Where(m).Count()
	if err != nil {
		return
	}

	res.Total = count
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize

	if res.Total%res.PageSize > 0 {
		res.PageTotal = 1
	}
	res.PageTotal += res.Total / res.PageSize

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

	_, err = dao.CronTask.Ctx(ctx).WhereIn("cron_uid", req.CronUids).Delete()
	return
}
