package auth

import (
	"github.com/gofiber/fiber/v3"
	"sportTrackerAPI/internal/config"
	"sportTrackerAPI/pkg/Validate"
	"sportTrackerAPI/pkg/middleware"
)

type Handler struct {
	*Service
	*config.Config
}

func NewAuthHandler(service *Service, config *config.Config) *Handler {
	handler := Handler{
		Service: service,
		Config:  config,
	}
	return &handler
}

func (handler *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/test", middleware.AuthMiddleware(handler.Config), handler.Test)
	app.Get("/register", handler.Register)
	app.Get("/login", handler.Login)
	app.Get("/refresh", handler.Refresh)
	app.Get("/logout", middleware.AuthMiddleware(handler.Config), handler.Logout)
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
	return ctx.Status(fiber.StatusOK).JSON(RegisterResponse{Message: "Registration successful."})
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
	response, err := handler.Service.Login(
		request.Email,
		request.Password,
		handler.Config.Auth.Secret)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (handler *Handler) Refresh(ctx fiber.Ctx) error {
	var request RefreshRequest
	err := ctx.Bind().JSON(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := handler.Service.Refresh(request.RefreshToken, handler.Config.Auth.Secret)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (handler *Handler) Logout(ctx fiber.Ctx) error {
	email := ctx.Locals("userEmail").(string)
	err := handler.Service.Logout(email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
