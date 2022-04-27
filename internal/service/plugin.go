package service

import (
	"context"
	v1 "go-ops/api/v1"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/util/guid"
)

type (
	sPlugin struct{}
)

var (
	insPlugin = sPlugin{}
)

func Plugin() *sPlugin {
	return &insPlugin
}

func (self *sPlugin) Create(ctx context.Context, req *v1.AddPluginReq) (res *v1.PluginItemRes, err error) {

	uuid := guid.S()
	plugin := &entity.Plugin{
		Name:        req.Name,
		Os:          req.Os,
		Arch:        req.Arch,
		PackageName: req.PackageName,
		Md5:         req.Md5,
		Creater:     req.Creater,
		Created:     gtime.New(),
		Uuid:        uuid,
	}

	_, err = dao.Plugin.Ctx(ctx).Data(plugin).Insert()
	if err != nil {
		return
	}

	res = &v1.PluginItemRes{
		Name:        req.Name,
		Os:          req.Os,
		Arch:        req.Arch,
		PackageName: req.PackageName,
		Md5:         req.Md5,
		Creater:     req.Creater,
		Created:     plugin.Created.String(),
		Uuid:        uuid,
	}
	return
}

func (self *sPlugin) Query(ctx context.Context, req *v1.QueryPluginReq) (res *v1.QueryPluginRes, err error) {

	m := g.Map{}

	if req.Name != "" {
		m["name"] = req.Name
	}

	if req.PackageName != "" {
		m["package_name"] = req.PackageName
	}

	if req.Os != "" {
		m["os"] = req.Os
	}
	if req.Arch != "" {
		m["arch"] = req.Arch
	}
	if req.Creater != "" {
		m["creater"] = req.Creater
	}
	if req.Updater != "" {
		m["updater"] = req.Updater
	}

	list := make([]*entity.Plugin, 0)

	err = dao.Plugin.Ctx(ctx).Where(m).Page(req.PageNum, req.PageSize).Scan(&list)

	if err != nil {
		return
	}

	count, err := dao.Plugin.Ctx(ctx).Where(m).Count()
	if err != nil {
		return
	}

	res = &v1.QueryPluginRes{
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

func (self *sPlugin) Update(ctx context.Context, req *v1.UpdatePluginReq) (res *v1.PluginItemRes, err error) {

	var plugin *entity.Plugin

	err = dao.Plugin.Ctx(ctx).Where("uuid = ?", req.Uuid).Scan(&plugin)
	if err != nil {
		return
	}

	if plugin == nil {
		return
	}

	if req.Name != "" {
		plugin.Name = req.Name
	}

	if req.Os != "" {
		plugin.Os = req.Os
	}

	if req.Arch != "" {
		plugin.Arch = req.Arch
	}

	if req.Md5 != "" {
		plugin.Md5 = req.Md5
	}

	if req.PackageName != "" {
		plugin.PackageName = req.PackageName
	}

	if req.Updater != "" {
		plugin.Updater = req.Updater
	}

	plugin.Updated = gtime.Now()

	_, err = dao.App.Ctx(ctx).Data(plugin).Where("uuid = ?", req.Uuid).Update()

	if err != nil {
		return
	}

	res = &v1.PluginItemRes{
		Name:        plugin.Name,
		Os:          plugin.Os,
		Arch:        plugin.Arch,
		PackageName: plugin.PackageName,
		Md5:         plugin.Md5,
		Creater:     plugin.Creater,
		Created:     plugin.Created.String(),
		Uuid:        plugin.Uuid,
		Updater:     plugin.Updater,
		Updated:     plugin.Updated.String(),
	}

	return
}

func (self *sPlugin) Delete(ctx context.Context, req *v1.DeletePluginReq) (err error) {
	_, err = dao.Plugin.Ctx(ctx).WhereIn("uuid", req.Uuids).Delete()
	return
}
