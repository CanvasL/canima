package setting

import (
	"github.com/spf13/viper"
	"fmt"
)

var Config = new(AppConfig)

func Init() (err error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err = viper.Unmarshal(Config); err != nil {
		fmt.Println("viper.Unmarshal failed, err:", err)
	}

	return
}