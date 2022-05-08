package controller

import (
	"context"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

type user struct{}

var User = user{}

func (self *user) CurrentInfo(ctx context.Context, req *v1.CurrentUserInfoReq) (res *v1.UserInfoRes, erro error) {

	id := service.Auth().GetIdentity(ctx)

	uid := gconv.String(id)

	user, err := service.User().Get(ctx, uid)
	if err != nil {
		return
	}

	res = &v1.UserInfoRes{
		Uid:      uid,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Phone:    user.Phone,
	}
	return

}
