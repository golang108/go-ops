package controller

import (
	"context"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"

	"github.com/gogf/gf/v2/os/glog"
)

var CheckItem *checkitem = new(checkitem)

type checkitem struct{}

func (self *checkitem) Create(ctx context.Context, req *v1.AddCheckItemReq) (res *v1.CheckItemRes, err error) {
	res, err = service.Checkitem().Create(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "create checkitem err:%v", err)
		return
	}
	return
}

func (self *checkitem) Update(ctx context.Context, req *v1.UpdateCheckItemReq) (res *v1.CheckItemRes, err error) {
	res, err = service.Checkitem().Update(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "update checkitem err:%v", err)
		return
	}
	return
}

func (self *checkitem) Query(ctx context.Context, req *v1.QueryCheckItemReq) (res *v1.QueryCheckItemRes, err error) {
	res, err = service.Checkitem().Query(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "query checkitem err:%v", err)
		return
	}
	return
}

func (self *checkitem) Delete(ctx context.Context, req *v1.DeleteCheckItemReq) (res v1.DeleteRes, err error) {
	err = service.Checkitem().Delete(ctx, req)
	if err != nil {
		glog.Errorf(ctx, "delete checkitem err:%v", err)
		res = v1.DeleteRes("删除失败")
		return
	}
	res = v1.DeleteRes("删除成功")
	return
}
