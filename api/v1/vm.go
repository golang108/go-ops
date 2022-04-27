package v1

import "github.com/gogf/gf/v2/frame/g"

type AddVmReq struct {
	g.Meta   `path:"/v1/m/vm/add" tags:"主机节点" method:"put" summary:"添加主机信息"`
	Name     string `json:"name"`
	Hostname string `json:"hostname" dc:"主机hostname"`
	PublicIp string `json:"publicIp"`
	Uuid     string `json:"uuid"`
	Os       string `json:"os"`
	Creater  string `json:"creater"`
}
