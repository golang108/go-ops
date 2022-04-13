package controller

import (
	"context"
	v1 "osp/api/v1"
	"osp/service"
)

var Script *script = new(script)

type script struct{}

func (self *script) Create(ctx context.Context, req *v1.AddScriptReq) (res *v1.ScriptItemRes, err error) {

	return service.Script().Create(ctx, req)
}

func (self *script) Update(ctx context.Context, req *v1.UpdateScriptReq) (res *v1.ScriptItemRes, err error) {
	return service.Script().Update(ctx, req)
}

func (self *script) Query(ctx context.Context, req *v1.ScriptQueryReq) (res *v1.ScriptInfoRes, err error) {
	return service.Script().Query(ctx, req)
}
