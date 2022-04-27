package config

import (
	"bookmark/logger"
	"github.com/spf13/viper"
)

var GlobalConfig *viper.Viper

func InitConfig() {
	//workDir, _ := os.Getwd()
	GlobalConfig = viper.New()
	GlobalConfig.AddConfigPath("/home/ubuntu/bookmark/")
	GlobalConfig.SetConfigName("application")
	GlobalConfig.SetConfigType("yaml")
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		logger.SugarLogger.Fatal("Error when reading config. Error: ", err.Error())
	}
}
