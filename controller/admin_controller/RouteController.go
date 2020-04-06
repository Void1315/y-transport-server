package admin_controller

import (
	"net/http"

	// "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/service/admin_service"
)

type Bid struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
type A struct {
	Data  []Bid `json:"data"`
	Total int   `json:"total"`
}

func RouteList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	data, err := bindListParam(c)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	resData := admin_service.RouteList(data)
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func RouteOne(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
