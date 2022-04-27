package service

import (
	"context"

	"go-ops/internal/model"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"

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
			PublicIp: req.PeerId,
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
