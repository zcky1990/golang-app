package controller

import (
	"golang_app/golangApp/middlewares"
	"golang_app/golangApp/models"

	"github.com/gofiber/fiber/v2"
)

func Signup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.User
		if err := c.BodyParser(&params); err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		user := models.GetUserByEmail(params.Email)
		if user != nil {
			return c.JSON(ErrorResponse("Email Has been taken"))
		}
		data, err := models.CreateUser(params)
		if err != nil {
			return c.JSON(ErrorResponse("Failed to create User :" + err.Error()))
		}
		return c.JSON(SuccessResponse(data))
	}
}

type LoginResponse struct {
	Authorization string       `json:"auth_token,omitempty"`
	Users         *models.User `json:"users,omitempty"`
}

func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.User
		var token string
		if err := c.BodyParser(&params); err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		responseUser, err := models.GetUserByEmailAndPassword(params.Email, params.Password)
		if err == nil {
			if responseUser != nil {
				token, err = middlewares.GenerateToken(params.Email, params.Password)
				if err != nil {
					return c.JSON(ErrorResponse(err.Error()))
				}
			} else {
				return c.JSON(ErrorResponse("User not found"))
			}
		} else {
			return c.JSON(ErrorResponse(err.Error()))
		}
		response := LoginResponse{
			Authorization: token,
			Users:         responseUser,
		}
		return c.JSON(SuccessResponse(response))
	}
}

func UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.User
		if err := c.BodyParser(&params); err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		user := models.GetUserByEmail(params.Email)
		if user != nil {
			response, err := models.UpdateUserById(user.Id.Hex(), params)
			if err != nil {
				return c.JSON(ErrorResponse(err.Error()))
			}
			return c.JSON(SuccessResponse(response))
		} else {
			return c.JSON(ErrorResponse("User not found"))
		}
	}
}
