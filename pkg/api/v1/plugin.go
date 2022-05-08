package v1

import (
	"go-ops/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AddPluginReq struct {
	g.Meta      `path:"/v1/m/plugin/create" tags:"插件" method:"post" summary:"添加插件"`
	Name        string `json:"name"   dc:"插件名"     `     // 插件名
	PackageName string `json:"packageName"  dc:"包名"`     // 包名
	Os          string `json:"os"       dc:"操作系统"   `    // 操作系统
	Arch        string `json:"arch"    dc:"架构"     `     // 架构
	Md5         string `json:"md5"      dc:"包md5名称"    ` // 包md5名称
	Creater     string `json:"creater" dc:"创建人"`
}

type PluginItemRes struct {
	Uuid        string `json:"uuid" dc:"插件uuid"`
	Name        string `json:"name"   dc:"插件名"     `     // 插件名
	PackageName string `json:"packageName"  dc:"包名"`     // 包名
	Os          string `json:"os"       dc:"操作系统"   `    // 操作系统
	Arch        string `json:"arch"    dc:"架构"     `     // 架构
	Md5         string `json:"md5"      dc:"包md5名称"    ` // 包md5名称
	Creater     string `json:"creater" dc:"创建人"`
	Updater     string `json:"updater" dc:"更新人"`
	Created     string `json:"created" dc:"创建时间"`
	Updated     string `json:"updated" dc:"更新时间"`
}

type UpdatePluginReq struct {
	g.Meta `path:"/v1/m/plugin/update" tags:"插件" method:"post" summary:"更新插件"`
	PluginItemRes
}

type DeletePluginReq struct {
	g.Meta `path:"/v1/m/plugin/delete" tags:"插件" method:"post" summary:"删除插件"`
	Uuids  []string `json:"uuids" dc:"插件uuid列表"`
}

type QueryPluginReq struct {
	g.Meta `path:"/v1/m/plugin/query" tags:"插件" method:"post" summary:"查询插件"`
	PageReq
	PluginItemRes
}

type QueryPluginRes struct {
	Page
	List []*entity.Plugin `json:"list"`
}
