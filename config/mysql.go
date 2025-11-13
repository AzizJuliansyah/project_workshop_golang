package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	DB_USER := viper.GetString("DATABASE.USER")
	DB_PASS := viper.GetString("DATABASE.PASS")
	DB_HOST := viper.GetString("DATABASE.HOST")
	DB_PORT := viper.GetString("DATABASE.PORT")
	DB_NAME := viper.GetString("DATABASE.NAME")

	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true&loc=Local"
	mysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return mysql, nil
}