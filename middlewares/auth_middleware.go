package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{
		Key: []byte(secret),
	},		
	ErrorHandler: authMiddlewareError,
	})
}

func authMiddlewareError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"message": "Unauthorized",
	})
}