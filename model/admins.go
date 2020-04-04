package model

import (
	"github.com/jinzhu/gorm"
)

//Admin Admins 表结构
type Admin struct {
	gorm.Model
	User     string `gorm:"size:20;NOT NULL;unique" json:"user"` // 列名为 `user` 用户名 登录账号 联合主键
	Password string `gorm:"size:100;NOT NULL;" json:"password"`  // 列名为 `password` 密码的sha256值
	Token    string `gorm:"size:512;DEFAULT:''" json:"token"`    // 列名为 `token` jwt登录凭证
}

// //CreactAdmin 创建管理员账户
// func CreactAdmin(user string, password string, email string, token string) (*Admin, error) {
// 	admin := Admin{User: user, Password: password, Email: email, Token: token}
// 	if result := Db.Save(&admin); result.Error != nil {
// 		return nil, result.Error
// 	} else {
// 		return result.Value.(*Admin), nil
// 	}
// }

// //CreatUser 创建用户
// func CreatUser(phone string, password string, hashPassword string, salt string, token string) {
// 	user := User{Phone: phone, Password: hashPassword, Token: token, Salt: salt}
// 	if Db.NewRecord(user) { // => 主键为空返回`true`
// 		Db.Create(&user)
// 	}
// }

// //GetUserWithPhone 通过phone查询用户
// func GetUserWithPhone(phone string) (*User, error) {
// 	var user User
// 	err := Db.Where("phone=?", phone).First(&user).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, err
// 	}
// 	return &user, err
// }
