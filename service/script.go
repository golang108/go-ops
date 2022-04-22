package service

import (
	"context"
	"encoding/json"
	v1 "go-ops/api/v1"
	"go-ops/model/entity"
	"go-ops/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
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

	argsBytes, err := json.Marshal(req.Args)
	if err != nil {
		return
	}

	script := &entity.Script{
		ScriptUid: scriptUid,
		Name:      req.Name,
		Content:   req.Content,
		Args:      string(argsBytes),
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

	m := g.Map{}

	if req.Name != "" {
		m["name"] = req.Name

	}

	if req.Type != "" {
		m["type"] = req.Type
	}

	list := make([]*entity.Script, 0)

	err = dao.Script.Ctx(ctx).Where(m).Scan(&list)

	if err != nil {
		return
	}

	res = &v1.ScriptInfoRes{
		List: list,
	}
	return

	return
}

func (self *sScript) Update(ctx context.Context, req *v1.UpdateScriptReq) (res *v1.ScriptItemRes, err error) {

	var script *entity.Script

	err = dao.Script.Ctx(ctx).Where("script_uid = ?", req.ScriptId).Scan(&script)

	if script == nil {
		return
	}

	if len(req.Args) != 0 {
		argsBytes, err := json.Marshal(req.Args)
		if err != nil {
			return nil, err
		}
		script.Args = string(argsBytes)
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
		Args:     req.Args,
		Desc:     script.Desc,
		Type:     script.Type,
	}

	return
}

func (self *sScript) Delete(ctx context.Context, req *v1.DeleteScriptReq) (err error) {
	_, err = dao.CronTask.Ctx(ctx).WhereIn("script_uid", req.ScriptIds).Delete()
	return
}
