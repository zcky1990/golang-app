package controllers

import (
	"context"
	constant "golang_app/golangApp/constants"
	"golang_app/golangApp/services"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gofiber/fiber/v2"
)

var ctx context.Context
var cld *cloudinary.Cloudinary

type ImageController struct {
	service     *services.CloudinaryService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewCloudinaryController(localize *localize.Localization, redis *redis.RedisClient) *ImageController {
	service := services.NewCloudinaryService(localize, redis)
	return &ImageController{service: service, translation: localize, redis: redis}
}

func (c *ImageController) SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		constant.STATUS: constant.SUCCESS,
		constant.DATA:   data,
	}
}

func (c *ImageController) ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		constant.STATUS:        constant.FAILED,
		constant.ERROR_MESSAGE: message,
	}
}

func (c *ImageController) UploadFile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var uploadResp *services.UploadImageResponse
		form, err := ctx.MultipartForm()
		if err != nil {
			return ctx.JSON(c.ErrorResponse(err.Error()))
		}

		files, fileExists := form.File["file"]
		if !fileExists || len(files) == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(c.ErrorResponse("File Params is required"))
		}

		file, err := files[0].Open()
		if err != nil {
			return ctx.JSON(c.ErrorResponse(c.translation.Localization(constant.FAILED_OPEN_FILE)))
		}

		defer file.Close()
		fileName := files[0].Filename
		directory := form.Value["directory"][0]

		if directory != "" {
			uploadResp, err = c.service.UploadImageToFolder(file, fileName, directory)
		} else {
			uploadResp, err = c.service.UploadImage(file, fileName)
		}
		if err != nil {
			return ctx.JSON(c.ErrorResponse(err.Error()))
		}
		return ctx.JSON(c.SuccessResponse(uploadResp))
	}
}
