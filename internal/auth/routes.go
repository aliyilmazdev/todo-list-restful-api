package auth

import "github.com/gofiber/fiber/v2"

func SetupAuthRoutes(router fiber.Router, service Service) {
	t := router.Group("/auth")
	t.Post("/register", RegisterHandler(service))
	t.Post("/login", LoginHandler(service))
}