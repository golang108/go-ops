package v1

import (
	"go-ops/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AddAppReq struct {
	g.Meta `path:"/v1/m/app" tags:"App管理" method:"post" summary:"创建一个app"`
	Name   string `json:"name" dc:"应用名"`  // 应用名
	Owner  string `json:"owner" dc:"拥有者"` // 拥有者
}

type AddAppRes struct {
	Appid    string `json:"appid"  dc:"appid"  `       //
	ApiKey   string `json:"apiKey"   `                 //
	SecKey   string `json:"secKey"   `                 //
	Owner    string `json:"owner"  dc:"拥有者"  `         //
	Name     string `json:"name"   dc:"应用名"  `         // 应用名
	Status   int    `json:"status" dc:"状态 1启用 0 禁用"  ` // 1启用 0 禁用
	OwnerUid string `json:"ownerUid" dc:"拥有者uid" `     // 拥有者uid
}

type UpdateAppReq struct {
	g.Meta `path:"/v1/m/app" tags:"App管理" method:"put" summary:"更新一个app"`
	Appid  string `json:"appid"  dc:"appid"  `       //
	Name   string `json:"name" dc:"应用名"`             // 应用名
	Owner  string `json:"owner" dc:"拥有者"`            // 拥有者
	Status int    `json:"status" dc:"状态 1启用 0 禁用"  ` // 1启用 0 禁用
}

type QueryAppReq struct {
	g.Meta `path:"/v1/m/app/query" tags:"App管理" method:"post" summary:"查询app"`
	Name   string `json:"name" dc:"应用名"`  // 应用名
	Owner  string `json:"owner" dc:"拥有者"` // 拥有者
	PageReq
}

type QueryAppRes struct {
	Page
	List []*entity.App `json:"list"`
}

type DeleteAppReq struct {
	g.Meta `path:"/v1/m/app/delete" tags:"App管理" method:"post" summary:"删除app"`
	Appids []string `json:"appids" dc:"app id 列表"`
}
