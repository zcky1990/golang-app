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

type WeddingController struct {
	service     *services.WeddingService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewWeddingController(database *mongo.Database, localize *localize.Localization, redis *redis.RedisClient) *WeddingController {
	service := services.NewWeddingService(database, localize, redis)
	return &WeddingController{service: service, translation: localize, redis: redis}
}

func (cntrl *WeddingController) SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		c.STATUS: c.SUCCESS,
		c.DATA:   data,
	}
}

func (cntrl *WeddingController) ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}

func (cntrl *WeddingController) CreateWeddingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.WeddingData
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(cntrl.ErrorResponse(err.Error()))
		}
		data, err := cntrl.service.CreateWeddingData(params)
		if err != nil {
			return ctx.JSON(cntrl.ErrorResponse(cntrl.translation.Localization(c.EMAIL_TAKEN)))
		}
		return ctx.JSON(cntrl.SuccessResponse(data))
	}
}

func (cntrl *WeddingController) GetWeddingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		weddingID := ctx.Params("id")

		data, err := cntrl.service.GetWeddingDataById(weddingID)
		if err != nil {
			return ctx.JSON(cntrl.ErrorResponse(cntrl.translation.Localization(c.EMAIL_TAKEN)))
		}
		return ctx.JSON(cntrl.SuccessResponse(data))
	}
}
