package middleware

import (
	"github.com/gofiber/fiber/v3"
	"net/http"
	"sportTrackerAPI/internal/config"
	"sportTrackerAPI/pkg/jwt"
)

func AuthMiddleware(config *config.Config) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Token Required"})
		}
		if len(token) < 7 || token[:7] != "Bearer " {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		isValid, claims := jwt.NewJWT(config.Auth.Secret).Parse(token[7:])
		if !isValid {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		ctx.Locals("userEmail", claims.Email)
		return ctx.Next()
	}
}
