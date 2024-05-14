package controllers

import (
	"golang_app/golangApp/constant"
	"golang_app/golangApp/models"
	"golang_app/golangApp/services"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"

	"github.com/gofiber/fiber/v2"
)

type WeddingController struct {
	service     *services.WeddingService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewWeddingController(service *services.WeddingService, localize *localize.Localization, redis *redis.RedisClient) *WeddingController {
	return &WeddingController{service: service, translation: localize, redis: redis}
}

func (c *WeddingController) CreateWeddingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.WeddingData
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ErrorResponse(err.Error()))
		}
		data, err := c.service.CreateWeddingData(params)
		if err != nil {
			return ctx.JSON(ErrorResponse(c.translation.Localization(constant.EMAIL_TAKEN)))
		}
		return ctx.JSON(SuccessResponse(data))
	}
}

func (c *WeddingController) GetWeddingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		weddingID := ctx.Params("id")

		data, err := c.service.GetWeddingDataById(weddingID)
		if err != nil {
			return ctx.JSON(ErrorResponse(c.translation.Localization(constant.EMAIL_TAKEN)))
		}
		return ctx.JSON(SuccessResponse(data))
	}
}
