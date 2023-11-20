package app

import (
	"flag"
	"log"
	"os"

	"github.com/heenzu/Go-Commerce/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
 	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error bang pas loading .env")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoCommerce")
	appConfig.AppEnv = getEnv("APP_ENV", "Development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "user")
	dbConfig.DBPassword = getEnv("DB_Password", "password")
	dbConfig.DBName = getEnv("DB_NAME", "dbname")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Intialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

}