package auth

import (
	"github.com/gofiber/fiber/v3"
)

func (handler *Handler) RegisterRoutes(router fiber.Router, isAuth fiber.Handler) {
	routes := router.Group("/auth")

	routes.Get("/test", isAuth, handler.Test)
	routes.Get("/register", handler.Register)
	routes.Get("/login", handler.Login)
	routes.Get("/refresh", handler.Refresh)
	routes.Get("/logout", isAuth, handler.Logout)
}
