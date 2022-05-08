package controller

import (
	"context"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"
)

var CheckItem *checkitem = new(checkitem)

type checkitem struct{}

func (self *checkitem) Create(ctx context.Context, req *v1.AddCheckItemReq) (res *v1.CheckItemRes, err error) {
	return service.Checkitem().Create(ctx, req)
}

func (self *checkitem) Update(ctx context.Context, req *v1.UpdateCheckItemReq) (res *v1.CheckItemRes, err error) {
	return service.Checkitem().Update(ctx, req)
}

func (self *checkitem) Query(ctx context.Context, req *v1.QueryCheckItemReq) (res *v1.QueryCheckItemRes, err error) {
	return service.Checkitem().Query(ctx, req)
}

func (self *checkitem) Delete(ctx context.Context, req *v1.DeleteCheckItemReq) (res v1.DeleteRes, err error) {
	err = service.Checkitem().Delete(ctx, req)
	if err != nil {
		res = v1.DeleteRes("删除失败")
		return
	}
	res = v1.DeleteRes("删除成功")
	return
}
