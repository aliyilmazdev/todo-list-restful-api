package auth

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router, service Service, validate *validator.Validate) {

	t := router.Group("/auth")
	t.Post("/register", RegisterHandler(service, validate))
	t.Post("/login", LoginHandler(service))
}