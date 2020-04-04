package cors

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// method := c.Request.Method
		// origin := c.Request.Header.Get("Origin")
		// // origin := "http://127.0.0.1:3000"
		// if origin != "" {
		// 	c.Header("Access-Control-Allow-Origin", origin)
		// 	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// 	c.Header("Access-Control-Allow-Headers", "*")
		// 	c.Header("Access-Control-Allow-Credentials", "true")
		// }
		// if method == "OPTIONS" {
		// 	c.AbortWithStatus(http.StatusNoContent)
		// }
		c.Next()
	}
}
