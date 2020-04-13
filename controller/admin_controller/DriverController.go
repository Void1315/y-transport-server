package admin_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/service/admin_service"
)

type driverCreate struct {
	Name       string            `json:"name" valid:"Required;"`
	Phone      string            `json:"phone" valid:"Required;"`
	Age        string            `json:"age" valid:"Required;"`
	DrivingAge string            `json:"driving_age" valid:"Required;"`
	Image      map[string]string `json:"image" valid:"Required;"`
}

//DriverList list
func DriverList(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	data, err := bindListParam(c)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	resData := admin_service.DriverList(data)
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func DriverOne(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := admin_service.DriverOne(id)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, result)
}

func DriverCreate(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		// form model.Driver
		form admin_service.DriverCreateForm
	)
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	resData, err := admin_service.DriverCreate(&form)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}
func DriverEdit(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form admin_service.DriverEditForm
	)
	id, _ := strconv.Atoi(c.Param("id"))
	_, errCode := app.BindAndValid(c, &form)
	form.ID = id
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	resData, err := admin_service.DriverEdit(&form)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

func DriverDelete(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	err := admin_service.DriverDelete(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, id)
}
