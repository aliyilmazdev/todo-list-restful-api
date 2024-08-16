package main

import (
	"github.com/aliyilmazdev/todo-list-restful-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	api := app.Group("/api")

	app.Listen(":3000")
}