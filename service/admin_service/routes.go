package admin_service

import (

	// "github.com/y-transport-server/controller/admin_controller"
	"strings"

	"github.com/goinggo/mapstructure"
	"github.com/y-transport-server/model"
)

type Route struct {
	PathJson string `json:"path_json"`
}

type ListParam struct {
	Sort   []string          `json:"sort"`
	Limit  int               `json:"limit"`
	Page   int               `json:"page"`
	Filter map[string]string `json:"filter"`
}

//RouteList 分页
func RouteList(data *ListParam) model.PageJson {
	routes := make([]model.Route, 0)
	var routeModel model.Route
	if err := mapstructure.Decode(data.Filter, &routeModel); err != nil {
		return model.PageJson{}
	}
	Db := model.Db.Where(&routeModel).Limit(data.Limit).Offset((data.Page - 1) * data.Limit)
	var total = 0
	Db.Order(strings.Join(data.Sort[:], " ")).Find(&routes).Count(&total)
	page := model.PageJson{
		Data:  routes,
		Page:  data.Page,
		Total: total,
		Size:  len(routes),
	}
	return page
}

//RouteOne 单一
func RouteOne() {

}
