package router

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/y-transport-server/controller"
	"github.com/y-transport-server/middleware/jwt"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("wqld1315"))
	router.Use(sessions.Sessions("ysession", store))
	testRouter(router)
	signUpRouter(router)
	authRouter(router)
	return router
}
func testRouter(router *gin.Engine) {
	router.GET("/ping", controller.Ping)
}
func signUpRouter(router *gin.Engine) {
	signUp := router.Group("/sign_up")
	{
		signUp.POST("/phone", controller.SignUpPhone)
		signUp.POST("/code", controller.SignUpCode)
		signUp.POST("/create", controller.SignUpCreate)
	}
}
func authRouter(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.Use(jwt.JWT())
	{
		auth.POST("/check", controller.Check)
	}
}
