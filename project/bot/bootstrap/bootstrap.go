package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
)

func Env(envDir string) {
	err := godotenv.Load(envDir)
	if err != nil {
		log.Fatalln("loading .env file failed")
	}
}
