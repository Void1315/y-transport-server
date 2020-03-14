package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/y-transport-server/model"
)

func Setup(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
