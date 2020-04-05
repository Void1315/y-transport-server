package model

import (
	"github.com/jinzhu/gorm"
)

type Route struct {
	gorm.Model
	PathJson string `gorm:"type:text;DEFAULT:''" json:"path_json"`
}
