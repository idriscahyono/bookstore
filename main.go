package main

import (
	"gin-bookstore/configs"
	"gin-bookstore/routers"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	configs.InitDB()
	routers.InitRouter()
}
