package providers

import (

	"os"

	"github.com/joho/godotenv"
)


// IConfig defines an interface for retrieving configuration values.
type IConfig interface {
	Get(key string) string
}
// Config implements the IConfig interface by loading values from the
// environment and making them available via the Get method.
type Config struct{}

func NewConfig() IConfig {
	return &Config{}
}

// Get retrieves the value for the given key from the environment.
// It loads the values from a .env file if present, otherwise
// it falls back to retrieving the value from the OS environment.
func (cfg *Config) Get(key string) string {
	// TODO: Comment the err section on deployment. It's only for development use
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
		// log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
