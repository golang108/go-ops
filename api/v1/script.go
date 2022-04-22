package v1

import (
	"go-ops/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AddScriptReq struct {
	g.Meta  `path:"/v1/m/script" tags:"脚本库管理" method:"post" summary:"创建一个脚本"`
	Name    string            `json:"name"    dc:"脚本名称"  `                  // 命令名称
	Content string            `json:"content" dc:"脚本内容"  `                  // 脚本内容
	Args    map[string]string `json:"args"    dc:"参数信息"  `                  // 参数信息
	Desc    string            `json:"desc"    dc:"描述信息"  `                  // 描述信息
	Type    string            `json:"type"   dc:"脚本类型shell或者powershell"   ` // 脚本类型shell或者powershell
	Creater string            `json:"creater" dc:"创建人"`
}

type ScriptQueryReq struct {
	g.Meta `path:"/v1/m/script/query" tags:"脚本库管理" method:"post" summary:"查询脚本库信息"`
	Name   string `json:"name"    dc:"脚本名称"  `
	Type   string `json:"type"   dc:"脚本类型shell或者powershell"`
}

type ScriptInfoRes struct {
	List []*entity.Script `json:"list"`
}

type UpdateScriptReq struct {
	g.Meta   `path:"/v1/m/script" tags:"脚本库管理" method:"put" summary:"更新脚本信息"`
	ScriptId string            `json:"scriptId" dc:"脚本id"`
	Name     string            `json:"name"    dc:"脚本名称"  `                  // 命令名称
	Content  string            `json:"content" dc:"脚本内容"  `                  // 脚本内容
	Args     map[string]string `json:"args"    dc:"参数信息"  `                  // 参数信息
	Desc     string            `json:"desc"    dc:"描述信息"  `                  // 描述信息
	Type     string            `json:"type"   dc:"脚本类型shell或者powershell"   ` // 脚本类型shell或者powershell
	Updater  string            `json:"updater" dc:"更新人"`
}

type ScriptItemRes struct {
	ScriptId string            `json:"scriptId" dc:"脚本id"`
	Name     string            `json:"name"    dc:"脚本名称"  `                  // 命令名称
	Content  string            `json:"content" dc:"脚本内容"  `                  // 脚本内容
	Args     map[string]string `json:"args"    dc:"参数信息"  `                  // 参数信息
	Desc     string            `json:"desc"    dc:"描述信息"  `                  // 描述信息
	Type     string            `json:"type"   dc:"脚本类型shell或者powershell"   ` // 脚本类型shell或者powershell
	Creater  string            `json:"creater" dc:"创建人"`
	Updater  string            `json:"updater" dc:"更新人"`
}

type DeleteScriptReq struct {
	g.Meta    `path:"/v1/m/script/delete" tags:"脚本库管理" method:"put" summary:"删除脚本信息"`
	ScriptIds string `json:"scriptIds" dc:"脚本id 列表"`
}

type DeleteScriptRes string
