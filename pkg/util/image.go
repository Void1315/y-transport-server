package util

import (
	"bufio"
	"encoding/base64"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/setting"
)

type UploadImage struct {
	Base64 string `json:"base64`
	Type   string `json:"type"`
	Name   string `json:"name"`
}

func SaveImage(images *[]UploadImage) (*[]model.SavedImage, error) {
	savedImages := make([]model.SavedImage, 0)
	for _, image := range *images {
		base64Str := strings.Split(image.Base64, ",")[1]
		imgs, err := base64.StdEncoding.DecodeString(base64Str)
		if err != nil {
			return nil, errors.New("base64解码错误")
		}
		timenow := time.Now().UnixNano()
		filename := strconv.FormatInt(timenow, 10) + "." + image.Type
		file, err2 := os.OpenFile("./static/img/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			return nil, errors.New("创建文件错误")
		}
		w := bufio.NewWriter(file) //创建新的 Writer 对象
		_, err3 := w.WriteString(string(imgs))
		if err3 != nil {
			return nil, errors.New("写入文件错误")
		}
		w.Flush()
		defer file.Close()
		savedImages = append(savedImages, model.SavedImage{
			Path:     setting.AppSetting.PrefixUrl + setting.AppSetting.ImageSavePath + filename,
			FileName: filename,
		})
	}
	return &savedImages, nil
}
func EditImage(images *[]UploadImage, oldImages *[]model.SavedImage) (*[]model.SavedImage, error) {
	for _, image := range *oldImages {
		path := setting.AppSetting.PrefixUrl + setting.AppSetting.ImageSavePath + image.FileName
		err := os.Remove(path)
		if err != nil {
			return nil, err
		}
	}
	savedImages, err1 := SaveImage(images)
	if err1 != nil {
		return nil, err1
	}
	return savedImages, nil
}
