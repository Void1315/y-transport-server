package admin_service

import (

	// "github.com/y-transport-server/controller/admin_controller"

	"strings"

	"github.com/goinggo/mapstructure"
	"github.com/y-transport-server/model"
)

type Route struct {
	Id       uint   `json:"id"`
	PathJson string `json:"path_json"`
	Name     string `json:"name"`
	Type     int    `json:"type"`    // 驾车方案
	Comment  string `json:"comment"` // 说明注释
}

type ListParam struct {
	Sort   []string          `json:"sort"`
	Limit  int               `json:"limit"`
	Page   int               `json:"page"`
	Filter map[string]string `json:"filter"`
}

func RouteAll() (*[]model.Route, error) {
	routers := &[]model.Route{}
	if err := model.Db.Find(&routers).Error; err != nil {
		return nil, err
	}
	return routers, nil
}

//RouteList 分页
func RouteList(data *ListParam) model.PageJson {
	routes := make([]model.Route, 0)
	var routeModel model.Route
	// model.Db.Find().Count()
	if err := mapstructure.Decode(data.Filter, &routeModel); err != nil {
		return model.PageJson{}
	}
	Db := model.Db.Where(&routeModel).Limit(data.Limit).Offset((data.Page - 1) * data.Limit).Order(strings.Join(data.Sort[:], " "))
	Db.Find(&routes)

	var total = 0
	model.Db.Model(&model.Route{}).Where(&routeModel).Count(&total)

	page := model.PageJson{
		Data:  routes,
		Page:  data.Page,
		Total: total,
		Size:  len(routes),
	}

	return page
}

//RouteOne 单一
func (r *Route) RouteOne() (*model.Route, error) {
	route := &model.Route{}

	if err := model.Db.First(&route, r.Id).Error; err != nil {
		return nil, err
	}
	return route, nil
}

//RouteCreate 创建
func (r *Route) RouteCreate() (*model.Route, error) {
	route := &model.Route{
		PathJson: r.PathJson,
		Name:     r.Name,
		Comment:  r.Comment,
		Type:     r.Type,
	}
	if err := model.Db.Save(route).Error; err != nil {
		return nil, err
	}
	return route, nil
}

func (r *Route) RouteEdit() (*model.Route, error) {
	route := &model.Route{
		Model:    model.Model{ID: r.Id},
		PathJson: r.PathJson,
		Comment:  r.Comment,
		Name:     r.Name,
		Type:     r.Type,
	}
	if err := model.Db.Save(route).Error; err != nil {
		return nil, err
	}
	return route, nil
}

func RouteDelete(id uint) error {
	route := &model.Route{
		Model: model.Model{ID: id},
	}
	if err := model.Db.First(route).Error; err != nil {
		return err
	}
	if err0 := model.Db.Delete(&route).Error; err0 != nil {
		return err0
	}
	// if err1 := model.Db.Table("cars").Where("route_id = ?", id).Updates(map[string]interface{}{"route_id": 0}).Error; err1 != nil {
	// 	return err1
	// }
	return nil
}

func RouteDeleteTrue(id uint) error {
	route := &model.Route{
		Model: model.Model{ID: id},
	}
	if err := model.Db.Unscoped().First(route).Error; err != nil {
		return err
	}
	if err0 := model.Db.Unscoped().Delete(&route).Error; err0 != nil {
		return err0
	}
	if err1 := model.Db.Table("cars").Where("route_id = ?", id).Updates(map[string]interface{}{"route_id": 0}).Error; err1 != nil {
		return err1
	}
	return nil
}
