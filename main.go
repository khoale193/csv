package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/khoale193/csv/cron"
	"github.com/khoale193/csv/models/dbcon"
	"github.com/khoale193/csv/models/migration"
	"github.com/khoale193/csv/pkg/logging"
	"github.com/khoale193/csv/pkg/setting"
	"github.com/khoale193/csv/routers"
)

func init() {
	setting.Setup()

	dbcon.Setup()
	migration.Migrate()

	logging.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/khoale193/csv
// @license.name MIT
// @license.url https://github.com/khoale193/csv/LICENSE
// @host localhost:8003/api
// @in header
// @name Authorization
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	cron.InitRouter()
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] currently server time %s\n", time.Now())
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
