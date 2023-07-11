package service

import (
	"context"
	"peony/api/v1"
)

type (
	IUser interface {
		Login(ctx context.Context, in *v1.LoginReq) error
		Register(ctx context.Context, in *v1.RegisterReq) error
		UserExist(email string) bool
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
