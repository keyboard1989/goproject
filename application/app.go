package application

import (
	"net/http"
	"time"

	"github.com/keyboard1989/goproject/models"
	"github.com/keyboard1989/goutil"
)

type App struct {
	config     *models.Configuration
	statis     *goutil.Statistics
	httpClient *http.Client
}

func NewApp() *App {
	config := GetConfig()

	statis := goutil.NewStatistics(config.Statis.ItemLength, config.Statis.DataLength, config.Statis.TimeFormat, GetTime)

	//初始化发送请求的httpClient
	timeOut := time.Duration(config.HTTPClient.TimeOut) * time.Millisecond
	maxIdleConnsPerHost := config.HTTPClient.MaxIdleConnsPerHost
	httpClient := initHTTPClient(timeOut, maxIdleConnsPerHost)

	app := &App{
		config:     config,
		statis:     statis,
		httpClient: httpClient,
	}

	return app
}
