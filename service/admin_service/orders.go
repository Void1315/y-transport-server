package admin_service

import (
	"strings"

	"github.com/goinggo/mapstructure"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/service"
)

type OrderCreateForm struct {
}

func OrderList(data *ListParam) model.PageJson {
	orders := make([]model.Order, 0)
	var orderModel model.Order
	// model.Db.Find().Count()
	if err := mapstructure.Decode(data.Filter, &orderModel); err != nil {
		return model.PageJson{}
	}
	Db := model.Db.Where(&orderModel).Limit(data.Limit).Offset((data.Page - 1) * data.Limit).Order(strings.Join(data.Sort[:], " "))
	Db.Find(&orders)

	var total = 0
	model.Db.Model(&model.Order{}).Where(&orderModel).Count(&total)

	page := model.PageJson{
		Data:  orders,
		Page:  data.Page,
		Total: total,
		Size:  len(orders),
	}
	return page
}

func OrderCreate(data *OrderCreateForm) (*model.Order, error) {
	order := &model.Order{}
	service.WebPageAlipay()
	return order, nil
}

func OrderOne(id int) (*model.Order, error) {
	order := &model.Order{}

	if err := model.Db.Set("gorm:auto_preload", true).First(&order, id).Error; err != nil {
		return nil, err
	}
	return order, nil
}
