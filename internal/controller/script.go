package controller

import (
	"context"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"

	"github.com/gogf/gf/v2/os/glog"
)

var Script *script = new(script)

type script struct{}

func (self *script) Create(ctx context.Context, req *v1.AddScriptReq) (res *v1.ScriptItemRes, err error) {

	res, err = service.Script().Create(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "create script err:%v", err)
	}
	return
}

func (self *script) Update(ctx context.Context, req *v1.UpdateScriptReq) (res *v1.ScriptItemRes, err error) {
	res, err = service.Script().Update(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "update script err:%v", err)
	}
	return
}

func (self *script) Query(ctx context.Context, req *v1.ScriptQueryReq) (res *v1.ScriptInfoRes, err error) {
	res, err = service.Script().Query(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "query script err:%v", err)
	}
	return
}

func (self *script) Delete(ctx context.Context, req *v1.DeleteScriptReq) (res v1.DeleteScriptRes, err error) {
	err = service.Script().Delete(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "delete script err:%v", err)
		res = v1.DeleteScriptRes("删除失败")
		return
	}
	res = v1.DeleteScriptRes("删除成功")
	return
}
