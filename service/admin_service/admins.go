package admin_service

import (
	"errors"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/util"
)

type Admin struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (a *Admin) Login() (string, error) {
	admin, err := model.GetAdminWithUser(a.User)
	if err != nil {
		return "", err
	}
	hashPassword := util.Make(a.Password, "")
	if hashPassword != admin.Password {
		return "", errors.New("账号或密码错误")
	}
	token, _ := util.GenerateTokenAdmin(a.User, a.Password)
	return token, nil
}
