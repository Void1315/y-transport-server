package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/pkg/logging"
)

type checkAuth struct {
	Token string `json:"token";valid:"Required;"`
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
