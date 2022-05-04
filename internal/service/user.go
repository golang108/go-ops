package service

import (
	"context"

	v1 "go-ops/api/v1"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"
	"go-ops/internal/service/internal/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	// sUser is service struct of module User.
	sUser struct{}
)

var (
	// insUser is the instance of service User.
	insUser = sUser{}
)

// User returns the interface of User service.
func User() *sUser {
	return &insUser
}

func (s *sUser) Get(ctx context.Context, uid string) (r *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where("uid = ?", uid).Scan(&r)
	if err != nil {
		return
	}
	if r == nil {
		return
	}
	return
}

func (s *sUser) IsUsernameAvailable(ctx context.Context, username string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Username: username,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// SignIn creates session for given user account.
func (s *sUser) SignIn(ctx context.Context, in *v1.AuthLoginReq) (r *entity.User, err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Username: in.Username,
		Passwd:   in.Passwd,
	}).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New(`username or Password not correct`)
	}
	return user, nil
}
