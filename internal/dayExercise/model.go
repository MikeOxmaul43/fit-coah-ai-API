package dayExercise

import (
	"gorm.io/gorm"
	"sportTrackerAPI/internal/exercise"
)

type DayExercise struct {
	gorm.Model
	ProgramDayID uint `gorm:"not null;index"`
	ExerciseID   uint `gorm:"not nuLL;index"`
	Exercise     exercise.Exercise
	Sets         int `gorm:"not null;default:1;check:sets >= 1"`
	Reps         int `gorm:"not null;default:1;check:reps >= 1"`
	OrderIndex   int `gorm:"not null;default:0;index"`
}
