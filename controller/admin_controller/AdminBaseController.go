package admin_controller

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/y-transport-server/service/admin_service"
)

func bindListParam(c *gin.Context) (*admin_service.ListParam, error) {
	valid := validation.Validation{}
	sort := ""
	if arg := c.Query("sort"); arg != "" {
		sort = com.StrTo(arg).String()
		valid.Required(sort, "sort")
	}
	limit := 0
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
		valid.Required(limit, "limit")
	}
	page := 0
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
		valid.Required(page, "page")
	}
	filter := ""
	if arg := c.Query("filter"); arg != "" {
		filter = com.StrTo(arg).String()
		valid.Required(filter, "filter")
	}
	jsonData := []byte(`
    {
        "sort": ` + sort + `,
        "limit": ` + strconv.Itoa(limit) + `,
        "page": ` + strconv.Itoa(page) + `,
        "filter": ` + filter + `
	}`)
	listJSONData := &admin_service.ListParam{}
	err := json.Unmarshal(jsonData, &listJSONData)
	if err != nil {
		return nil, err
	}
	return listJSONData, nil
}
