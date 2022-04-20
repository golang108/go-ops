package service

import (
	"context"
	v1 "go-ops/api/v1"
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

	return
}

func (self *sCheckitem) Update(ctx context.Context, req *v1.UpdateCheckItemReq) (res *v1.CheckItemRes, err error) {

	return
}

func (self *sCheckitem) Query(ctx context.Context, req *v1.QueryCheckItemReq) (res *v1.QueryCheckItemRes, err error) {

	return
}
