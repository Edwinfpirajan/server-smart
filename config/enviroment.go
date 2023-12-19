package config

import (
	"log"
	"sync"

	"github.com/andresxlp/gosuite/config"
)

var (
	Cfg  Config
	Once sync.Once
)

type Config struct {
	Server Server `mapstructure:"server" validate:"required"`
	MainDb MainDb `mapstructure:"main_db" validate:"required"`
}

type Server struct {
	Port int `mapstructure:"port" validate:"required"`
}

type MainDb struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DbName   string `mapstructure:"name" validate:"required"`
}

func Environments() Config {
	Once.Do(func() {
		if err := config.SetEnvsFromFile(".env"); err != nil {
			log.Panicf("Error can't loaded .env file %#v", err)
		}
		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return Cfg
}
