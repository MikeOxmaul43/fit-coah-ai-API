package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sportTrackerAPI/internal/dayExercise"
	"sportTrackerAPI/internal/exercise"
	"sportTrackerAPI/internal/program"
	"sportTrackerAPI/internal/programDay"
	"sportTrackerAPI/internal/user"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&exercise.Exercise{})
	db.AutoMigrate(&dayExercise.DayExercise{})
	db.AutoMigrate(&programDay.ProgramDay{})
	db.AutoMigrate(&program.Program{})
}
