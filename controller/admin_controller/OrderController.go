package admin_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/service/admin_service"
)

//OrderList list
func OrderList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	data, err := bindListParam(c)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	resData := admin_service.OrderList(data)
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func OrderCreate(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		// form model.Driver
		form admin_service.OrderCreateForm
	)
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	resData, err := admin_service.OrderCreate(&form)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func OrderReturn(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, "完事")
}

func OrderOne(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := admin_service.OrderOne(id)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, result)
}
