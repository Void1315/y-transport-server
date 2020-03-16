package service

import (
	// 导入session包
	"github.com/gin-contrib/sessions"
	// 导入gin框架包
	"github.com/gin-gonic/gin"
	"github.com/y-transport-server/pkg/util"
)

//PhoneSignUp 手机号注册 用以获取验证码
func PhoneSignUp(c *gin.Context, phone string) {
	session := sessions.Default(c)
	session.Set("sms_code", "1315")
	session.Set("phone", phone)
	session.Save()
}

//CodeSignUp 验证码注册
func CodeSignUp(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	smsCode := session.Get("sms_code")
	if code == smsCode {
		return true
	}
	return false
}

//CreateSignUp 完善用户信息
func CreateSignUp(c *gin.Context, phone string, password string) (string, string, string, error) {
	salt := util.GetSalt(10)
	hashPassword := util.Make(password, salt)
	token, err := util.GenerateToken(phone, hashPassword)
	return salt, hashPassword, token, err

}
