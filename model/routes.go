package model

type Route struct {
	Model
	Name     string `gorm:"size:30;DEFAULT:''" mapstructure:"name" json:"name"`     // 路线名称
	Type     int    `gorm:"type:tinyint;DEFAULT:0" mapstructure:"type" json:"type"` // 驾车方案
	Comment  string `gorm:"type:text;" mapstructure:"comment" json:"comment"`       // 说明注释
	PathJson string `gorm:"type:text;" mapstructure:"path_json" json:"path_json"`   // 路径点信息
}
