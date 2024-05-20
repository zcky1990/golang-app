package controllers

import (
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"

	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewHomeController(localize *localize.Localization, redis *redis.RedisClient) *HomeController {
	return &HomeController{translation: localize, redis: redis}
}

func (cntrl *HomeController) IndexPage() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/application")
	}
}
