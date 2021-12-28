package main

import (
	"gin-bookstore/app/configs"
	"gin-bookstore/app/routers"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	configs.InitDB()
	routers.InitRouter()
}
