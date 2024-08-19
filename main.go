package main

import (
	"log"

	"github.com/aliyilmazdev/todo-list-restful-api/database"
	"github.com/aliyilmazdev/todo-list-restful-api/internal/auth"
	"github.com/aliyilmazdev/todo-list-restful-api/internal/user"
	"github.com/aliyilmazdev/todo-list-restful-api/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	if database.DB == nil {
		log.Fatal("Database connection is not initialized")
	}

	jwtClient := jwt.NewClient()

	api := app.Group("/api", logger.New())

	userRepository := user.NewRepository(database.DB)
	userService := user.NewService(userRepository)
	user.SetupUserRoutes(api, userService)

	authService := auth.NewService(userService, jwtClient)
	auth.SetupAuthRoutes(api, authService)

	app.Listen(":3000")
}