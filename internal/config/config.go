package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db DbСonfig
}

type DbСonfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env")
	}
	return &Config{
		Db: DbСonfig{Dsn: os.Getenv("DSN")},
	}
}
