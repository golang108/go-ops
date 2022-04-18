package controller

import (
	"context"
	v1 "go-ops/api/v1"
	"go-ops/service"
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
