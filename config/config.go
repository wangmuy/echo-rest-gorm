package config

import (
	"github.com/tkanos/gonfig"
)

type Config struct {
	DB_DIALECT        string
	DB_CONN           string
	DB_MAX_IDLE_CONNS int
	DB_MAX_OPEN_CONNS int
}

func GetConfig() *Config {
	c := Config{}
	gonfig.GetConf("config/config.json", &c)
	return &c
}
