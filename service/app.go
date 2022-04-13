package service

import (
	"context"
	v1 "osp/api/v1"
	"osp/model/entity"
	"osp/service/internal/dao"

	"github.com/gogf/gf/v2/util/guid"
)

type (
	sApp struct{}
)

var (
	insApp = sApp{}
)

func App() *sApp {
	return &insApp
}

func (self *sApp) Create(ctx context.Context, req *v1.AddAppReq) (res *v1.AddAppRes, err error) {

	appid := guid.S()
	apikey := guid.S() + guid.S()
	seckey := guid.S() + guid.S() + guid.S() + guid.S()
	app := &entity.App{
		Name:   req.Name,
		Owner:  req.Owner,
		Appid:  appid,
		ApiKey: apikey,
		SecKey: seckey,
	}

	_, err = dao.App.Ctx(ctx).Data(app).Insert()
	if err != nil {
		return
	}

	res = &v1.AddAppRes{
		Appid:  appid,
		SecKey: seckey,
		ApiKey: apikey,
		Name:   req.Name,
		Owner:  req.Owner,
	}
	return
}
