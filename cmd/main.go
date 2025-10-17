package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"sportTrackerAPI/internal/auth"
	"sportTrackerAPI/internal/config"
	"sportTrackerAPI/internal/user"
	"sportTrackerAPI/pkg/db"
)

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()
	database := db.NewDb(cfg)

	//Repositories
	userRepository := user.NewUserRepository(database)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handlers
	authHandler := auth.NewAuthHandler(authService)

	//RegisterRoutes
	authHandler.RegisterRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
