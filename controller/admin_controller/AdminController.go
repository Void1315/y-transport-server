package admin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/service/admin_service"
)

type login struct {
	User     string `json:"user" valid:"Required; MaxSize(20);"`                // 用户名 也是登录账号
	Password string `json:"password" valid:"Required; MaxSize(20);MinSize(6);"` // 密码被加密过的
}

//AdminCheck 后台的身份验证
func AdminCheck(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//Login 后台登录
func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form login
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	adminService := admin_service.Admin{User: form.User, Password: form.Password}
	token, err := adminService.Login()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_ADMIN_USER, err)
	} else {
		c.SetCookie("token", token, 86400, "/", "", false, false)
		appG.Response(http.StatusOK, e.SUCCESS, nil)
	}
}

//Logout 退出登录
func Logout(c *gin.Context) {
	appG := app.Gin{C: c}
	c.SetCookie("token", "token", -1, "/", "", false, false)
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
