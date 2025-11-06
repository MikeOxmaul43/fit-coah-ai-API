package exercise

import "github.com/gofiber/fiber/v3"

func (handler *Handler) RegisterRoutes(router fiber.Router) {
	routes := router.Group("exercise")

	routes.Get("/", handler.GetAll)
	routes.Get("/:muscleGroup", handler.GetByMuscleGroup)

}
