//config
package config

//convig
import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl   string
	Port    string
	BaseURL string // ← tambahkan ini
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}
	return Config{
		DBUrl:   os.Getenv("DB_URL"),
		Port:    os.Getenv("PORT"),
		BaseURL: os.Getenv("BASE_URL"),
	}
}
