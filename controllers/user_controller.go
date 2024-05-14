package controllers

import (
	"fmt"
	"golang_app/golangApp/constant"
	"golang_app/golangApp/middlewares"
	"golang_app/golangApp/models"
	"golang_app/golangApp/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{service: userService}
}

func (controller *UserController) Signup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.User
		if err := c.BodyParser(&params); err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		user := controller.service.GetUserByEmail(params.Email)
		if user != nil {
			return c.JSON(ErrorResponse(controller.service.Locale.Localization(constant.EMAIL_TAKEN)))
		}
		data, err := controller.service.CreateUser(params)
		if err != nil {
			return c.JSON(ErrorResponse(fmt.Sprintf("%s : %s", constant.MESSAGE_ERROR_FAILED_CREATE_USER, err.Error())))
		}
		return c.JSON(SuccessResponse(data))
	}
}

type LoginResponse struct {
	Authorization string       `json:"auth_token,omitempty"`
	Users         *models.User `json:"users,omitempty"`
}

func (controller *UserController) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.User
		var token string
		if err := c.BodyParser(&params); err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		responseUser, err := controller.service.GetUserByEmailAndPassword(params.Email, params.Password)
		if err == nil {
			if responseUser != nil {
				token, err = middlewares.GenerateToken(params.Email, params.Password)
				if err != nil {
					return c.JSON(ErrorResponse(err.Error()))
				}
			} else {
				return c.JSON(ErrorResponse(controller.service.Locale.Localization(constant.USER_NOT_FOUND)))
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

func (controller *UserController) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.User
		if err := c.BodyParser(&params); err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		user := controller.service.GetUserByEmail(params.Email)
		if user != nil {
			response, err := controller.service.UpdateUserById(user.Id.Hex(), params)
			if err != nil {
				return c.JSON(ErrorResponse(err.Error()))
			}
			return c.JSON(SuccessResponse(response))
		} else {
			return c.JSON(ErrorResponse(controller.service.Locale.Localization(constant.USER_NOT_FOUND)))
		}
	}
}
