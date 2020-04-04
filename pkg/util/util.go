package util

import (
	"regexp"

	"github.com/y-transport-server/pkg/setting"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.JwtSetting.JwtSecret)
}

//GetDomain 从Origin提取domain
func GetDomain(origin string) string {
	var reg = regexp.MustCompile(`^https?:\/\/([\w-.]+)(:\d+)?`)
	return reg.FindStringSubmatch(origin)[1]
}
