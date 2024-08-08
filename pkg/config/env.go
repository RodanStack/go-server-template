package config

import "github.com/spf13/viper"

type Env struct {
	ServerPort         string
	MaxMultipartMemory int64
	Environment        string
}

func NewEnv() *Env {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.SetDefault("SERVER_PORT", "8080")
	const tenMB = 10 * 1024 * 1024

	viper.SetDefault("MAX_MULTIPART_MEMORY", tenMB)

	return &Env{
		ServerPort:         viper.GetString("SERVER_PORT"),
		MaxMultipartMemory: viper.GetInt64("MAX_MULTIPART_MEMORY"),
		Environment:        viper.GetString("ENVIRONMENT"),
	}
}
