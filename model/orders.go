package model

type Order struct {
	Model
	UuId    string  `gorm:"size:36;DEFAULT:''" mapstructure:"uu_id" json:"uu_id"`        // 路线名称
	Status  int     `gorm:"type:tinyint;DEFAULT:-1" mapstructure:"status" json:"status"` // 订单状态
	UserId  uint    `gorm:"DEFAULT:0" mapstructure:"user_id" json:"user_id"`             // 驾车方案
	TripId  uint    `gorm:"DEFAULT:0" mapstructure:"trip_id" json:"trip_id"`             // 车次外键
	StartId uint    `gorm:"DEFAULT:0" mapstructure:"start_id" json:"start_id"`
	EndId   uint    `gorm:"DEFAULT:0" mapstructure:"end_id" json:"end_id"`
	Price   float32 `gorm:"type:DOUBLE(6,2);DEFAULT:0.0" json:"price"`
}
