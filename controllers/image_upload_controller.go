package controller

import (
	"context"
	"golang_app/golangApp/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gofiber/fiber/v2"
)

var ctx context.Context
var cld *cloudinary.Cloudinary

func UploadFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var uploadResp *config.UploadImageResponse
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		files := form.File["file"]

		file, err := files[0].Open()
		if err != nil {
			return c.JSON(ErrorResponse(Localization("FAILED_OPEN_FILE")))
		}
		defer file.Close()
		fileName := files[0].Filename
		folder := form.Value["folder"][0]

		if folder != "" {
			uploadResp, err = config.UploadImageToFolder(file, fileName, folder)
		} else {
			uploadResp, err = config.UploadImage(file, fileName)
		}
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		return c.JSON(SuccessResponse(uploadResp))
	}
}
