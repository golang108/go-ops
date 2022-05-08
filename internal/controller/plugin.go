package controller

import (
	"context"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"
)

var Plugin *plugin = new(plugin)

type plugin struct{}

func (self *plugin) Create(ctx context.Context, req *v1.AddPluginReq) (res *v1.PluginItemRes, err error) {
	return service.Plugin().Create(ctx, req)
}

func (self *plugin) Update(ctx context.Context, req *v1.UpdatePluginReq) (res *v1.PluginItemRes, err error) {
	return service.Plugin().Update(ctx, req)
}

func (self *plugin) Query(ctx context.Context, req *v1.QueryPluginReq) (res *v1.QueryPluginRes, err error) {
	return service.Plugin().Query(ctx, req)
}

func (self *plugin) Delete(ctx context.Context, req *v1.DeletePluginReq) (res v1.DeleteRes, err error) {
	err = service.Plugin().Delete(ctx, req)
	if err != nil {
		res = v1.DeleteRes("删除失败")
		return
	}
	res = v1.DeleteRes("删除成功")
	return
}
