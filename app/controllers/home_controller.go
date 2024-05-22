package controllers

import (
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"

	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	jsImportPath  string
	cssImportPath string
	translation   *localize.Localization
	redis         *redis.RedisClient
}

func NewHomeController(JsPath string, CssPath string, localize *localize.Localization, redis *redis.RedisClient) *HomeController {
	return &HomeController{jsImportPath: JsPath, cssImportPath: CssPath, translation: localize, redis: redis}
}

func (c *HomeController) IndexPage() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"JSPath":     c.jsImportPath,
			"CSSPath":    c.cssImportPath,
			"JSFileName": "home.js",
			"Title":      "Hello, World!",
		}, "layouts/application")
	}
}
