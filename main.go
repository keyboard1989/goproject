package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/keyboard1989/goproject/handlers"
	"github.com/keyboard1989/goutil"
)

var (
	statis     *goutil.Statistics
	httpClient *http.Client
)

func initApp() {
	statis = goutil.NewStatistics(config.Statis.ItemLength, config.Statis.DataLength, config.Statis.TimeFormat, GetTime)

	//初始化发送请求的httpClient
	timeOut := time.Duration(config.HTTPClient.TimeOut) * time.Millisecond
	maxIdleConnsPerHost := config.HTTPClient.MaxIdleConnsPerHost
	httpClient = initHTTPClient(timeOut, maxIdleConnsPerHost)
}

func main() {

	initConfig()
	initApp()

	gin.SetMode(gin.DebugMode)

	if err := engine().Run(config.Server.Addr + ":" + config.Server.Port); err != nil {
		fmt.Println("start error")
	}
}

func engine() *gin.Engine {
	engine := gin.New()

	apiHandler := handlers.NewApiHandler(httpClient, statis)
	cmdHandler := handlers.NewCmdHandler(statis)
	debugHandler := handlers.NewDebugHandler()
	statusHandler := handlers.NewStatusHandler(statis)

	engine.GET("/", defaultHandler)
	engine.POST("/api", apiHandler.Api())
	engine.GET("/cmd", cmdHandler.Cmd())
	engine.GET("/debug", debugHandler.Debug())
	engine.GET("/status", statusHandler.Status())

	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "not found url \n")
	})
	return engine
}

func defaultHandler(c *gin.Context) {
	c.String(http.StatusOK, "goproject\n")
}
