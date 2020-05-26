package admin_service

import (
	"strings"

	"github.com/goinggo/mapstructure"
	"github.com/y-transport-server/model"
)

type TripCreateForm struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      int    `json:"type" mapstructure:"type" ` // 订单状态
	CarId     uint   `json:"car_id" mapstructure:"car_id" `
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type TripWithRoute struct {
	RouteId uint `json:"route_id"`
}

func TripList(data *ListParam) model.PageJson {
	trips := make([]model.Trip, 0)
	var tripModel model.Trip
	if err := mapstructure.Decode(data.Filter, &tripModel); err != nil {
		return model.PageJson{}
	}
	Db := model.Db.Where(&tripModel).Limit(data.Limit).Offset((data.Page - 1) * data.Limit).Order(strings.Join(data.Sort[:], " "))
	Db.Set("gorm:auto_preload", true).Find(&trips)
	var total = 0
	model.Db.Model(&model.Trip{}).Where(&tripModel).Count(&total)

	page := model.PageJson{
		Data:  trips,
		Page:  data.Page,
		Total: total,
		Size:  len(trips),
	}
	return page
}

func TripCreate(data *TripCreateForm) (*model.Trip, error) {
	trip := &model.Trip{
		Name:      data.Name,
		CarId:     data.CarId,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		Type:      data.Type,
	}
	if err2 := model.Db.Save(trip).Error; err2 != nil {
		return nil, err2
	}
	return trip, nil
}

func TripOne(id int) (*model.Trip, error) {
	trip := &model.Trip{}

	if err := model.Db.Set("gorm:auto_preload", true).First(&trip, id).Error; err != nil {
		return nil, err
	}
	return trip, nil
}

func TripEdit(data *TripCreateForm) (*model.Trip, error) {
	trip := &model.Trip{
		Model:     model.Model{ID: uint(data.ID)},
		Name:      data.Name,
		CarId:     data.CarId,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		Type:      data.Type,
	}
	if err := model.Db.Save(trip).Error; err != nil {
		return nil, err
	}
	return trip, nil
}

func TripDelete(id int) error {
	trip := &model.Trip{
		Model: model.Model{ID: uint(id)},
	}
	if err := model.Db.Find(&trip).Error; err != nil {
		return err
	}
	if err0 := model.Db.Delete(&trip).Error; err0 != nil {
		return err0
	}
	return nil
}

func TripDeleteTrue(id int) error {
	trip := &model.Trip{
		Model: model.Model{ID: uint(id)},
	}
	if err := model.Db.Unscoped().Find(&trip).Error; err != nil {
		return err
	}
	// 物理删除
	if err0 := model.Db.Unscoped().Delete(&trip).Error; err0 != nil {
		return err0
	}
	// 关联删除
	if err1 := model.Db.Table("orders").Where("trip_id = ?", id).Updates(map[string]interface{}{"trip_id": 0}).Error; err1 != nil {
		return err1
	}
	return nil
}

func GetTripWithRoute(routeId uint) ([]model.Trip, error) {
	trips := make([]model.Trip, 0)
	if err := model.Db.Table("trips").Where("car_id IN (?)", model.Db.Table("cars").Select("id").Where("route_id = ?", routeId).SubQuery()).Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}
