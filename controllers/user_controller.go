package controllers

import (
	"fmt"
	"golang_app/golangApp/constant"
	"golang_app/golangApp/middlewares"
	"golang_app/golangApp/models"
	"golang_app/golangApp/services"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service     *services.UserService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewUserController(userService *services.UserService, localize *localize.Localization, redis *redis.RedisClient) *UserController {
	return &UserController{service: userService, translation: localize, redis: redis}
}

func (c *UserController) Signup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.User
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ErrorResponse(err.Error()))
		}
		user := c.service.GetUserByEmail(params.Email)
		if user != nil {
			return ctx.JSON(ErrorResponse(c.translation.Localization(constant.EMAIL_TAKEN)))
		}
		data, err := c.service.CreateUser(params)
		if err != nil {
			return ctx.JSON(ErrorResponse(fmt.Sprintf("%s : %s", constant.MESSAGE_ERROR_FAILED_CREATE_USER, err.Error())))
		}
		return ctx.JSON(SuccessResponse(data))
	}
}

type Authorization struct {
	Token    string `json:"auth_token,omitempty"`
	AuthType string `json:"auth_type,omitempty"`
}

type LoginResponse struct {
	Authorization Authorization `json:"authorization,omitempty"`
	Users         *models.User  `json:"users,omitempty"`
}

func (c *UserController) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.User
		var token string
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ErrorResponse(err.Error()))
		}
		responseUser, err := c.service.GetUserByEmailAndPassword(params.Email, params.Password)
		if err == nil {
			if responseUser != nil {
				token, err = middlewares.GenerateToken(params.Email, params.Password)
				if err != nil {
					return ctx.JSON(ErrorResponse(err.Error()))
				}
			} else {
				return ctx.JSON(ErrorResponse(c.translation.Localization(constant.USER_NOT_FOUND)))
			}
		} else {
			return ctx.JSON(ErrorResponse(err.Error()))
		}
		response := LoginResponse{
			Authorization: Authorization{
				Token:    token,
				AuthType: "Bearer",
			},
			Users: responseUser,
		}
		return ctx.JSON(SuccessResponse(response))
	}
}

func (c *UserController) UpdateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.User
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ErrorResponse(err.Error()))
		}
		user := c.service.GetUserByEmail(params.Email)
		if user != nil {
			response, err := c.service.UpdateUserById(user.Id.Hex(), params)
			if err != nil {
				return ctx.JSON(ErrorResponse(err.Error()))
			}
			return ctx.JSON(SuccessResponse(response))
		} else {
			return ctx.JSON(ErrorResponse(c.translation.Localization(constant.USER_NOT_FOUND)))
		}
	}
}
