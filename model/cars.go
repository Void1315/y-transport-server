package model

type Car struct {
	Model
	Name     string `gorm:"size:30;DEFAULT:''" mapstructure:"name" json:"name"`             // 车辆名称
	Type     int    `gorm:"type:tinyint;DEFAULT:0" mapstructure:"type" json:"type"`         // 车辆类型
	Number   string `gorm:"size:30;DEFAULT:''" mapstructure:"number" json:"number"`         // 车牌号
	Phone    string `gorm:"size:20;DEFAULT:'';NOT NULL" mapstructure:"phone" json:"phone"`  // 车辆联系方式
	Image    string `gorm:"type:text;" mapstructure:"image" json:"image"`                   // 车辆图片
	Capacity int    `gorm:"type:tinyint;DEFAULT:0" mapstructure:"capacity" json:"capacity"` // 最大载客量
}
