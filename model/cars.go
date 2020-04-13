package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Car struct {
	Model
	Name     string         `gorm:"size:30;DEFAULT:''" mapstructure:"name" json:"name"`             // 车辆名称
	Type     int            `gorm:"type:tinyint;DEFAULT:0" mapstructure:"type" json:"type"`         // 车辆类型
	Number   string         `gorm:"size:30;DEFAULT:''" mapstructure:"number" json:"number"`         // 车牌号
	Phone    string         `gorm:"size:20;DEFAULT:'';NOT NULL" mapstructure:"phone" json:"phone"`  // 车辆联系方式
	Image    *SavedImageMap `gorm:"type:json;" mapstructure:"image" json:"image"`                   // 车辆图片
	Capacity int            `gorm:"type:tinyint;DEFAULT:0" mapstructure:"capacity" json:"capacity"` // 最大载客量
	RouteId  uint           `json:"route_id" mapstructure:"route_id"`                               // belongs to Route 外键关联
	Route    Route          `json:"route" gorm:"foreignkey:RouteId;PRELOAD:true"`                   // 关联模型
	Cars     []Car          `json:"cars" gorm:"foreignkey:CarId"`
}
type SavedImage struct {
	Path     string `json:"path"`
	FileName string `json:"file_name"`
}
type SavedImageMap struct {
	SavedImage []SavedImage `json:"image"`
	Valid      bool         `json:"valid"`
}

func NewEmptySavedImageMap() *SavedImageMap {
	return &SavedImageMap{
		SavedImage: make([]SavedImage, 0),
		Valid:      true,
	}
}

func NewSavedImageMap(src []SavedImage) *SavedImageMap {
	return &SavedImageMap{
		SavedImage: src,
		Valid:      true,
	}
}

func (ls *SavedImageMap) Scan(value interface{}) error {
	if value == nil {
		ls.SavedImage, ls.Valid = make([]SavedImage, 0), false
		return nil
	}
	t := make([]SavedImage, 0)
	if e := json.Unmarshal(value.([]byte), &t); e != nil {
		return e
	}
	ls.Valid = true
	ls.SavedImage = t
	return nil
}

func (ls *SavedImageMap) Value() (driver.Value, error) {
	if ls == nil {
		return nil, nil
	}
	if !ls.Valid {
		return nil, nil
	}

	b, e := json.Marshal(ls.SavedImage)
	return b, e
}
