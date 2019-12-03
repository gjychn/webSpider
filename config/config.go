package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var DefaultConfig = new(config)

type config struct {
	Api  appInfo
	Log  logInfo
}

type appInfo struct {
	Url string	`yaml:"url"`
}

type logInfo struct {
	Url string `yaml:"url"`
}

func LoadConfig(url string) (*config, error){
	vip := viper.New()
	vip.SetConfigName("config")
	vip.SetConfigFile(url)
	err := vip.ReadInConfig()
	if err != nil {
		fmt.Println("123")
		//TODO  日志记录
		return nil, err
	}
	viper.Unmarshal(DefaultConfig)

	return  DefaultConfig, nil
}