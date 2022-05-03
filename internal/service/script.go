package service

import (
	"context"
	"encoding/json"
	"errors"
	v1 "go-ops/api/v1"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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
		Created:   gtime.Now(),
		Cmd:       req.Cmd,
		WaitTime:  req.WaitTime,
	}

	_, err = dao.Script.Ctx(ctx).Data(script).Insert()
	if err != nil {
		return
	}

	res = &v1.ScriptItemRes{
		ScriptId: script.ScriptUid,
		Name:     script.Name,
		Content:  script.Content,
		Args:     req.Args,
		Desc:     script.Desc,
		Type:     script.Type,
		Cmd:      script.Cmd,
		WaitTime: script.WaitTime,
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

	if req.ScriptId != "" {
		m["script_uid"] = req.ScriptId
	}

	list := make([]*entity.Script, 0)

	err = dao.Script.Ctx(ctx).Where(m).Page(req.PageNum, req.PageSize).Scan(&list)

	if err != nil {
		return
	}

	count, err := dao.Script.Ctx(ctx).Where(m).Count()
	if err != nil {
		return
	}

	res = &v1.ScriptInfoRes{
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

func (self *sScript) Update(ctx context.Context, req *v1.UpdateScriptReq) (res *v1.ScriptItemRes, err error) {

	var script *entity.Script

	err = dao.Script.Ctx(ctx).Where("script_uid = ?", req.ScriptId).Scan(&script)

	if script == nil {
		err = errors.New("not found:" + req.ScriptId)
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

	if req.Cmd != "" {
		script.Cmd = req.Cmd
	}

	if req.WaitTime != 0 {
		script.WaitTime = req.WaitTime
	}

	script.Updated = gtime.Now()

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
		Cmd:      script.Cmd,
		WaitTime: script.WaitTime,
	}

	return
}

func (self *sScript) Delete(ctx context.Context, req *v1.DeleteScriptReq) (err error) {
	_, err = dao.Script.Ctx(ctx).WhereIn("script_uid", req.ScriptIds).Delete()
	return
}
