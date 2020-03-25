package auth_service

import (
	"errors"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/util"
)

type Auth struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (a *Auth) SignIn() (*model.User, error) {
	user, err := model.GetUserWithPhone(a.Phone)
	if err != nil {
		return nil, err
	}
	hashPassword := util.Make(a.Password, user.Salt)
	if hashPassword != user.Password {
		return nil, errors.New("账号或密码错误")
	}
	return user, nil
}
