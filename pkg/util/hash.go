package util

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//GetSalt 获取指定长度加密盐
func GetSalt(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//Make 加密
func Make(text string, salt string) string {
	md5 := sha256.New()
	md5.Write([]byte(text))
	md5.Write([]byte(salt))
	bs := md5.Sum(nil)
	return hex.EncodeToString(bs)
}
