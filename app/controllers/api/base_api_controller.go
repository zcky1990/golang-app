package api

import (
	"github.com/gofiber/fiber/v2"
)

type BaseApiController interface {
	getLanguange(ctx *fiber.Ctx) string
	SuccessResponse(data interface{}) fiber.Map
	ErrorResponse(message string) fiber.Map
}
