package main

import (
	"github.com/gofiber/fiber/v3"
	"sportTrackerAPI/internal/auth"
	"sportTrackerAPI/internal/config"
	"sportTrackerAPI/internal/user"
	"sportTrackerAPI/pkg/db"
)

const HttpPort = ":8080"

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()
	database := db.NewDb(cfg)

	//Repositories
	userRepository := user.NewUserRepository(database)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handlers
	authHandler := auth.NewAuthHandler(authService, cfg)

	//RegisterRoutes
	authHandler.RegisterRoutes(app)

	err := app.Listen(HttpPort)
	if err != nil {
		panic(err)
	}
}
