package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// 导入session包
	"github.com/gin-contrib/sessions"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/pkg/logging"
	"github.com/y-transport-server/service"
	"github.com/y-transport-server/service/auth_service"
)

type checkAuth struct {
	Token string `json:"token";valid:"Required;"`
}

type signIn struct {
	Phone    string `json:"phone";valid:"Required;"`
	Password string `json:"password";valid:"Required;"`
}
type signUpPhone struct {
	Phone string `json:"phone";valid:"Required; MaxSize(20);"`
}

type signUpCode struct {
	Code string `json:"code";valid:"Required; MaxSize(5);"`
}

type signUpCreate struct {
	Password string `json:"password";valid:"Required; MaxSize(20);MinSize(6);"`
}

//SignUpPhone 手机号注册
func SignUpPhone(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form signUpPhone
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	service.PhoneSignUp(c, form.Phone)
	appG.Response(httpCode, e.SUCCESS, nil)
}

//SignUpCode 验证码注册
func SignUpCode(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form signUpCode
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	checkCode := service.CodeSignUp(c, form.Code)
	if checkCode {

		appG.Response(httpCode, e.SUCCESS, nil)
	} else {
		appG.Response(httpCode, e.ERROR_AUTH_SMS_CODE, nil)
	}

}

//SignUpCreate 完善用户信息
func SignUpCreate(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form signUpCreate
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	session := sessions.Default(c)
	phone := session.Get("phone")
	if phoneStr, ok := phone.(string); ok {
		salt, hashPassword, token, err := service.CreateSignUp(c, phoneStr, form.Password)
		if err != nil {
			logging.Warn(err)
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		} else {
			model.CreatUser(phoneStr, form.Password, hashPassword, salt, token)
			appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
				"token": token,
			})
		}
	} else {
		appG.Response(httpCode, e.ERROR, nil)
	}
}

//CheckAuth 检测身份
func Check(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form checkAuth
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	logging.Info("登录")
	appG.Response(httpCode, e.SUCCESS, nil)
}

//SignIn 正常通过账号密码登录
func SignIn(c *gin.Context) {
	logging.Info("账号密码登录")
	var (
		appG = app.Gin{C: c}
		form signIn
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	authService := auth_service.Auth{Phone: form.Phone, Password: form.Password}

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	user, err := authService.SignIn()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, user)
	}
}

//SignInWithToken 通过token登录
func SignInWithToken(c *gin.Context) {

}
