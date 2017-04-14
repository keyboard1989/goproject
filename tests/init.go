package tests

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/keyboard1989/goproject/application"
	"github.com/spf13/viper"
)

var (
	service *application.Service
)

func init() {
	gin.SetMode(gin.TestMode)
	service = initService()
}

func initService() *application.Service {
	dir, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}

	defer dir.Close()
	confYmlContent := initConfYml()
	ioutil.WriteFile(dir.Name(), []byte(confYmlContent), 0666)
	viper.Set("conf", dir.Name())

	app := application.NewApp()
	service := application.NewService(app)

	return service
}

func initConfYml() string {
	return `go_version: 1.7
server:
  addr: 127.0.0.1
  port: 8099
db:
  db_source: root:@tcp(127.0.0.1:3306)/adx?parseTime=true&charset=utf8&loc=Local
statis:
  item_length: 10000
  data_length: 10
  time_format: 2006-01-02 15:04:05
http_client:
  time_out: 700
  max_idle_conns_per_host: 50
cache:
  redis_source: 127.0.0.1:6379 
`
}
