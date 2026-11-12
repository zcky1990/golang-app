package api

import (
	"github.com/gofiber/fiber/v3"
)

// BaseApiController is the shared contract of the JSON controllers.
//
// Fiber v3 handlers receive fiber.Ctx (an interface) instead of *fiber.Ctx.
type BaseApiController interface {
	getLanguange(ctx fiber.Ctx) string
	SuccessResponse(data interface{}) fiber.Map
	ErrorResponse(message string) fiber.Map
}
