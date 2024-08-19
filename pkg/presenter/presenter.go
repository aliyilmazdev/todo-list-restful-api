package presenter

import "github.com/gofiber/fiber/v2"

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"message": err.Error(),
	}	
}

func SuccessResponse(response interface{}, message string) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data": response,
		"message": message,
	}
}