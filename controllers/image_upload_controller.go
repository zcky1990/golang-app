package controllers

import (
	"context"
	"golang_app/golangApp/constant"
	"golang_app/golangApp/services"
	"golang_app/golangApp/utils/localize"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gofiber/fiber/v2"
)

var ctx context.Context
var cld *cloudinary.Cloudinary

type ImageController struct {
	service     *services.CloudinaryService
	Translation *localize.Localization
}

func NewCloudinaryController(cloudinaryService *services.CloudinaryService, locale *localize.Localization) *ImageController {
	return &ImageController{service: cloudinaryService, Translation: locale}
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
			return ctx.JSON(ErrorResponse(c.Translation.Localization(constant.FAILED_OPEN_FILE)))
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
