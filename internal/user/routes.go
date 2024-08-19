package user

import (
	"github.com/aliyilmazdev/todo-list-restful-api/config"
	"github.com/aliyilmazdev/todo-list-restful-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router, service Service) {
	jwt := middlewares.NewAuthMiddleware(config.Config("JWT_SECRET"))

	t := router.Group("/user",)
	t.Get("/", jwt, GetUserByIDHandler(service))
	
}