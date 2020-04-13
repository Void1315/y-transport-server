package admin_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/service/admin_service"
)

func TripList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	data, err := bindListParam(c)
	if err != nil {
		appG.Response(e.SUCCESS, e.ERROR, err)
		return
	}
	resData := admin_service.TripList(data)
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func TripCreate(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		// form model.Driver
		form admin_service.TripCreateForm
	)
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	resData, err := admin_service.TripCreate(&form)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func TripOne(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := admin_service.TripOne(id)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, result)
}

func TripEdit(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form admin_service.TripCreateForm
	)
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	resData, err := admin_service.TripEdit(&form)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}
