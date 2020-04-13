package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/y-transport-server/pkg/setting"
)

var jwtSecret []byte

//Claims 前台token结构
type Claims struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//AdminClaims 后台token结构
type AdminClaims struct {
	User     string `json:"user"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//GenerateTokenAdmin 签发针对后台管理平台的token
func GenerateTokenAdmin(user string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(setting.JwtSetting.ExpireTime) * time.Hour)
	claims := AdminClaims{
		user,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "y-transport-admin",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

//ParseTokenAdmin 解析后台token
func ParseTokenAdmin(token string) (*AdminClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

//GenerateToken 签发token 针对用户端
func GenerateToken(phone string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(setting.JwtSetting.ExpireTime) * time.Hour)
	claims := Claims{
		phone,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "y-transport",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

//ParseToken 解析前台token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
