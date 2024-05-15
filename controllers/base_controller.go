package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type BaseController interface {
	SuccessResponse(data interface{}) fiber.Map
	ErrorResponse(message string) fiber.Map
}
