package admin_service

import (
	"bufio"
	"encoding/base64"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/goinggo/mapstructure"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/setting"
)

type Driver struct {
}

type DriverCreateForm struct {
	ID         int               `json:"id"`
	Name       string            `json:"name" valid:"Required"`
	Phone      string            `json:"phone" valid:"Required;Mobile"`
	Age        int               `json:"age" valid:"Required";Max(70);Min(18)`
	DrivingAge int               `json:"driving_age" valid:"Required"`
	Image      map[string]string `json:"image" valid:"Required"`
}

func DriverList(data *ListParam) model.PageJson {
	drivers := make([]model.Driver, 0)
	var driverModel model.Driver
	// model.Db.Find().Count()
	if err := mapstructure.Decode(data.Filter, &driverModel); err != nil {
		return model.PageJson{}
	}
	Db := model.Db.Where(&driverModel).Limit(data.Limit).Offset((data.Page - 1) * data.Limit).Order(strings.Join(data.Sort[:], " "))
	Db.Find(&drivers)

	var total = 0
	model.Db.Model(&model.Route{}).Where(&driverModel).Count(&total)

	page := model.PageJson{
		Data:  drivers,
		Page:  data.Page,
		Total: total,
		Size:  len(drivers),
	}
	return page
}

func DriverCreate(data *DriverCreateForm) (*model.Driver, error) {
	image := strings.Split(data.Image["base64"], ",")[1]
	imgs, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return nil, errors.New("base64解码错误")
	}
	timenow := time.Now().Unix()
	filename := strconv.FormatInt(timenow, 10) + "." + data.Image["type"]
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
	driverModel := model.Driver{
		Name:       data.Name,
		Phone:      data.Phone,
		Image:      setting.AppSetting.PrefixUrl + setting.AppSetting.ImageSavePath + filename,
		DrivingAge: data.DrivingAge,
		Age:        data.Age,
	}
	if err := model.Db.Save(&driverModel).Error; err != nil {
		return nil, err
	}
	return &driverModel, nil
}

func DriverOne(id int) (*model.Driver, error) {
	driver := &model.Driver{}
	if err := model.Db.First(&driver, id).Error; err != nil {
		return nil, err
	}
	return driver, nil
}

func DriverEdit(data *DriverCreateForm) (*model.Driver, error) {
	oldDriver := &model.Driver{}
	model.Db.Find(&oldDriver, data.ID)
	var re = regexp.MustCompile(`\/(\d+\.\w+)$`)
	filename := re.FindAllString(oldDriver.Image, -1)[1]
	image := strings.Split(data.Image["base64"], ",")[1]
	imgs, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return nil, errors.New("base64解码错误")
	}
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

	driver := &model.Driver{
		Model:      model.Model{ID: uint(data.ID)},
		Name:       data.Name,
		Phone:      data.Phone,
		Age:        data.Age,
		DrivingAge: data.DrivingAge,
		Image:      setting.AppSetting.PrefixUrl + setting.AppSetting.ImageSavePath + filename,
	}
	if err := model.Db.Save(driver).Error; err != nil {
		return nil, err
	}
	return driver, nil
}
