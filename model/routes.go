package model

type Route struct {
	Model
	PathJson string `gorm:"type:text;" json:"path_json"`
}
