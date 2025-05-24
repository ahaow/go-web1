package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenCons  int
	}
}

var Appconfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file, %v", err)
	}

	Appconfig = &Config{}

	if err := viper.Unmarshal(Appconfig); err != nil {
		log.Fatal("Error parsing config file, %v", err)
	}

	initDB()
	initRedis()
}
