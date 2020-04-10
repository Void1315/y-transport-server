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
	Name       string            `json:"name" `
	Phone      string            `json:"phone" `
	Age        string            `json:"age" `
	DrivingAge string            `json:"driving_age" `
	Image      map[string]string `json:"image" `
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
	}
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

func DriverEdit(c *gin.Context) {

}
