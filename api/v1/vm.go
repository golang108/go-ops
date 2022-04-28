package v1

import (
	"go-ops/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AddVmReq struct {
	g.Meta   `path:"/v1/m/vm/add" tags:"主机节点" method:"post" summary:"添加主机信息"`
	Name     string `json:"name"`
	Hostname string `json:"hostname" dc:"主机hostname"`
	PublicIp string `json:"publicIp"`
	Uuid     string `json:"uuid"`
	Os       string `json:"os"`
	Creater  string `json:"creater"`
}

type VmItemRes struct {
	PeerId   string `json:"peerId"`
	Name     string `json:"name"`
	Hostname string `json:"hostname" dc:"主机hostname"`
	PublicIp string `json:"publicIp"`
	Uuid     string `json:"uuid"`
	Os       string `json:"os"`
	Creater  string `json:"creater"`
	Updater  string `json:"updater"`
}

type UpdateVmReq struct {
	g.Meta  `path:"/v1/m/vm/update" tags:"主机节点" method:"post" summary:"更新主机信息"`
	Name    string `json:"name"`
	Uuid    string `json:"uuid"`
	Os      string `json:"os"`
	Updater string `json:"updater"`
}

type DeleteVmReq struct {
	g.Meta `path:"/v1/m/vm/delete" tags:"主机节点" method:"post" summary:"删除主机信息"`
	Uuids  string `json:"uuids"`
}

type QueryVmReq struct {
	g.Meta   `path:"/v1/m/vm/query" tags:"主机节点" method:"post" summary:"查询主机节点"`
	Name     string `json:"name"`
	Uuid     string `json:"uuid"`
	Hostname string `json:"hostname"`
	PageReq
}

type QueryVmRes struct {
	Page
	List []*entity.Vm `json:"list"`
}
