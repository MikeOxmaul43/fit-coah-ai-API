package programDay

import (
	"gorm.io/gorm"
	"sportTrackerAPI/internal/dayExercise"
)

type ProgramDay struct {
	gorm.Model
	ProgramID   uint `gorm:"not null;index"`
	DayNumber   int  `gorm:"not null"`
	Description string
	Exercises   []dayExercise.DayExercise `gorm:"constraint:OnDelete:CASCADE"`
}
