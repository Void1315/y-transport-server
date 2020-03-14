package migration

import (
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
