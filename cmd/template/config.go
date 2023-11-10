package template

func ConfigTemplate() []byte {
	return []byte(`package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        int ` + "`mapstructure:\"PORT\"`" + `
	Environment string ` + "`mapstructure:\"APP_ENV\"`" + `
}

const (
	LocalEnvironment      = "local"
	DevEnvironment        = "dev"
	ProductionEnvironment = "prod"
)

var AppConfig Config

func init() {
	viper.AddConfigPath("../../internal/config")
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		// If no config is present we immediately exit.
		panic(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		// If config cannot be unmarshalled we immediately exit.
		panic(err)
	}
}
`)
}

func AppEnvTemplate() []byte {
	return []byte(`PORT=8080
APP_ENV=local
`)
}
