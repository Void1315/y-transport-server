package model

type Trip struct {
	Model
	Name      string `gorm:"size:20;DEFAULT:''" mapstructure:"name" json:"name"`     // 路线名称
	Type      int    `gorm:"type:tinyint;DEFAULT:0" mapstructure:"type" json:"type"` // 订单状态
	CarId     uint   `gorm:"DEFAULT:0" mapstructure:"car_id" json:"car_id"`
	Car       Car    `gorm:"foreignkey:CarId" json:"car"`
	StartTime string `gorm:"type:time;DEFAULT:'01:00'" mapstructure:"start_time" json:"start_time"`
	EndTime   string `gorm:"type:time;DEFAULT:'01:00'" mapstructure:"end_time" json:"end_time"`
}
