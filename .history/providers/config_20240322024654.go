package providers

import (
	// "log"
	"os"

	// "github.com/joho/godotenv"
)

type IConfig interface {
	Get(key string) string
}

type Config struct{}

func NewConfig() IConfig {
	return &Config{}
}

func (cfg *Config) Get(key string) string {
	// TODO: Comment the err section on deployment. It's only for
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
