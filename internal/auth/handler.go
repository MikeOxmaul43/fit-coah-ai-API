package auth

import (
	"github.com/gofiber/fiber/v3"
	"sportTrackerAPI/pkg/Validate"
)

type Handler struct {
	*Service
}

func NewAuthHandler(service *Service) *Handler {
	handler := Handler{service}
	return &handler
}

func (handler *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/test", handler.Test)
	app.Get("/register", handler.Register)
}
func (handler *Handler) Test(ctx fiber.Ctx) error {
	return ctx.SendString("HelloWorld")
}

func (handler Handler) Register(ctx fiber.Ctx) error {
	var request RegisterRequest
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
	_, err = handler.Service.Register(
		request.Email,
		request.Password,
		request.UserName)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).SendString("Successful")
}

func (handler *Handler) Login(ctx fiber.Ctx) error {
	var request LoginRequest
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
	_, err = handler.Service.Login(
		request.Email,
		request.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).SendString("Successful")

}
