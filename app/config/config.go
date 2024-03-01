package config

import "github.com/spf13/viper"

func InitConfig() {
	viper.AddConfigPath("./app/config/")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err!= nil {
		panic("Error while loading config file")
	}

	viper.AddConfigPath("./app/database/")
	viper.SetConfigName("datasource.mongo")
	viper.SetConfigType("json")

	err = viper.MergeInConfig()
	if err!= nil {
		panic("Error while loading database config file")
	}
}