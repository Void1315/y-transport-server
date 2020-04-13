package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/y-transport-server/pkg/setting"
)

var Db *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id" mapstructure:"id"`
	CreatedAt time.Time  `json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" mapstructure:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" mapstructure:"deleted_at"`
}

// PageJson 分页结构体
type PageJson struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
	Size  int         `json:"size"`
	Page  int         ` json:"page"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	Db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	Db.LogMode(true)

	gorm.DefaultTableNameHandler = func(Db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
}
