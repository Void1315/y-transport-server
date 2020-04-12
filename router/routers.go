package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/controller"
	"github.com/y-transport-server/controller/admin_controller"
	"github.com/y-transport-server/middleware/cors"
	"github.com/y-transport-server/middleware/jwt"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Cors())
	store := cookie.NewStore([]byte("wqld1315"))
	router.Static("/static", "./static")
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
		route := admin.Group("/routes_data")
		{
			route.GET("", admin_controller.RouteList)
			route.POST("", admin_controller.RouteCreate)
			route.POST("/all", admin_controller.RouteAll)
			route.GET("/:id", admin_controller.RouteOne)
			route.POST("/edit/:id", admin_controller.RouteEdit)

		}
		driver := admin.Group("/driver")
		{
			driver.GET("", admin_controller.DriverList)
			driver.POST("", admin_controller.DriverCreate)
			driver.GET("/:id", admin_controller.DriverOne)
			driver.POST("/edit/:id", admin_controller.DriverEdit)
			driver.DELETE("/:id", admin_controller.DriverDelete)
		}
		car := admin.Group("/car")
		{
			car.GET("", admin_controller.CarList)
			car.POST("", admin_controller.CarCreate)
			car.GET("/:id", admin_controller.CarOne)
			car.POST("/edit/:id", admin_controller.CarEdit)
		}
		admin.POST("/login", admin_controller.Login)
		admin.GET("/logout", admin_controller.Logout)
		admin.GET("/check", jwt.JWT(), admin_controller.AdminCheck)
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
