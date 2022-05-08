package service

import (
	"context"
	"errors"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"
	v1 "go-ops/pkg/api/v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

type (
	sCheckitem struct{}
)

var (
	insCheckitem = sCheckitem{}
)

func Checkitem() *sCheckitem {
	return &insCheckitem
}

func (self *sCheckitem) Create(ctx context.Context, req *v1.AddCheckItemReq) (res *v1.CheckItemRes, err error) {

	item := &entity.CheckItem{
		CheckItemId: guid.S(),
		Name:        req.Name,
		Content:     req.Content,
		Desc:        req.Content,
		Type:        req.Type,
		WaitTime:    req.WaitTime,
		Cmd:         req.Cmd,
		Weight:      req.Weight,
	}

	_, err = dao.CheckItem.Ctx(ctx).Data(item).Insert()
	if err != nil {
		return
	}
	res = &v1.CheckItemRes{
		CheckItemId: item.CheckItemId,
		Name:        item.Name,
		Content:     item.Content,
		Desc:        item.Desc,
		Type:        item.Type,
		Cmd:         item.Cmd,
		WaitTime:    item.WaitTime,
		Weight:      item.Weight,
	}

	return
}

func (self *sCheckitem) Update(ctx context.Context, req *v1.UpdateCheckItemReq) (res *v1.CheckItemRes, err error) {

	var item *entity.CheckItem

	err = dao.CheckItem.Ctx(ctx).Where("check_item_id = ?", req.CheckItemId).Scan(&item)

	if err != nil {
		return
	}

	if item == nil {
		err = errors.New("not found " + req.CheckItemId)
		return
	}

	if req.Content != "" {
		item.Content = req.Content
	}

	if req.Name != "" {
		item.Name = req.Name
	}

	if req.Desc != "" {
		item.Desc = req.Desc
	}

	if req.Type != "" {
		item.Type = req.Type
	}
	if req.Cmd != "" {
		item.Cmd = req.Cmd
	}
	if req.Weight != 0 {
		item.Weight = req.Weight
	}
	if req.WaitTime != 0 {
		item.WaitTime = req.WaitTime
	}

	_, err = dao.CheckItem.Ctx(ctx).Data(item).Where("check_item_id = ?", req.CheckItemId).Update()
	if err != nil {
		return
	}

	res = &v1.CheckItemRes{
		CheckItemId: item.CheckItemId,
		Name:        item.Name,
		Content:     item.Content,
		Desc:        item.Desc,
		Type:        item.Type,
		Cmd:         item.CheckItemId,
		WaitTime:    item.WaitTime,
		Weight:      item.Weight,
	}
	return
}

func (self *sCheckitem) Query(ctx context.Context, req *v1.QueryCheckItemReq) (res *v1.QueryCheckItemRes, err error) {

	m := g.Map{}

	if req.CheckItemId != "" {
		m["check_item_id"] = req.CheckItemId
	}

	if req.Name != "" {
		m["name"] = req.Name

	}

	if req.Type != "" {
		m["type"] = req.Type
	}

	list := make([]*entity.CheckItem, 0)

	err = dao.CheckItem.Ctx(ctx).Where(m).Page(req.PageNum, req.PageSize).Scan(&list)
	if err != nil {
		return
	}

	count, err := dao.CheckItem.Ctx(ctx).Where(m).Count()
	if err != nil {
		return
	}

	res = &v1.QueryCheckItemRes{
		List: make([]*v1.CheckItemRes, 0),
	}

	for _, item := range list {
		res.List = append(res.List, &v1.CheckItemRes{
			CheckItemId: item.CheckItemId,
			Name:        item.Name,
			Content:     item.Content,
			Desc:        item.Desc,
			Type:        item.Type,
			Cmd:         item.Cmd,
			Weight:      item.Weight,
			WaitTime:    item.WaitTime,
		})
	}

	res.Total = count
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize

	if res.Total%res.PageSize > 0 {
		res.PageTotal = 1
	}
	res.PageTotal += res.Total / res.PageSize
	return
}

func (self *sCheckitem) Delete(ctx context.Context, req *v1.DeleteCheckItemReq) (err error) {
	_, err = dao.CheckItem.Ctx(ctx).WhereIn("check_item_id", req.CheckItemIds).Delete()
	return
}
