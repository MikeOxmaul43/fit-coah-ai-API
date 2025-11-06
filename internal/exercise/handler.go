package exercise

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	*Repository
}

func NewExerciseHandler(repository *Repository) *Handler {
	handler := Handler{
		Repository: repository,
	}
	return &handler
}

func (handler *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("exercise/", handler.GetAll)
	app.Get("exercise/:muscleGroup", handler.GetByMuscleGroup)

}

func (handler *Handler) GetAll(ctx fiber.Ctx) error {
	exercises, err := handler.Repository.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(GetAllResponse{Exercises: exercises})
}

func (handler *Handler) GetByMuscleGroup(ctx fiber.Ctx) error {
	muscleGroup := ctx.Params("muscleGroup")
	exercises, err := handler.Repository.GetByMuscleGroup(muscleGroup)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(GetByMuscleGroupResponse{Exercises: exercises})
}
