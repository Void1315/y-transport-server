package util

import "github.com/y-transport-server/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.JwtSetting.JwtSecret)
}
