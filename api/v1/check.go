package v1

import "github.com/gogf/gf/v2/frame/g"

type AddCheckItemReq struct {
	g.Meta  `path:"/check/item/add" tags:"巡检管理" method:"post" summary:"添加巡检项"`
	Name    string `json:"name"  dc:"检查项名称"      ` // 检查项名称
	Desc    string `json:"desc"  dc:"检查项描述"      ` // 检查项描述
	Type    string `json:"type"  dc:"类型"      `    //
	Content string `json:"content" dc:"内容"    `    // 检查项内容
}

type CheckItemRes struct {
	CheckItemId string `json:"checkItemId" dc:"检查项id"`
	Name        string `json:"name"  dc:"检查项名称"      ` // 检查项名称
	Desc        string `json:"desc"  dc:"检查项描述"      ` // 检查项描述
	Type        string `json:"type"  dc:"类型"      `    //
	Content     string `json:"content" dc:"内容"    `    // 检查项内容
}

type UpdateCheckItemReq struct {
	g.Meta `path:"/check/item/update" tags:"巡检管理" method:"post" summary:"更新检查项"`
	CheckItemRes
}

type QueryCheckItemReq struct {
	g.Meta `path:"/check/item/query" tags:"巡检管理" method:"post" summary:"查询检查项"`
	Name   string `json:"name"  dc:"检查项名称"      ` // 检查项名称
	Type   string `json:"type"  dc:"类型"      `    //
	PageReq
}

type QueryCheckItemRes struct {
	Page
	List []*CheckItemRes `json:"list"`
}
