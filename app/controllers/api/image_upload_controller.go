package api

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

func (cntrl *ImageController) SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		c.STATUS: c.SUCCESS,
		c.DATA:   data,
	}
}

func (cntrl *ImageController) ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}

func (cntrl *ImageController) getLanguange(ctx *fiber.Ctx) string {
	lang := ctx.Get("Accept-Language")
	if lang == "" {
		return c.LOCALE_ENGLISH
	}
	if lang == c.LOCALE_ENGLISH {
		return c.LOCALE_ENGLISH
	} else {
		return c.LOCALE_INDONESIA
	}
}

func (cntrl *ImageController) UploadFile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var uploadResp *services.UploadImageResponse
		form, err := ctx.MultipartForm()
		if err != nil {
			return ctx.JSON(cntrl.ErrorResponse(err.Error()))
		}

		files, fileExists := form.File["file"]
		if !fileExists || len(files) == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(cntrl.ErrorResponse(cntrl.translation.GetLocalizationMessageWithLocale("FILE_PARAMS_REQUIRED", cntrl.getLanguange(ctx))))
		}

		file, err := files[0].Open()
		if err != nil {
			return ctx.JSON(cntrl.ErrorResponse(cntrl.translation.GetLocalizationMessageWithLocale("FAILED_OPEN_FILE", cntrl.getLanguange(ctx))))
		}

		defer file.Close()
		fileName := files[0].Filename
		directory := form.Value["directory"][0]

		if directory != "" {
			uploadResp, err = cntrl.service.UploadImageToFolder(file, fileName, directory)
		} else {
			uploadResp, err = cntrl.service.UploadImage(file, fileName)
		}
		if err != nil {
			return ctx.JSON(cntrl.ErrorResponse(err.Error()))
		}
		return ctx.JSON(cntrl.SuccessResponse(uploadResp))
	}
}
