package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"uniqueIndex"`
	Email          string `gorm:"uniqueIndex;not null"`
	HashedPassword string `gorm:"not null"`
	Name           string
	Gender         string `gorm:"index"`
	Height         float64
	Weight         float64
	BirthDate      time.Time `gorm:"type:date"`
}
