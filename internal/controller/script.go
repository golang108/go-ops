package controller

import (
	"context"
	v1 "osp/api/v1"
)

var Script *script = new(script)

type script struct{}

func (self *script) Create(ctx context.Context, req *v1.AddScriptReq) (res *v1.ScriptItemRes, err error) {
	return
}

func (self *script) Update(ctx context.Context, req *v1.UpdateScriptReq) (res *v1.ScriptItemRes, err error) {
	return
}

func (self *script) Query(ctx context.Context, req *v1.ScriptQueryReq) (res *v1.ScriptInfoRes, err error) {
	return
}
