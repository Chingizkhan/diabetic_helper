package config

import (
	"diabetHelperTelegramBot/diabetHelper/utils/env"
	"log"
)

var (
	config *Config
)

type Config struct {
	Port      string `env:"PORT"`
	AllowedIp string `env:"ALLOWED_IP"`
	Postgres
}

func Init() {
	conf := &Config{}
	if err := env.Unmarshal(conf); err != nil {
		log.Fatalln(err.Error())
	}
	config = conf
}

func Get() *Config {
	return config
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	DbName   string `env:"POSTGRES_DB"`
}
