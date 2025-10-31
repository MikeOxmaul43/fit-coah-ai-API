package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"sportTrackerAPI/db"
	"sportTrackerAPI/internal/auth"
	"sportTrackerAPI/internal/config"
	"sportTrackerAPI/internal/exercise"
	"sportTrackerAPI/internal/user"
	"sportTrackerAPI/redisDb"
)

const HttpPort = ":8080"

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()
	database := db.NewDb(cfg)
	redisDataBase := redisDb.NewRDb(cfg)

	//Middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip} ${status} - ${method} ${path} ${latency}\n",
	}))

	//Repositories
	userRepository := user.NewUserRepository(database)
	authRedisRepository := auth.NewAuthRepository(redisDataBase)
	exerciseRepository := exercise.NewExerciseRepository(database)

	//Services
	authService := auth.NewAuthService(userRepository, authRedisRepository)
	//Handlers
	authHandler := auth.NewAuthHandler(authService, cfg)
	exerciseHandler := exercise.NewExerciseHandler(exerciseRepository)

	//RegisterRoutes
	authHandler.RegisterRoutes(app)
	exerciseHandler.RegisterRoutes(app)

	err := app.Listen(HttpPort)
	if err != nil {
		panic(err)
	}
}
