package controllers

import (
	"golang_app/golangApp/app/services"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v2"
)

type ImageController struct {
	service     *services.CloudinaryService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewCloudinaryController(localize *localize.Localization, redis *redis.RedisClient) *ImageController {
	service := services.NewCloudinaryService(localize, redis)
	return &ImageController{service: service, translation: localize, redis: redis}
}

func (ctrl *ImageController) SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		c.STATUS: c.SUCCESS,
		c.DATA:   data,
	}
}

func (ctrl *ImageController) ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}

func (ctrl *ImageController) UploadFile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var uploadResp *services.UploadImageResponse
		form, err := ctx.MultipartForm()
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}

		files, fileExists := form.File["file"]
		if !fileExists || len(files) == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(ctrl.ErrorResponse(c.MESSAGE_ERROR_FILE_PARAMS_REQUIRED))
		}

		file, err := files[0].Open()
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.Localization(c.FAILED_OPEN_FILE)))
		}

		defer file.Close()
		fileName := files[0].Filename
		directory := form.Value["directory"][0]

		if directory != "" {
			uploadResp, err = ctrl.service.UploadImageToFolder(file, fileName, directory)
		} else {
			uploadResp, err = ctrl.service.UploadImage(file, fileName)
		}
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		return ctx.JSON(ctrl.SuccessResponse(uploadResp))
	}
}
