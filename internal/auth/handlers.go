package auth

import (
	"github.com/aliyilmazdev/todo-list-restful-api/internal/user"
	"github.com/aliyilmazdev/todo-list-restful-api/pkg/presenter"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
	var user user.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(presenter.ErrorResponse(err))
	}

	res, err := service.Register(user)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(presenter.ErrorResponse(err))
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(presenter.SuccessResponse(res, "User registered successfully."))

	}
}

func LoginHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
	var request LoginRequest
	if err := c.BodyParser(&request); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(presenter.ErrorResponse(err))
	}

	res, err := service.Login(request)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(presenter.ErrorResponse(err))
	}

	c.Status(fiber.StatusOK)
	return c.JSON(presenter.SuccessResponse(res, "User logged in successfully."))
	}
}