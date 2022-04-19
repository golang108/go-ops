// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentInfo is the golang structure of table agent_info for DAO operations like Where/Data.
type AgentInfo struct {
	g.Meta         `orm:"table:agent_info, do:true"`
	Id             interface{} //
	Peerid         interface{} // 节点id
	Name           interface{} // agent名称
	ExpectedStatus interface{} // 期望状态
	Status         interface{} //
	IsDefault      interface{} //
	Timeout        interface{} // 启动超时时间
	Created        *gtime.Time //
	Updated        *gtime.Time //
	Version        interface{} // 版本信息
}
