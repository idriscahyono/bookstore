package main

import (
	"github.com/idriscahyono/bookstore/app/configs"
	"github.com/idriscahyono/bookstore/app/routers"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	configs.InitDB()
	routers.InitRouter()
}
