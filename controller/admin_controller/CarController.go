package admin_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/service/admin_service"
)

func CarList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	data, err := bindListParam(c)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	resData := admin_service.CarList(data)
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func CarCreate(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		// form model.Driver
		form admin_service.CarCreateForm
	)
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	resData, err := admin_service.CarCreate(&form)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func CarOne(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := admin_service.CarOne(id)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, result)
}

func CarEdit(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form admin_service.CarCreateForm
	)
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	resData, err := admin_service.CarEdit(&form)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}
