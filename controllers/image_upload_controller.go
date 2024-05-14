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
		files := form.File["file"]

		file, err := files[0].Open()
		if err != nil {
			return ctx.JSON(ErrorResponse(c.translation.Localization(constant.FAILED_OPEN_FILE)))
		}
		defer file.Close()
		fileName := files[0].Filename
		folder := form.Value[constant.FOLDER][0]

		if folder != "" {
			uploadResp, err = c.service.UploadImageToFolder(file, fileName, folder)
		} else {
			uploadResp, err = c.service.UploadImage(file, fileName)
		}
		if err != nil {
			return ctx.JSON(ErrorResponse(err.Error()))
		}
		return ctx.JSON(SuccessResponse(uploadResp))
	}
}
