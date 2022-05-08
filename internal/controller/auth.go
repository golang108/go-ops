package controller

import (
	"context"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"
)

type authController struct{}

var Auth = authController{}

func (c *authController) Login(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	res = &v1.AuthLoginRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return
}

func (c *authController) RefreshToken(ctx context.Context, req *v1.AuthRefreshTokenReq) (res *v1.AuthRefreshTokenRes, err error) {
	res = &v1.AuthRefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *authController) Logout(ctx context.Context, req *v1.AuthLogoutReq) (res *v1.AuthLogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
