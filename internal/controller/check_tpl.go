package controller

import (
	"context"
	v1 "go-ops/api/v1"
)

var CheckTpl *checktpl = new(checktpl)

type checktpl struct{}

func (self *checktpl) Create(ctx context.Context, req *v1.AddCheckTplReq) (res *v1.CheckTplItemRes, err error) {
	return
}

func (self *checktpl) Update(ctx context.Context, req *v1.UpdateCheckTplReq) (res *v1.CheckTplItemRes, err error) {
	return
}

func (self *checktpl) Query(ctx context.Context, req *v1.QueryCheckTplReq) (res *v1.CheckTplItemRes, err error) {
	return
}

func (self *checktpl) Delete(ctx context.Context, req *v1.DeleteCheckTplReq) (res v1.DeleteRes, err error) {
	return
}
