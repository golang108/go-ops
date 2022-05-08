package service

import (
	"context"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"
	v1 "go-ops/pkg/api/v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

type (
	sTaskPreset struct{}
)

var (
	insTaskPreset = sTaskPreset{}
)

func TaskPreset() *sTaskPreset {
	return &insTaskPreset
}

func (self *sTaskPreset) Create(ctx context.Context, req *v1.AddTaskPresetReq) (res *v1.TaskPresetItemRes, err error) {

	item := &entity.TaskPreset{
		Created: gtime.Now(),
		Name:    req.Name,
		Content: req.Content,
		Creater: req.Creater,
		Type:    req.Type,
		Uuid:    guid.S(),
	}

	_, err = dao.TaskPreset.Ctx(ctx).Data(item).Insert()

	if err != nil {
		return
	}

	res = &v1.TaskPresetItemRes{
		Name:    req.Name,
		Content: req.Content,
		Creater: req.Creater,
		Type:    req.Type,
		Uuid:    item.Uuid,
		Created: item.Created.String(),
	}

	return
}

func (self *sTaskPreset) Update(ctx context.Context, req *v1.UpdateTaskPresetReq) (res *v1.TaskPresetItemRes, err error) {

	var item *entity.TaskPreset

	err = dao.TaskPreset.Ctx(ctx).Where("uuid = ?", req.Uuid).Scan(&item)

	if item == nil {
		return
	}

	if req.Name != "" {
		item.Name = req.Name
	}

	if req.Content != "" {
		item.Content = req.Content
	}

	if req.Updater != "" {
		item.Updater = req.Updater
	}

	if req.Type != "" {
		item.Type = req.Type
	}

	_, err = dao.TaskPreset.Ctx(ctx).Data(item).Where("uuid = ?", req.Uuid).Update()

	if err != nil {
		return
	}

	res = &v1.TaskPresetItemRes{
		Name:    req.Name,
		Content: req.Content,
		Updater: req.Updater,
		Type:    req.Type,
		Uuid:    item.Uuid,
		Updated: item.Updated.String(),
	}

	return
}

func (self *sTaskPreset) Query(ctx context.Context, req *v1.QueryTaskPresetReq) (res *v1.QueryTaskPresetRes, err error) {

	m := g.Map{}

	if req.Name != "" {
		m["name"] = req.Name

	}

	if req.Uuid != "" {
		m["uuid"] = req.Uuid
	}

	if req.Creater != "" {
		m["creater"] = req.Creater
	}

	list := make([]*entity.TaskPreset, 0)

	err = dao.TaskPreset.Ctx(ctx).Where(m).Page(req.PageNum, req.PageSize).Scan(&list)

	if err != nil {
		return
	}

	count, err := dao.TaskPreset.Ctx(ctx).Where(m).Count()
	if err != nil {
		return
	}

	res = new(v1.QueryTaskPresetRes)
	res.Total = count
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize

	if res.Total%res.PageSize > 0 {
		res.PageTotal = 1
	}
	res.PageTotal += res.Total / res.PageSize

	for _, item := range list {
		res.List = append(res.List, &v1.TaskPresetItemRes{
			Name:    item.Name,
			Content: item.Content,
			Updater: item.Updater,
			Creater: item.Creater,
			Type:    item.Type,
			Uuid:    item.Uuid,
			Updated: item.Updated.String(),
			Created: item.Created.String(),
		})
	}

	return
}

func (self *sTaskPreset) Delete(ctx context.Context, req *v1.DeleteTaskPresetReq) (err error) {

	_, err = dao.TaskPreset.Ctx(ctx).WhereIn("uuid", req.Uuids).Delete()
	return
}
