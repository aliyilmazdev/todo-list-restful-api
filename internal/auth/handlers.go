package auth

import (
	"github.com/aliyilmazdev/todo-list-restful-api/pkg/presenter"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(service Service, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
	var registerRequest RegisterRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(presenter.ErrorResponse(err))
	}

	if err := validate.Struct(registerRequest); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(presenter.ErrorResponse(err))
	}

	res, err := service.Register(registerRequest)

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