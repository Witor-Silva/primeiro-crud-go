package config

import (
	"os"

	"github.com/joho/godotenv"
)

type LoadedEnv struct {
	Environment string
	App         string
	ApiPort     string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
}

func LoadConfig() LoadedEnv {
	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error Loading ,env file")
		}

	}

	return LoadedEnv{
		Environment: os.Getenv("ENVIRONMENT"),
		App:         os.Getenv("APP"),
		ApiPort:     os.Getenv("API_PORT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPass:      os.Getenv("DB_PASS"),
		DBName:      os.Getenv("DB_NAME"),
	}
}
