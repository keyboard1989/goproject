package application

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/keyboard1989/goproject/models"
	"github.com/spf13/viper"
)

func getConfig() *models.Configuration {
	var config *models.Configuration
	viper.SetEnvPrefix("goproject")
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.SetDefault("conf", "./conf.yml")

	conf := viper.GetString("conf")
	spew.Dump(conf)
	file, _ := os.Open(conf)
	defer file.Close()

	err := viper.ReadConfig(file)
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
