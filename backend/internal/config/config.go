package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() Config {
	// Karena dijalankan dari cmd/api/, .env ada di ../
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	return Config{
		DBUrl: os.Getenv("DB_URL"),
		Port:  os.Getenv("PORT"),
	}
}
