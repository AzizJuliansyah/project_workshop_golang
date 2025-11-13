package config

import (
	"github.com/spf13/viper"
)

func InitViper() error {
	viper.SetConfigName("app.conf.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}