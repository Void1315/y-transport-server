package admin_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/y-transport-server/pkg/app"
	"github.com/y-transport-server/pkg/e"
	"github.com/y-transport-server/service/admin_service"
)

type routeCreate struct {
	PathJson string `json:"path_json" valid:"Required"`
}

type routeEdit struct {
	PathJson interface{} `json:"path_json" valid:"Required"`
	Name     string      `json:"name" valid:"Required"`
	Type     int         `json:"type"`
	Comment  string      `json:"comment"`
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
	id, _ := strconv.Atoi(c.Param("id"))
	adminService := admin_service.Route{Id: uint(id)}
	result, err := adminService.RouteOne()
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, result)
}

//RouteCreate 创建路线
func RouteCreate(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form routeCreate
	)
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	adminService := admin_service.Route{PathJson: form.PathJson}
	resData, err := adminService.RouteCreate()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}

//RouteEdit 修改路线
func RouteEdit(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form routeEdit
	)
	id, _ := strconv.Atoi(c.Param("id"))
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, nil)
		return
	}
	b, _ := json.Marshal(form.PathJson)
	fmt.Println(string(b))
	adminService := admin_service.Route{Id: uint(id), PathJson: string(b), Name: form.Name, Type: form.Type, Comment: form.Comment}
	resData, err := adminService.RouteEdit()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, resData)
}
