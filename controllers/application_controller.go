package controller

import (
	"encoding/json"
	"io"

	"golang_app/golangApp/config"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		"status": "Success",
		"data":   data,
	}
}

func ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		"status": "Error",
		"error":  message,
	}
}

func SetParams(body io.Reader, v interface{}) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func Localization(message string) string {
	locale, _ := config.GetInstance().GetMessage(message)
	return locale
}
