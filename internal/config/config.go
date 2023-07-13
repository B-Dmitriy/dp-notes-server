package config

import (
	u "webservice/pgk/utils"

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

func InitConfig() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	config.Port = u.NilToStr(viper.Get("port"))
	config.Endpoint = u.NilToStr(viper.Get("endpoint"))
	config.Host = u.NilToStr(viper.Get("host"))
	config.User = u.NilToStr(viper.Get("user"))
	config.Password = u.NilToStr(viper.Get("password"))
	config.Schema = u.NilToStr(viper.Get("schema"))

	return config, nil
}
