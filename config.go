package main

import (
	"os"

	"github.com/keyboard1989/goproject/models"

	"github.com/spf13/viper"
)

var config models.Configuration

func initConfig() {
	viper.SetEnvPrefix("goproject")
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	viper.SetDefault("conf", "./conf.yml")

	conf := viper.GetString("conf")
	file, _ := os.Open(conf)
	defer file.Close()

	viper.ReadConfig(file)

	viper.Unmarshal(&config)
}
