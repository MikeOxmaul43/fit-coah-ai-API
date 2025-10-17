package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"index"`
	Email          string
	HashedPassword string
	Name           string
	Height         float64
	Weight         float64
	BirthDate      time.Time `gorm:"type:date"`
}
