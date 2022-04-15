package model

type AgentInfo struct {
	Name       string `json:"name" dc:"agent 名称"`
	Version    string `json:"version" dc:"版本信息"`
	UrlAddress string `json:"urlAddress" dc:"下载地址"`
	Status     string `json:"status" dc:"agent状态"`
	RunTimeout int    `json:"runTimeout" dc:"启动超时时间"`
}

type AgentDetails struct {
	AgentInfo
	Details string `json:"details" dc:"agent 运行详细信息"`
}
