package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/y-transport-server/migration"
	"github.com/y-transport-server/model"
	"github.com/y-transport-server/pkg/logging"
	"github.com/y-transport-server/pkg/setting"
	"github.com/y-transport-server/router"
)

func init() {
	setting.Setup()
	model.Setup()
	migration.Setup(model.Db) // 第一次运行 迁移表结构
	logging.Setup()
	// util.Setup()

}
func main() {
	r := router.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening http://127.0.0.1%s", endPoint)

	server.ListenAndServe()
}
