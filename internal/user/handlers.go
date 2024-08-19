package user

import (
	"errors"
	"strconv"

	"github.com/aliyilmazdev/todo-list-restful-api/pkg/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserByIDHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Locals("user") == nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(presenter.ErrorResponse(errors.New("Unauthorized")))
		}

		rawUser := c.Locals("user").(*jwt.Token)
		userClaims := rawUser.Claims.(jwt.MapClaims)
		userId := userClaims["id"].(float64)
		s := strconv.FormatFloat(userId, 'f', -1, 64)

		u, err := service.GetUserByID(s)

		if err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(presenter.ErrorResponse(err))
		}

		return c.JSON(presenter.SuccessResponse(u, "User loaded successfully."))
	}
}