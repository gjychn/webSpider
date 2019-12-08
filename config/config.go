package config

import (
	"github.com/spf13/viper"
	"webSpider/log"
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
		log.Error("config is err")
		return nil, err
	}
	viper.Unmarshal(DefaultConfig)

	return  DefaultConfig, nil
}