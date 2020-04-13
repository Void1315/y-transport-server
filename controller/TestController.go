package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
)

// Ping 测试连接，放回pong
func Ping(c *gin.Context) {
	var appG = app.Gin{C: c}
	appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, "pong")
}
