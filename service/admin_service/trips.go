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
