package providers

import (
	"github.com/godo"
)

type IConfig interface {
	Get(key string) string
}

type Config struct {}

func NewConfig() IConfig {
	return &Config{}
}

func (cfg *Config) Get(key string) string {
	err := g
}