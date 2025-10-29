package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"sportTrackerAPI/db"
	"sportTrackerAPI/internal/auth"
	"sportTrackerAPI/internal/config"
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
	authRepository := auth.NewAuthRepository(redisDataBase)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handlers
	authHandler := auth.NewAuthHandler(authService, cfg, authRepository)

	//RegisterRoutes
	authHandler.RegisterRoutes(app)

	err := app.Listen(HttpPort)
	if err != nil {
		panic(err)
	}
}
