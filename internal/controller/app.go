package controller

import (
	"context"
	v1 "go-ops/api/v1"
	"go-ops/internal/service"
)

var App *app = new(app)

type app struct{}

func (self *app) Create(ctx context.Context, req *v1.AddAppReq) (res *v1.AddAppRes, err error) {
	return service.App().Create(ctx, req)
}

func (self *app) Update(ctx context.Context, req *v1.UpdateAppReq) (res *v1.AddAppRes, err error) {
	return service.App().Update(ctx, req)
}

func (self *app) Query(ctx context.Context, req *v1.QueryAppReq) (res *v1.QueryAppRes, err error) {
	return service.App().Query(ctx, req)
}

func (self *app) SingleQuery(ctx context.Context, req *v1.QuerySingleAppReq) (res *v1.AddAppRes, err error) {
	return service.App().SingleQuery(ctx, req)
}

func (self *app) Delete(ctx context.Context, req *v1.DeleteAppReq) (res v1.DeleteRes, err error) {
	err = service.App().Delete(ctx, req)
	if err != nil {
		res = v1.DeleteRes("删除失败")
		return
	}
	res = v1.DeleteRes("删除成功")
	return
}
