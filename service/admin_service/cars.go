package admin_service

import (
	"strings"

	"github.com/goinggo/mapstructure"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/util"
)

type CarCreateForm struct {
	ID       int                `json:"id"`
	Name     string             `json:"name" valid:"Required"`
	Phone    string             `json:"phone" valid:"Required;Mobile"`
	Type     int                `json:"type"  valid:"Required;"`
	Image    []util.UploadImage `json:"image" valid:"Required"`
	Capacity int                `json:"capacity"  valid:"Required;"`
	Number   string             `json:"number"  valid:"Required;"`
}

func CarList(data *ListParam) model.PageJson {
	cars := make([]model.Car, 0)
	var carModel model.Car
	if err := mapstructure.Decode(data.Filter, &carModel); err != nil {
		return model.PageJson{}
	}
	Db := model.Db.Where(&carModel).Limit(data.Limit).Offset((data.Page - 1) * data.Limit).Order(strings.Join(data.Sort[:], " "))
	Db.Find(&cars)

	var total = 0
	model.Db.Model(&model.Car{}).Where(&carModel).Count(&total)

	page := model.PageJson{
		Data:  cars,
		Page:  data.Page,
		Total: total,
		Size:  len(cars),
	}
	return page
}

func CarCreate(data *CarCreateForm) (*model.Car, error) {
	savedImages, err := util.SaveImage(&data.Image)
	if err != nil {
		return nil, err
	}
	car := &model.Car{
		Name:     data.Name,
		Phone:    data.Phone,
		Type:     data.Type,
		Capacity: data.Capacity,
		Number:   data.Number,
		Image: &model.SavedImageMap{
			SavedImage: *savedImages,
			Valid:      true,
		},
	}
	if err2 := model.Db.Save(car).Error; err2 != nil {
		return nil, err2
	}
	return car, nil
}

func CarEdit(data *CarCreateForm) (*model.Car, error) {
	oldCar := &model.Car{}
	oldImages := make([]model.SavedImage, 0)
	model.Db.Find(&oldCar, data.ID)
	savedImages, err1 := util.EditImage(&data.Image, &oldImages)
	if err1 != nil {
		return nil, err1
	}
	car := &model.Car{
		Model:    model.Model{ID: uint(data.ID)},
		Name:     data.Name,
		Phone:    data.Phone,
		Type:     data.Type,
		Capacity: data.Capacity,
		Number:   data.Number,
		Image: &model.SavedImageMap{
			SavedImage: *savedImages,
			Valid:      true,
		},
	}
	if err2 := model.Db.Save(car).Error; err2 != nil {
		return nil, err2
	}
	return car, nil
}

func CarOne(id int) (*model.Car, error) {
	car := &model.Car{}

	if err := model.Db.First(&car, id).Error; err != nil {
		return nil, err
	}
	return car, nil
}
