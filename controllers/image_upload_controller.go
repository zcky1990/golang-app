package controllers

import (
	"context"
	"golang_app/golangApp/constant"
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

func NewCloudinaryController(service *services.CloudinaryService, locale *localize.Localization, redis *redis.RedisClient) *ImageController {
	return &ImageController{service: service, translation: locale, redis: redis}
}

func (c *ImageController) UploadFile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var uploadResp *services.UploadImageResponse
		form, err := ctx.MultipartForm()
		if err != nil {
			return ctx.JSON(ErrorResponse(err.Error()))
		}

		files, fileExists := form.File["file"]
		if !fileExists || len(files) == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse("File Params is required"))
		}

		file, err := files[0].Open()
		if err != nil {
			return ctx.JSON(ErrorResponse(c.translation.Localization(constant.FAILED_OPEN_FILE)))
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
			return ctx.JSON(ErrorResponse(err.Error()))
		}
		return ctx.JSON(SuccessResponse(uploadResp))
	}
}
