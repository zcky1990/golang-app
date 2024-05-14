package controllers

import (
	"encoding/json"
	"golang_app/golangApp/constant"
	"io"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		constant.STATUS: constant.SUCCESS,
		constant.DATA:   data,
	}
}

func ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		constant.STATUS:        constant.FAILED,
		constant.ERROR_MESSAGE: message,
	}
}

func SetParams(body io.Reader, v interface{}) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
