package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       int32

	DbConnectionString string
	DbName             string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		PublicHost:         "localhost",
		Port:               3000,
		DbConnectionString: os.Getenv("MONGO_DB_URI"),
		DbName:             "MONKEYCOM_GO_DB",
	}
}
