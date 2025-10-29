package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DbСonfig
	Auth AuthConfig
	Rdb  RdbConfig
}

type DbСonfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}
type RdbConfig struct {
	Address string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env")
	}
	return &Config{
		Db:   DbСonfig{Dsn: os.Getenv("DSN")},
		Auth: AuthConfig{Secret: os.Getenv("SECRET")},
		Rdb:  RdbConfig{Address: os.Getenv("REDIS")},
	}
}
