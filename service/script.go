package service

import (
	"context"
	v1 "osp/api/v1"
	"osp/model/entity"
	"osp/service/internal/dao"

	"github.com/gogf/gf/v2/util/guid"
)

type (
	sScript struct{}
)

var (
	insScript = sScript{}
)

func Script() *sScript {
	return &insScript
}

func (self *sScript) Create(ctx context.Context, req *v1.AddScriptReq) (res *v1.ScriptItemRes, err error) {

	scriptUid := guid.S()

	script := &entity.Script{
		ScriptUid: scriptUid,
		Name:      req.Name,
		Content:   req.Content,
		Args:      req.Args,
		Desc:      req.Desc,
		Type:      req.Type,
	}

	_, err = dao.Script.Ctx(ctx).Data(script).Insert()
	if err != nil {
		return
	}

	res = &v1.ScriptItemRes{
		ScriptId: script.ScriptUid,
	}

	return
}

func (self *sScript) Query(ctx context.Context, req *v1.ScriptQueryReq) (res *v1.ScriptInfoRes, err error) {

	//da := dao.Script.Ctx(ctx).Where("1 = 1")

	return
}

func (self *sScript) Update(ctx context.Context, req *v1.UpdateScriptReq) (res *v1.ScriptItemRes, err error) {

	var script *entity.Script

	err = dao.Script.Ctx(ctx).Where("script_uid = ?", req.ScriptId).Scan(&script)

	if script == nil {
		return
	}

	if req.Args != "" {
		script.Args = req.Args
	}

	if req.Content != "" {
		script.Content = req.Content
	}

	if req.Desc != "" {
		script.Desc = req.Desc
	}

	if req.Name != "" {
		script.Name = req.Name
	}

	if req.Type != "" {
		script.Type = req.Type
	}

	_, err = dao.Script.Ctx(ctx).Data(script).Where("script_uid = ?", req.ScriptId).Update()

	if err != nil {
		return
	}

	res = &v1.ScriptItemRes{
		ScriptId: req.ScriptId,
		Name:     script.Name,
		Content:  script.Content,
		Args:     script.Content,
		Desc:     script.Desc,
		Type:     script.Type,
	}

	return
}
