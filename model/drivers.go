package model

type Driver struct {
	Model
	Name       string `gorm:"size:20;DEFAULT:'';NOT NULL" mapstructure:"name" json:"name"`          // 司机名称
	Phone      string `gorm:"size:20;DEFAULT:'';NOT NULL" mapstructure:"phone" json:"phone"`        // 司机联系方式
	Image      string `gorm:"size:255;DEFAULT:'';NOT NULL" mapstructure:"image" json:"image"`       // 司机照片
	DrivingAge int    `gorm:"type:tinyint;DEFAULT:0" mapstructure:"driving_age" json:"driving_age"` // 司机驾龄
	Age        int    `gorm:"type:tinyint;DEFAULT:0" mapstructure:"age" json:"age"`                 // 司机年龄
	CarId      uint   `gorm:"DEFAULT:0" mapstructure:"car_id" json:"car_id"`
}
