package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"net/http"
)

// type createAdminUser struct {
// 	User     string `json:"user" valid:"Required; MaxSize(20);"`                // 用户名 也是登录账号
// 	Password string `json:"password" valid:"Required; MaxSize(20);MinSize(6);"` // 密码被加密过的
// 	Token    string `json:"token" valid:"Required; MaxSize(20);MinSize(6);"`    // 列名为 `token` jwt登录凭证
// }

//AdminCheck 后台的身份验证
func AdminCheck(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// func CreateAdminUser(c *gin.Context) {
// 	var (
// 		appG = app.Gin{C: c}
// 		form createAdminUser
// 	)
// 	httpCode, errCode := app.BindAndValid(c, &form)
// 	if errCode != e.SUCCESS {
// 		appG.Response(httpCode, errCode, nil)
// 		return
// 	}
// 	// adminService.PhoneSignUp(c, form.Phone)
// 	result, err := model.CreactAdmin(form.User, form.Password, form.Email, form.Salt, form.Token)
// 	if err != nil {
// 		appG.Response(httpCode, e.ERROR, err)
// 		return
// 	}
// 	appG.Response(httpCode, e.SUCCESS, result)
// }
