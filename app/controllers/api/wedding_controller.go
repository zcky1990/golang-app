package api

import (
	"golang_app/golangApp/app/models"
	"golang_app/golangApp/app/services"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ BaseApiController = (*WeddingController)(nil)

type WeddingController struct {
	service     *services.WeddingService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewWeddingController(database *mongo.Database, localize *localize.Localization, redis *redis.RedisClient) *WeddingController {
	service := services.NewWeddingService(database, localize, redis)
	return &WeddingController{service: service, translation: localize, redis: redis}
}

func (ctrl *WeddingController) SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		c.STATUS: c.SUCCESS,
		c.DATA:   data,
	}
}

func (ctrl *WeddingController) ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}

func (ctrl *WeddingController) getLanguange(ctx *fiber.Ctx) string {
	locale := ctx.Get("Accept-Language")
	if locale == "" {
		return c.LOCALE_ENGLISH
	}
	if locale == c.LOCALE_ENGLISH {
		return c.LOCALE_ENGLISH
	} else {
		return c.LOCALE_INDONESIA
	}
}

func (ctrl *WeddingController) CreateWeddingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.WeddingData
		locale := ctrl.getLanguange(ctx)
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		data, err := ctrl.service.CreateWeddingData(params)
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("EMAIL_TAKEN", locale)))
		}
		return ctx.JSON(ctrl.SuccessResponse(data))
	}
}

func (ctrl *WeddingController) GetWeddingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		weddingID := ctx.Params("id")
		locale := ctrl.getLanguange(ctx)
		data, err := ctrl.service.GetWeddingDataById(weddingID)
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("EMAIL_TAKEN", locale)))
		}
		return ctx.JSON(ctrl.SuccessResponse(data))
	}
}
