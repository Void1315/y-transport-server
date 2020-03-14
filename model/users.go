package model

import "github.com/jinzhu/gorm"

//User Users 表结构
type User struct {
	gorm.Model
	Username   string `gorm:"size:50;DEFAULT:''"`  // 列名为 `username` 用户名称
	UsernameZh string `gorm:"size:10;DEFAULT:''"`  // 列名为 `username_zh` 用户中文名称
	Email      string `gorm:"size:30;DEFAULT:''"`  // 列名为 `email` 账户邮箱，
	Avatar     string `gorm:"size:255;DEFAULT:''"` // 列名为 `avatar` 头像绝对路径,url
	Password   string `gorm:"size:100;DEFAULT:''"` // 列名为 `password` 密码的sha256值
	Salt       string `gorm:"size:20;DEFAULT:''"`  // 列名为 `salt` 密码盐
	Token      string `gorm:"size:512;DEFAULT:''"` // 列名为 `salt` 密码盐
	Type       bool   `gorm:"NOT NULL;DEFAULT:0"`  // 列名为 `salt` 密码盐
}
