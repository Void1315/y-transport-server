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
	authRouter(router)
	adminRoute(router)
	return router
}
func testRouter(router *gin.Engine) {
	router.GET("/ping", controller.Ping)
}
func adminRoute(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		// admin.POST("/create", controller.CreateAdminUser)
		admin.GET("/check", controller.AdminCheck)
	}

}

func authRouter(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/check", jwt.JWT(), controller.Check)
		auth.POST("/sign_in", controller.SignIn)
		signUp := auth.Group("/sign_up")
		{
			signUp.POST("/phone", controller.SignUpPhone)
			// signUp.POST("/code", controller.SignUpCode)
			signUp.POST("/create", controller.SignUpCreate)
		}
	}
}
