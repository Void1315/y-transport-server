package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 测试连接，放回pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
