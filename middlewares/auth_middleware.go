package middlewares

import (
	"github.com/aliyilmazdev/todo-list-restful-api/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{
		Key: []byte(config.Config("SECRET")),
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