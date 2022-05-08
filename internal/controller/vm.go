package controller

import (
	"context"
	"go-ops/internal/model"
	"go-ops/internal/peer"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"
	"strings"
	"time"
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
	res, err = service.VM().Query(ctx, req)
	if err != nil {
		return
	}
	if len(res.List) > 0 {
		return
	}

	req.Hostname = strings.TrimSpace(req.Hostname)

	if req.Hostname == "" {
		return
	}

	// 没有查找到主机信息，需要到网络中查找

	err = peer.SendMsgBroadCast(peer.GetOspPeer().PNet, &model.GetPeerInfo{HostName: req.Hostname})
	if err != nil {
		return
	}

	for i := 0; i < 30; i++ {
		res, err = service.VM().Query(ctx, req)
		if err != nil {
			return
		}
		if len(res.List) > 0 {
			return
		}
		time.Sleep(time.Second * 2)
	}
	return
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
