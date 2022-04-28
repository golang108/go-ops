package controller

import (
	"context"
	v1 "go-ops/api/v1"
	"go-ops/internal/service"
)

var VM *vm = new(vm)

type vm struct{}

func (self *vm) Create(ctx context.Context, req *v1.AddVmReq) (res *v1.VmItemRes, err error) {
	return service.VM().Create(ctx, req)
}

func (self *vm) Update(ctx context.Context, req *v1.UpdateVmReq) (res *v1.VmItemRes, err error) {
	return service.VM().Update(ctx, req)
}

func (self *vm) Query(ctx context.Context, req *v1.QueryVmReq) (res *v1.QueryVmRes, err error) {
	return service.VM().Query(ctx, req)
}

func (self *vm) Delete(ctx context.Context, req *v1.DeleteVmReq) (res v1.DeleteRes, err error) {
	err = service.VM().Delete(ctx, req)
	if err != nil {
		res = v1.DeleteRes("删除失败")
		return
	}
	res = v1.DeleteRes("删除成功")
	return
}
