package admin_service

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/goinggo/mapstructure"
	"github.com/skip2/go-qrcode"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/service"
)

type OrderCreateForm struct {
	StartId uint `json:"start_id"`
	EndId   uint `json:"end_id"`
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

func OrderCreate(data *OrderCreateForm) (string, error) {
	price := "0.01"
	u2, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	order := &model.Order{
		UuId:    u2.String(),
		StartId: data.StartId,
		EndId:   data.EndId,
		Price:   0.01,
		Status:  0, // 状态未支付
	}
	if err0 := model.Db.Save(&order).Error; err0 != nil {
		return "", err0
	}

	returnURL := "http://www.yhy1315.cn/api/admin/zfb/return"
	payURL, err := service.WebPageAlipay(order.UuId, price, returnURL)
	return payURL, err
}

func OrderOne(id int) (*model.Order, error) {
	order := &model.Order{}

	if err := model.Db.Set("gorm:auto_preload", true).First(&order, id).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func OrderComplete(uuid string) (string, error) {
	order := &model.Order{}
	if err := model.Db.Where("uu_id = ?", uuid).First(&order).Error; err != nil {
		return "", err
	}
	if order.Status == 0 { // 未支付
		order.Status = 1 // 已支付
		model.Db.Save(&order)
	}

	err := qrcode.WriteFile("http://www.yhy1315.cn/api/check_order/"+uuid, qrcode.Medium, 256, "./static/pay/"+uuid+".png")
	if err != nil {
		return "", err
	}
	return "http://www.yhy1315.cn/api/static/pay/" + uuid + ".png", nil
}

func CheckOrder(uuid string) error {
	order := &model.Order{}
	if err := model.Db.Where("uu_id = ?", uuid).First(&order).Error; err != nil {
		return err
	}
	if order.Status == 1 { // 已支付
		order.Status = 2
		return model.Db.Save(&order).Error
	}
	if order.Status == 2 {
		return errors.New("订单已完成")
	}
	if order.Status == 0 {
		return errors.New("订单未支付")
	}
	return errors.New("订单错误")
}
