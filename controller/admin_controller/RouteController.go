package admin_controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"net/http"
	// "github.com/y-transport-server/service/admin_service"
)

func RouteList(c *gin.Context) {
	var (
		appG  = app.Gin{C: c}
		valid = validation.Validation{}
	)
	sort := ""
	if arg := c.Query("sort"); arg != "" {
		sort = com.StrTo(arg).String()
		valid.Required(sort, "sort")
	}
	// range := []int{}
	// filter := []string{}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func RouteOne(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
