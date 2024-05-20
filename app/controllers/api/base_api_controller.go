package api

import (
	"github.com/gofiber/fiber/v2"
)

type BaseApiController interface {
	SuccessResponse(data interface{}) fiber.Map
	ErrorResponse(message string) fiber.Map
}
