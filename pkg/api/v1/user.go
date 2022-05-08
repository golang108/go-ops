package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type AuthLoginReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"用户" summary:"登录"`
	Username string `json:"username" dc:"username"`
	Passwd   string `json:"passwd"`
}

type AuthLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AuthRefreshTokenReq struct {
	g.Meta `path:"/user/refresh_token" method:"post" tags:"用户" summary:"token续期"`
}

type AuthRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AuthLogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" tags:"用户" summary:"登出"`
}

type AuthLogoutRes struct {
}

type CurrentUserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"用户信息"`
}

type UserInfoRes struct {
	Uid      string `json:"uid" dc:"uid"`
	Username string `json:"username" dc:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}
