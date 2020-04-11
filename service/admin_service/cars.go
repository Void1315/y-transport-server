package admin_service

import (
	// "bufio"

	"bufio"
	"encoding/base64"
	"errors"

	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goinggo/mapstructure"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/setting"
)

type UploadImage struct {
	Base64 string `json:"base64`
	Type   string `json:"type"`
	Name   string `json:"name"`
}

// type SavedImage struct {
// 	Path     string `json:"path"`
// 	FileName string `json:"file_name"`
// }
type CarCreateForm struct {
	ID       int           `json:"id"`
	Name     string        `json:"name" valid:"Required"`
	Phone    string        `json:"phone" valid:"Required;Mobile"`
	Type     int           `json:"type"  valid:"Required;"`
	Image    []UploadImage `json:"image" valid:"Required"`
	Capacity int           `json:"capacity"  valid:"Required;"`
	Number   string        `json:"number"  valid:"Required;"`
}

func CarList(data *ListParam) model.PageJson {
	cars := make([]model.Car, 0)
	var carModel model.Car
	if err := mapstructure.Decode(data.Filter, &carModel); err != nil {
		return model.PageJson{}
	}
	Db := model.Db.Where(&carModel).Limit(data.Limit).Offset((data.Page - 1) * data.Limit).Order(strings.Join(data.Sort[:], " "))
	Db.Find(&cars)

	var total = 0
	model.Db.Model(&model.Car{}).Where(&carModel).Count(&total)

	page := model.PageJson{
		Data:  cars,
		Page:  data.Page,
		Total: total,
		Size:  len(cars),
	}
	return page
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
func CarCreate(data *CarCreateForm) (*model.Car, error) {
	savedImages, err := SaveImage(&data.Image)
	if err != nil {
		return nil, err
	}
	// jsonByte, err1 := json.Marshal(savedImages)
	// if err1 != nil {
	// 	return nil, err1
	// }
	car := &model.Car{
		Name:     data.Name,
		Phone:    data.Phone,
		Type:     data.Type,
		Capacity: data.Capacity,
		Number:   data.Number,
		// Image:    string(jsonByte),
		Image: &model.SavedImageMap{
			SavedImage: *savedImages,
			Valid:      true,
		},
	}
	if err2 := model.Db.Save(car).Error; err2 != nil {
		return nil, err2
	}
	return car, nil
}

func CarEdit(data *CarCreateForm) (*model.Car, error) {
	oldCar := &model.Car{}
	oldImages := make([]model.SavedImage, 0)
	model.Db.Find(&oldCar, data.ID)
	// err := json.Unmarshal([]byte(oldCar.Image.SavedImage), &oldImages)
	// if err != nil {
	// 	return nil, err
	// }
	savedImages, err1 := EditImage(&data.Image, &oldImages)
	if err1 != nil {
		return nil, err1
	}
	// jsonByte, err1 := json.Marshal(savedImages)
	// if err1 != nil {
	// 	return nil, err1
	// }
	car := &model.Car{
		Model:    model.Model{ID: uint(data.ID)},
		Name:     data.Name,
		Phone:    data.Phone,
		Type:     data.Type,
		Capacity: data.Capacity,
		Number:   data.Number,
		Image: &model.SavedImageMap{
			SavedImage: *savedImages,
			Valid:      true,
		},
	}
	if err2 := model.Db.Save(car).Error; err2 != nil {
		return nil, err2
	}
	return car, nil
}

func CarOne(id int) (*model.Car, error) {
	car := &model.Car{}

	if err := model.Db.First(&car, id).Error; err != nil {
		return nil, err
	}
	return car, nil
}
