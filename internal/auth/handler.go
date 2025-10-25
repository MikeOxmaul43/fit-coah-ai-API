package auth

import (
	"github.com/gofiber/fiber/v3"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"sportTrackerAPI/internal/config"
	"sportTrackerAPI/pkg/Validate"
	"sportTrackerAPI/pkg/jwt"
	"sportTrackerAPI/pkg/middleware"
	"time"
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
	expiredAt := time.Now().Add(24 * time.Hour)
	claims := jwt.Claims{
		Email:            request.Email,
		RegisteredClaims: jwtLib.RegisteredClaims{ExpiresAt: jwtLib.NewNumericDate(expiredAt), IssuedAt: jwtLib.NewNumericDate(time.Now())},
	}
	token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(claims)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(LoginResponse{Token: token, Expires: expiredAt})

}
