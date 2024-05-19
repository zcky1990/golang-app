package controllers

import (
	"golang_app/golangApp/app/models"
	"golang_app/golangApp/app/services"
	c "golang_app/golangApp/constants"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func (ctrl *WeddingController) CreateWeddingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.WeddingData
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		data, err := ctrl.service.CreateWeddingData(params)
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.Localization(c.EMAIL_TAKEN)))
		}
		return ctx.JSON(ctrl.SuccessResponse(data))
	}
}

func (ctrl *WeddingController) GetWeddingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		weddingID := ctx.Params("id")

		data, err := ctrl.service.GetWeddingDataById(weddingID)
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.Localization(c.EMAIL_TAKEN)))
		}
		return ctx.JSON(ctrl.SuccessResponse(data))
	}
}
