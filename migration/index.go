package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/y-transport-server/model"
)

func Setup(db *gorm.DB) {
	var admin model.Admin
	db.AutoMigrate(&model.User{}, &model.Admin{}, &model.Route{}, &model.Driver{}, &model.Car{})
	db.Where("user=?", "yhy1315").Find(&admin)
	if admin.ID == 0 {
		var adminUser = model.Admin{User: "yhy1315", Password: "9283f1b7821af07bc5d3a7dba9525fdaa9fcd68aa5e579192fd527d0e69db74a"}
		db.Create(&adminUser)
	}
}
