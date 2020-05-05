package admin_controller

import (
	"fmt"
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
	uuid := c.Query("out_trade_no")
	data, err := admin_service.OrderComplete(uuid)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	s := fmt.Sprintf(`<h1>请保存车票二维码,二维码用于乘车验票!</h1><img src="%s"></img>`, data)
	c.String(200, s)
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

func CheckOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	uuid := c.Param("uuid")
	err := admin_service.CheckOrder(uuid)
	if err != nil {
		appG.Response(e.ERROR, e.ERROR, nil)
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	s := fmt.Sprintf(`<h1>订单验证完成!请上车</h1>`)
	c.String(200, s)
}
