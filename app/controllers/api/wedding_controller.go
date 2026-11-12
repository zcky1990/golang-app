package api

import (
	"golang_app/golangApp/app/models"
	"golang_app/golangApp/app/services"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var _ BaseApiController = (*WeddingController)(nil)

type WeddingController struct {
	service     *services.WeddingService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewWeddingController(database *gorm.DB, localize *localize.Localization, redis *redis.RedisClient) *WeddingController {
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

func (ctrl *WeddingController) getLanguange(ctx fiber.Ctx) string {
	locale := ctx.Get("Accept-Language")
	if locale == "" {
		return ""
	}
	if locale == c.LOCALE_ENGLISH {
		return c.LOCALE_ENGLISH
	}
	return c.LOCALE_INDONESIA
}

func (ctrl *WeddingController) CreateWeddingData() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var params models.WeddingData
		locale := ctrl.getLanguange(ctx)
		// Fiber v3: BodyParser() became Bind().Body()
		if err := ctx.Bind().Body(&params); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(ctrl.ErrorResponse(err.Error()))
		}
		data, err := ctrl.service.CreateWeddingData(params, locale)
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("FAILED_CREATE_WEDDING_DATA", locale)))
		}
		return ctx.Status(fiber.StatusCreated).JSON(ctrl.SuccessResponse(data))
	}
}

func (ctrl *WeddingController) GetWeddingData() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		// Fiber v3: Params() accepts a default value, so no more empty-string
		// lookups when the route param is missing.
		weddingID := ctx.Params("id")
		locale := ctrl.getLanguange(ctx)
		if !services.ValidUUID(weddingID) {
			return ctx.Status(fiber.StatusBadRequest).JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("INVALID_WEDDING_ID", locale)))
		}
		data, err := ctrl.service.GetWeddingDataById(weddingID)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("WEDDING_DATA_NOT_FOUND", locale)))
		}
		return ctx.JSON(ctrl.SuccessResponse(data))
	}
}
