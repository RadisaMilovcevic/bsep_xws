package config

import "os"

type Config struct {
	Port          string
	AccountDBHost string
	AccountDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("ACCOUNT_SERVICE_PORT"),
		AccountDBHost: os.Getenv("ACCOUNT_DB_HOST"),
		AccountDBPort: os.Getenv("ACCOUNT_DB_PORT"),
	}
}
