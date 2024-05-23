package controllers

import (
	env "golang_app/golangApp/config/environments"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

var _ BaseController = (*HomeController)(nil)

type HomeController struct {
	translation *localize.Localization
	redis       *redis.RedisClient
	envConf     env.EnvironmentConfiguration
}

type response struct {
	JsImportPath  string
	CssImportPath string
}

func NewHomeController(env env.EnvironmentConfiguration, localize *localize.Localization, redis *redis.RedisClient) *HomeController {
	return &HomeController{
		envConf:     env,
		translation: localize,
		redis:       redis,
	}
}

func (c *HomeController) structToMap(v interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	} else if val.Kind() != reflect.Struct {
		panic("Input must be a struct or pointer to struct")
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i).Interface()
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
			JsImportPath:  c.envConf.GetJSFilePath(),
			CssImportPath: c.envConf.GetCSSFilePath(),
		}
		return ctx.Render("index", fiber.Map{
			"JSPath":     c.envConf.GetJSFilePath(),
			"CSSPath":    c.envConf.GetCSSFilePath(),
			"JSFileName": "home.js",
			"Title":      "Hello, World!",
			"Data":       c.structToMap(response),
		}, "layouts/application")
	}
}
