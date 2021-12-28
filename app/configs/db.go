package configs

import (
	"fmt"
	"github.com/idriscahyono/bookstore/app/database/seeders"
	"github.com/idriscahyono/bookstore/app/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	User   string
	Passwd string
	Addr   string
	Port   int
	DBName string
}

var DB *gorm.DB
var err error

func InitDB() {
	db := config{
		User:   viper.Get("DB_USERNAME").(string),
		Passwd: viper.Get("DB_PASSWORD").(string),
		Addr:   viper.Get("BASE_URL").(string),
		Port:   viper.Get("DB_PORT").(int),
		DBName: viper.Get("DB_NAME").(string),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User, db.Passwd, db.Addr, db.Port, db.DBName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	//Migration
	DB.AutoMigrate(&models.Book{})

	//Seeders
	seeders.DBSeed(DB)
}
