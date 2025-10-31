package exercise

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	MuscleGroup string `gorm:"index"`
	Description string
	Instruction string
	Difficulty  int `gorm:"default:1;check:difficulty >= 1 AND difficulty <= 5"`
}
