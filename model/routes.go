package model

type Route struct {
	Model
	PathJson string `gorm:"type:text;" mapstructure:"path_json" json:"path_json"`
}
