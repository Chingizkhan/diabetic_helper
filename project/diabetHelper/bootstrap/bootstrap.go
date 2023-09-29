package bootstrap

import (
	"diabetHelperTelegramBot/diabetHelper/storage/postgres"
	"diabetHelperTelegramBot/diabetHelper/utils"
	"github.com/joho/godotenv"
	"log"
)

func Migrate(dir string) {
	dbConn := postgres.GetDbConnection()

	files, err := utils.GetFiles(dir)
	if err != nil {
		log.Fatalf("migration failed! \n get file error: %v", err)
	}

	for _, fileName := range files {
		queryString := utils.ReadFile(dir + fileName)

		//log.Printf("(%d/%d) %v%v", key+1, len(files), dir, fileName)

		_, err = dbConn.Exec(queryString)
		if err != nil {
			log.Fatalf("migration failed! \n migration=%s, query error=%v\n", fileName, err)
		}
	}
}

func Env(envDir string) {
	err := godotenv.Load(envDir)
	if err != nil {
		log.Fatalln("loading .env file failed")
	}
}

func PingDbConnect() {
	postgres.GetDbConnection()
}
