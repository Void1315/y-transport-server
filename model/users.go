package model

import (
	"github.com/jinzhu/gorm"
)

//User Users 表结构
type User struct {
	Model
	Username   string `gorm:"size:50;DEFAULT:''" json:"username"`       // 列名为 `username` 用户名称
	UsernameZh string `gorm:"size:10;DEFAULT:''" json:"username_zh"`    // 列名为 `username_zh` 用户中文名称
	Email      string `gorm:"size:30;DEFAULT:''" json:"email"`          // 列名为 `email` 账户邮箱，
	Avatar     string `gorm:"size:255;DEFAULT:''" json:"avatar"`        // 列名为 `avatar` 头像绝对路径,url
	IdNumber   string `gorm:"size:30;DEFAULT:''" json:"id_number"`      // 列名为 `id_number` 身份证号码
	Phone      string `gorm:"size:20;DEFAULT:'';NOT NULL" json:"phone"` // 列名为 `phone` 手机号码
	Password   string `gorm:"size:100;DEFAULT:''" json:"password"`      // 列名为 `password` 密码的sha256值
	Salt       string `gorm:"size:20;DEFAULT:''" json:"salt"`           // 列名为 `salt` 密码盐
	Token      string `gorm:"size:512;DEFAULT:''" json:"token"`         // 列名为 `token` jwt登录凭证
	Type       bool   `gorm:"NOT NULL;DEFAULT:0" json:"type"`           // 列名为 `type` 用户类型 0本站 1第三方
}

//CreatUser 创建用户
func CreatUser(phone string, password string, hashPassword string, salt string, token string) {
	user := User{Phone: phone, Password: hashPassword, Token: token, Salt: salt}
	if Db.NewRecord(user) { // => 主键为空返回`true`
		Db.Create(&user)
	}
}

//GetUserWithPhone 通过phone查询用户
func GetUserWithPhone(phone string) (*User, error) {
	var user User
	err := Db.Where("phone=?", phone).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, err
}
