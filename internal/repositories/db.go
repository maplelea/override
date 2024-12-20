package repositories

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"override/config"
)

var DB *gorm.DB

func InitDB() {
	dsn := config.Viper.GetString("database.dsn")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(any("failed to connect database"))
	}
}
