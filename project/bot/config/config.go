package config

import (
	"diabetHelperTelegramBot/bot/utils/env"
	"log"
)

var (
	config *Config
)

type Config struct {
	Port      string `env:"PORT"`
	AppEnv    string `env:"APP_ENV"`
	AllowedIp string `env:"ALLOWED_IP"`
	Token     string `env:"TELEGRAM_TOKEN"`
	Service
}

type Service struct {
	DiabetHelper string `env:"SERVICE_DIABET_HELPER"`
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
