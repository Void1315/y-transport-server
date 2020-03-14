package util

import (
	"crypto/sha256"
	"encoding/hex"
)

//Make 加密
func Make(text string, salt string) (string, error) {
	md5 := sha256.New()
	md5.Write([]byte(text))
	md5.Write([]byte(salt))
	bs := md5.Sum(nil)
	return hex.EncodeToString(bs), nil
}
