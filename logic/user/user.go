package user

import (
	"context"
	"errors"
	"fmt"
	"peony/api/v1"
	"peony/model"
	"peony/service"
	"peony/utils"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (u *sUser) Register(ctx context.Context, in *v1.RegisterReq) error {
	// 查询邮箱是否已存在
	m := u.search(in)
	if m.Email != "" || m.Email == in.Email {
		return errors.New("邮箱已存在")
	}

	// 加密密码
	hPassword, _ := utils.NewHash([]byte(in.Password))

	// 注册
	user := model.User{
		UUID:     utils.UUID.String(),
		Email:    in.Email,
		Password: string(hPassword),
	}
	return utils.DB.Create(&user).Error
}

// Login 用户登录
func (u *sUser) Login(ctx context.Context, in *v1.LoginReq, emailCh chan string) error {
	m := u.search(in)
	if m.Email != in.Email {
		return errors.New("用户不存在")
	} else if err := utils.NewDeHash([]byte(m.Password), []byte(in.Password)); err != nil {
		return errors.New("密码错误")
	} else {
		emailCh <- m.Email
		return nil
	}
}

func (u *sUser) search(in any) *model.User {
	var user model.User
	var email string

	switch m := in.(type) {
	case *v1.LoginReq:
		email = m.Email
	case *v1.RegisterReq:
		email = m.Email
	default:
		return nil
	}

	utils.DB.Model(model.User{}).Where("email=?", email).Scan(&user)
	return &user
}

func (u *sUser) UserExist(email string) bool {
	result := utils.DB.Model(&model.User{}).Where("email=?", &email).Find(&model.User{}).RowsAffected
	fmt.Println(result)
	return result > 0
}
