package picture

import (
	"bookmark/logger"
	"github.com/spf13/viper"
)

var (
	pictureConfig *viper.Viper
	keywordMap    map[string][]string
)

func init() {
	pictureConfig = viper.New()
	//workDir, _ := os.Getwd()
	pictureConfig.AddConfigPath("/home/ubuntu/bookmark/")
	pictureConfig.SetConfigType("yaml")
	pictureConfig.SetConfigName("pictureSource")
	err := pictureConfig.ReadInConfig()
	if err != nil {
		logger.SugarLogger.Fatal("Error when reading picture source. Error: ", err.Error())
	}
	keywordMap = pictureConfig.GetStringMapStringSlice("keyword")
}
