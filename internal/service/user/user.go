package usercommon

import (
	"context"
	"github.com/just-zb/webook/internal/entity"
)

type UserRepo interface {
	AddUser(ctx context.Context, user *entity.User) (err error)
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
