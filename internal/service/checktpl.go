package service

import (
	"context"
	"errors"
	v1 "go-ops/api/v1"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/guid"
)

type sCheckTpl struct {
}

var (
	insCheckTpl = sCheckTpl{}
)

func CheckTpl() *sCheckTpl {
	return &insCheckTpl
}

func (selef *sCheckTpl) Create(ctx context.Context, req *v1.AddCheckTplReq) (err error) {

	checktpl := &entity.CheckTpl{
		Tid:         guid.S(),
		Name:        req.Name,
		Description: req.Desc,
		Type:        req.Type,
	}

	checktplDetails := make([]*entity.CheckTplDetail, 0)

	for _, item := range req.Items {

		checktplDetails = append(checktplDetails, &entity.CheckTplDetail{
			Tid:    checktpl.Tid,
			Cid:    item.ItemId,
			Weight: item.Weight,
		})

	}

	err = dao.CheckTpl.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.CheckTpl.Ctx(ctx).Data(checktpl).Insert()

		if err != nil {
			return err
		}

		_, err = dao.CheckTplDetail.Ctx(ctx).Data(checktplDetails).Insert()
		return err
	})
	return
}

func (self *sCheckTpl) Update(ctx context.Context, req *v1.UpdateCheckTplReq) (err error) {
	var checkTpl *entity.CheckTpl

	err = dao.CheckTpl.Ctx(ctx).Where("tid = ?", req.Tid).Scan(&checkTpl)
	if err != nil {
		return
	}

	if checkTpl == nil {
		err = errors.New("not found:" + req.Tid)
		return
	}

	checktplDetails := make([]*entity.CheckTplDetail, 0)

	for _, item := range req.Items {

		checktplDetails = append(checktplDetails, &entity.CheckTplDetail{
			Tid:    checkTpl.Tid,
			Cid:    item.ItemId,
			Weight: item.Weight,
		})

	}

	dao.CheckTplDetail.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.CheckTplDetail.Ctx(ctx).Where("tid = ?", req.Tid).Delete()
		if err != nil {
			return err
		}

		_, err = dao.CheckTplDetail.Ctx(ctx).Data(checktplDetails).Insert()
		if err != nil {
			return err
		}
		// _, err = dao.CheckTpl.Ctx(ctx).Data(checktplDetails).Insert()
		// if err != nil {
		// 	return err
		// }

		return err
	})

}
