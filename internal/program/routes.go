package program

import "github.com/gofiber/fiber/v3"

func (handler *Handler) RegisterRoutes(router fiber.Router, isAuth fiber.Handler) {
	routes := router.Group("/program")
	routes.Get("/:id", isAuth, handler.GetProgramById)
	routes.Post("/create", isAuth, handler.Create)
}
