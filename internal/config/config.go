package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Cfg struct {
	DbAddress string `mapstructure:"db_address"`
	DbLogin   string `mapstructure:"db_login"`
	DbPass    string `mapstructure:"db_pass"`
	DbPort    string `mapstructure:"db_port"`
	DbName    string `mapstructure:"db_name"`
	HttpPort  string `mapstructure:"http_port"`
}

func SetConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("HW")
	v.SetDefault("DB_ADDRESS", "localhost")
	v.SetDefault("DB_LOGIN", "postgres")
	v.SetDefault("DB_PASS", "postgres")
	v.SetDefault("DB_PORT", "5432")
	v.SetDefault("DB_NAME", "postgres")
	v.SetDefault("HTTP_PORT", "8082")

	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		fmt.Println(err)
	}

	return cfg
}
