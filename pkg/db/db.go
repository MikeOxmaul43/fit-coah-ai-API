package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sportTrackerAPI/internal/config"
)

type Db struct {
	*gorm.DB
}

func NewDb(cfg *config.Config) *Db {
	db, err := gorm.Open(postgres.Open(cfg.Db.Dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return &Db{db}
}
