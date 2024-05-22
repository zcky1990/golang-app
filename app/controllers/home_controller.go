package controllers

import (
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

var _ BaseController = (*HomeController)(nil)

type HomeController struct {
	jsImportPath  string
	cssImportPath string
	translation   *localize.Localization
	redis         *redis.RedisClient
}

type response struct {
	JsImportPath  string
	CssImportPath string
}

func NewHomeController(JsPath string, CssPath string, localize *localize.Localization, redis *redis.RedisClient) *HomeController {
	return &HomeController{jsImportPath: JsPath, cssImportPath: CssPath, translation: localize, redis: redis}
}

func (c *HomeController) structToMap(v interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Ensure v is a struct or pointer to struct
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // Dereference pointer to get the underlying struct value
	} else if val.Kind() != reflect.Struct {
		panic("Input must be a struct or pointer to struct")
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i).Interface()
		// Check if the field is a nested struct
		if reflect.TypeOf(fieldValue).Kind() == reflect.Struct {
			result[field.Name] = c.structToMap(fieldValue)
		} else {
			result[field.Name] = fieldValue
		}
	}
	return result
}

func (c *HomeController) IndexPage() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response := &response{
			JsImportPath:  c.jsImportPath,
			CssImportPath: c.cssImportPath,
		}
		return ctx.Render("index", fiber.Map{
			"JSPath":     c.jsImportPath,
			"CSSPath":    c.cssImportPath,
			"JSFileName": "home.js",
			"Title":      "Hello, World!",
			"Data":       c.structToMap(response),
		}, "layouts/application")
	}
}
