package admin_service

import (
	"github.com/y-transport-server/model"
	// "github.com/y-transport-server/pkg/util"
)

type Route struct {
	PathJson string `json:"path_json"`
}

//RouteList 分页
func RouteList() model.PageJson {
	routes := make([]model.Route, 0)
	Db := model.Db.Limit(10).Offset(0)
	Db.Find(&routes)
	page := model.PageJson{
		Data:  routes,
		Total: 2,
		Page:  1,
		Size:  2,
	}
	return page
}
