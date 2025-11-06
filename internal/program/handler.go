package program

import (
	"github.com/gofiber/fiber/v3"
	"sportTrackerAPI/pkg/Validate"
	"strconv"
)

type Handler struct {
	*Service
}

func NewProgramHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (handler *Handler) GetProgramById(ctx fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id not uint",
		})
	}

	program, err := handler.Service.Repository.GetById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(GetByIdRequest{Program: *program})
}

func (handler *Handler) Create(ctx fiber.Ctx) error {
	var request CreateProgramRequest
	err := ctx.Bind().JSON(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = Validate.IsValid(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userEmail := ctx.Locals("userEmail").(string)
	err = handler.Service.Create(request, userEmail)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Program created successfully"})
}
