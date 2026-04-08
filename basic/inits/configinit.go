package inits

import (
	"ShengXian/basic/config"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigFile("C:\\Users\\35305\\Desktop\\ShengXian\\ShengXian\\config.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&config.GlobalConf)
}
