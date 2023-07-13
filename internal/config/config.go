package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	Endpoint string
	Host     string
	User     string
	Password string
	Schema   string
}

func nilReplaceEmptyStr(arg interface{}) string {
	if arg == nil {
		return ""
	} else {
		return fmt.Sprint(arg)
	}
}

func InitConfig() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	config.Port = nilReplaceEmptyStr(viper.Get("port"))
	config.Endpoint = nilReplaceEmptyStr(viper.Get("endpoint"))
	config.Host = nilReplaceEmptyStr(viper.Get("host"))
	config.User = nilReplaceEmptyStr(viper.Get("user"))
	config.Password = nilReplaceEmptyStr(viper.Get("password"))
	config.Schema = nilReplaceEmptyStr(viper.Get("schema"))

	return config, nil
}
