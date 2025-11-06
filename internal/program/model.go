package program

import (
	"gorm.io/gorm"
	"sportTrackerAPI/internal/programDay"
)

type Program struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Description   string
	Level         string
	DurationWeeks int
	CreatorType   string                  `gorm:"not null;index"` //system or user
	CreatedBy     *uint                   `gorm:"index"`
	ProgramDays   []programDay.ProgramDay `gorm:"constraint:OnDelete:CASCADE"`
}
