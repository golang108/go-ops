package service

import (
	"context"

	v1 "go-ops/api/v1"
	"go-ops/internal/model"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

type (
	sVM struct{}
)

var (
	insVM = sVM{}
)

func VM() *sVM {
	return &insVM
}

func (self *sVM) Check(ctx context.Context, req *model.PeerInfo) (err error) {
	var vm *entity.Vm

	err = dao.Vm.Ctx(ctx).Where("hostname = ?", req.HostName).Where("public_ip = ?", req.PublicIp).Scan(&vm)
	if err != nil {
		return
	}

	if vm == nil {

		uuid := guid.S()
		vm = &entity.Vm{
			PeerId:   req.PeerId,
			OsInfo:   req.Os,
			PublicIp: req.PublicIp,
			Hostname: req.HostName,
			Name:     req.Name,
			Uuid:     uuid,
			Address:  req.Address,
			Creater:  "go-ops",
		}
		_, err = dao.Vm.Ctx(ctx).Data(vm).Insert()
		return
	}

	if vm.PeerId == req.PeerId && vm.Address == req.Address {
		return
	}

	vm.PeerId = req.PeerId
	vm.Address = req.Address

	_, err = dao.Vm.Ctx(ctx).Data(vm).Where("id = ?", vm.Id).Update()

	return

}

func (self *sVM) Create(ctx context.Context, req *v1.AddVmReq) (res *v1.VmItemRes, err error) {

	vm := &entity.Vm{
		Uuid:     req.Uuid,
		Name:     req.Name,
		Hostname: req.Hostname,
		Creater:  req.Creater,
		OsInfo:   req.Os,
		PublicIp: req.PublicIp,
		Created:  gtime.Now(),
	}

	_, err = dao.Vm.Ctx(ctx).Data(vm).Insert()
	if err != nil {
		return
	}

	res = &v1.VmItemRes{
		Uuid:     vm.Uuid,
		Name:     req.Name,
		Hostname: req.Hostname,
		Creater:  req.Creater,
		Os:       req.Os,
		PublicIp: req.PublicIp,
	}

	return
}

func (self *sVM) Update(ctx context.Context, req *v1.UpdateVmReq) (res *v1.VmItemRes, err error) {

	var vm *entity.Vm

	err = dao.Vm.Ctx(ctx).Where("uuid = ?", req.Uuid).Scan(&vm)
	if err != nil {
		return
	}

	if vm == nil {
		return
	}

	if req.Name != "" {
		vm.Name = req.Name
	}

	if req.Os != "" {
		vm.OsInfo = req.Os
	}

	if req.Uuid != "" {
		vm.Uuid = req.Uuid
	}

	_, err = dao.Vm.Ctx(ctx).Data(vm).Where("id = ?", vm.Id).Update()

	if err != nil {
		return
	}

	res = &v1.VmItemRes{
		Name:     vm.Name,
		Hostname: vm.Hostname,
		Os:       vm.OsInfo,
		PublicIp: vm.PublicIp,
		Uuid:     vm.Uuid,
		PeerId:   vm.PeerId,
	}

	return
}

func (self *sVM) Query(ctx context.Context, req *v1.QueryVmReq) (res *v1.QueryVmRes, err error) {

	m := g.Map{}

	if req.Name != "" {
		m["name"] = req.Name

	}

	if req.Hostname != "" {
		m["hostname"] = req.Hostname
	}

	if req.Uuid != "" {
		m["uuid"] = req.Uuid
	}

	list := make([]*entity.Vm, 0)

	err = dao.Vm.Ctx(ctx).Where(m).Page(req.PageNum, req.PageSize).Scan(&list)

	if err != nil {
		return
	}

	count, err := dao.Vm.Ctx(ctx).Where(m).Count()
	if err != nil {
		return
	}

	res = &v1.QueryVmRes{
		List: list,
	}

	res.Total = count
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize

	if res.Total%res.PageSize > 0 {
		res.PageTotal = 1
	}
	res.PageTotal += res.Total / res.PageSize
	return
}

func (self *sVM) Delete(ctx context.Context, req *v1.DeleteVmReq) (err error) {
	_, err = dao.Vm.Ctx(ctx).WhereIn("uuid", req.Uuids).Delete()
	return
}
