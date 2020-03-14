package router

import (
	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/controller"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	testRouter(router)
	authRouter(router)
	return router
}
func testRouter(router *gin.Engine) {
	router.GET("/ping", controller.Ping)
}
func authRouter(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", controller.SignUp)
	}
}
