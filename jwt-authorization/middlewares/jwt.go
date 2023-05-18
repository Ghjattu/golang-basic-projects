package middlewares

import (
	"jwt-authorization/utils"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthorization() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := utils.ValidateToken(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Next()
	}
}
