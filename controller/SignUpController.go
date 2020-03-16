package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// 导入session包
	"github.com/gin-contrib/sessions"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/pkg/logging"
	"github.com/y-transport-server/service"
)

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
