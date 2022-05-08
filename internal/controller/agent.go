package controller

import (
	"context"
	v1 "go-ops/pkg/api/v1"
)

var Agent *agent = new(agent)

type agent struct{}

func (self *agent) Create(ctx context.Context, req *v1.AddAgentReq) (res *v1.AgentInfoRes, err error) {
	return
}

func (self *agent) Update(ctx context.Context, req *v1.UpdateAgentReq) (res *v1.AgentInfoRes, err error) {
	return
}

func (self *agent) Query(ctx context.Context, req *v1.QueryAgentStatusReq) (res *v1.QueryAgentStatusRes, err error) {
	return
}
