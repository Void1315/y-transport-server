package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/y-transport-server/pkg/setting"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(phone string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(setting.JwtSetting.ExpireTime) * time.Hour)
	// expireTime := nowTime.Add(24 * time.Hour)
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
