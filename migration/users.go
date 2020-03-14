package migration

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
    Username string // 列名为 `username`
    Password string // 列名为 `password`
}

func Setup(db *gorm.DB){

}