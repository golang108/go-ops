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

type DeleteCheckItemReq struct {
	g.Meta       `path:"/check/item/delete" tags:"巡检管理" method:"post" summary:"删除检查项"`
	CheckItemIds []string `json:"checkItemIds" dc:"检查项id列表"`
}

type CheckTplItemInfo struct {
	ItemId string  `json"itemId" dc:"巡检项id"`
	Weight float64 `json:"weight" dc:"权重"`
}

type CheckTplItem struct {
	Tid   string              `json:"tid"`
	Name  string              `json:"name"  dc:"检查项名称"      ` // 检查项名称
	Desc  string              `json:"desc"  dc:"检查项描述"      ` // 检查项描述
	Type  string              `json:"type"  dc:"类型"      `    //
	Items []*CheckTplItemInfo `json:"items" dc:"items"`       //
}

type CheckTplItemRes *CheckTplItem

type AddCheckTplReq struct {
	g.Meta `path:"/check/tpl/create" tags:"巡检管理" method:"post" summary:"添加巡检模版"`
	CheckTplItem
}

type UpdateCheckTplReq struct {
	g.Meta `path:"/check/tpl/update" tags:"巡检管理" method:"post" summary:"更新巡检模版"`
	CheckTplItem
}

type DeleteCheckTplReq struct {
	g.Meta `path:"/check/tpl/delete" tags:"巡检管理" method:"post" summary:"删除巡检模版"`
	Tids   []string `json:"tids"`
}

type QueryCheckTplReq struct {
	g.Meta `path:"/check/tpl/query" tags:"巡检管理" method:"post" summary:"查询巡检模版"`
	Tid    string `json:"tid"`
	Name   string `json:"name"  dc:"检查项名称"      ` // 检查项名称
	PageReq
}
