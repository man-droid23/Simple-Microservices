package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/simple-microservices/pkg/auth"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
		})
	}
	tokenValidation, err := auth.ValidateToken(authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
		})
	}
	c.Locals("token", tokenValidation)
	return c.Next()
}
