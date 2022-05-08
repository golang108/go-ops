package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddAgentReq struct {
	g.Meta `path:"/peer/agent/add" tags:"Agent管理" method:"post" summary:"节点添加agent"`
	List   []*AgentInfo `json:"list"`
}

type AgentInfo struct {
	Peerid         string `json:"peerid"   dc:"节点id, 空表示默认版本,所有节点将使用这个版本"      ` // 节点id
	Name           string `json:"name"     dc:"agent 名称"      `                  // agent名称
	Version        string `json:"version" dc:"agent 版本"`
	ExpectedStatus string `json:"expectedStatus" dc:"期望状态 running,stopped,deleted" ` // 期望状态
	DownloadUrl    string `json:"downloadUrl" dc:"下载地址"`                             // 下载地址
	Status         string `json:"status"   dc:"agent 当前状态"      `                    // 当前agent状态
	IsDefault      int    `json:"isDefault"  dc:"是否默认安装 0-否  1-是"    `               //
	Timeout        int    `json:"timeout"    dc:"agent 启动超时时间"    `                  // 启动超时时间
}

type AgentInfoRes struct {
	List []*AgentInfo `json:"list"`
}

type UpdateAgentReq struct {
	g.Meta `path:"/peer/agent/update" tags:"Agent管理" method:"post" summary:"节点更新agent"`
	List   []*AgentInfo `json:"list"`
}

type QueryAgentStatusReq struct {
	g.Meta `path:"/peer/agent/status" tags:"Agent管理" method:"post" summary:"查询agent状态"`
	List   []*QueryAgentStatus `json:"list"`
}

type QueryAgentStatus struct {
	Peerid string   `json:"peerid" dc:"节点id"`
	Agents []string `json:"agents" dc:"需要查询的agent名称列表"`
}

type PeerAgentStatus struct {
	Peerid string `json:"peerid" dc:"节点id"`
	Agents []*AgentInfo
}

type QueryAgentStatusRes struct {
	List []*PeerAgentStatus
}
