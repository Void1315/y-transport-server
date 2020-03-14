package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
)

type signUp struct {
	Email    string `valid:"Required; MaxSize(30);Email"`
	Password string `valid:"Required; MaxSize(30);"`
}

//SignUp 注册
func SignUp(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form signUp
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	c.String(http.StatusOK, "pong")
}
