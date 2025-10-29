package auth

import (
	"github.com/gofiber/fiber/v3"
	"net/http"
	"sportTrackerAPI/internal/config"
	"sportTrackerAPI/pkg/Validate"
	"sportTrackerAPI/pkg/jwt"
	"sportTrackerAPI/pkg/middleware"
	"time"
)

type Handler struct {
	*Service
	*config.Config
	*Repository
}

func NewAuthHandler(service *Service, config *config.Config, repo *Repository) *Handler {
	handler := Handler{
		Service:    service,
		Config:     config,
		Repository: repo,
	}
	return &handler
}

func (handler *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/test", middleware.AuthMiddleware(handler.Config), handler.Test)
	app.Get("/register", handler.Register)
	app.Get("/login", handler.Login)
	app.Get("/refresh", handler.Refresh)
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
	_, err = handler.Service.Login(
		request.Email,
		request.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	accessToken, refreshToken, accessExp, refreshExp, err := jwt.GenerateTokens(handler.Config.Auth.Secret, request.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = handler.Repository.Set(request.Email, refreshToken, time.Until(refreshExp))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(LoginResponse{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		AccessExpires:  accessExp,
		RefreshExpires: refreshExp,
	})

}

func (handler *Handler) Refresh(ctx fiber.Ctx) error {
	var request RefreshRequest
	err := ctx.Bind().JSON(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	isValid, claims := jwt.NewJWT(handler.Config.Auth.Secret).Parse(request.RefreshToken)
	if !isValid {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}
	accessToken, refreshToken, accessExp, refreshExp, err := jwt.GenerateTokens(handler.Config.Auth.Secret, claims.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = handler.Repository.Set(claims.Email, refreshToken, time.Until(refreshExp))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(RefreshResponse{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		AccessExpires:  accessExp,
		RefreshExpires: refreshExp,
	})
}
