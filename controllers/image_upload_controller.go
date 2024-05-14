package controllers

import (
	"context"
	"golang_app/golangApp/constant"
	"golang_app/golangApp/services"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gofiber/fiber/v2"
)

var ctx context.Context
var cld *cloudinary.Cloudinary

const FOLDER = "folder"

type ImageController struct {
	service *services.CloudinaryService
}

func NewCloudinaryController(cloudinaryService *services.CloudinaryService) *ImageController {
	return &ImageController{service: cloudinaryService}
}

func (ctrl *ImageController) UploadFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var uploadResp *services.UploadImageResponse
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		files := form.File["file"]

		file, err := files[0].Open()
		if err != nil {
			return c.JSON(ErrorResponse(ctrl.service.Locale.Localization(constant.FAILED_OPEN_FILE)))
		}
		defer file.Close()
		fileName := files[0].Filename
		folder := form.Value[FOLDER][0]

		if folder != "" {
			uploadResp, err = ctrl.service.UploadImageToFolder(file, fileName, folder)
		} else {
			uploadResp, err = ctrl.service.UploadImage(file, fileName)
		}
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		return c.JSON(SuccessResponse(uploadResp))
	}
}
