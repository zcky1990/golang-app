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
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	service     *services.UserService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewUserController(database *mongo.Database, localize *localize.Localization, redis *redis.RedisClient) *UserController {
	service := services.NewUserService(database, localize, redis)
	return &UserController{service: service, translation: localize, redis: redis}
}

func (c *UserController) SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		constant.STATUS: constant.SUCCESS,
		constant.DATA:   data,
	}
}

func (c *UserController) ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		constant.STATUS:        constant.FAILED,
		constant.ERROR_MESSAGE: message,
	}
}

func (c *UserController) Signup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.User
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(c.ErrorResponse(err.Error()))
		}
		user := c.service.GetUserByEmail(params.Email)
		if user != nil {
			return ctx.JSON(c.ErrorResponse(c.translation.Localization(constant.EMAIL_TAKEN)))
		}
		data, err := c.service.CreateUser(params)
		if err != nil {
			return ctx.JSON(c.ErrorResponse(fmt.Sprintf("%s : %s", constant.MESSAGE_ERROR_FAILED_CREATE_USER, err.Error())))
		}
		return ctx.JSON(c.SuccessResponse(data))
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
			return ctx.JSON(c.ErrorResponse(err.Error()))
		}
		responseUser, err := c.service.GetUserByEmailAndPassword(params.Email, params.Password)
		if err == nil {
			if responseUser != nil {
				token, err = middlewares.GenerateToken(params.Email, params.Password)
				if err != nil {
					return ctx.JSON(c.ErrorResponse(err.Error()))
				}
			} else {
				return ctx.JSON(c.ErrorResponse(c.translation.Localization(constant.USER_NOT_FOUND)))
			}
		} else {
			return ctx.JSON(c.ErrorResponse(err.Error()))
		}
		response := LoginResponse{
			Authorization: Authorization{
				Token:    token,
				AuthType: "Bearer",
			},
			Users: responseUser,
		}
		return ctx.JSON(c.SuccessResponse(response))
	}
}

func (c *UserController) UpdateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.User
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(c.ErrorResponse(err.Error()))
		}
		user := c.service.GetUserByEmail(params.Email)
		if user != nil {
			response, err := c.service.UpdateUserById(user.Id.Hex(), params)
			if err != nil {
				return ctx.JSON(c.ErrorResponse(err.Error()))
			}
			return ctx.JSON(c.SuccessResponse(response))
		} else {
			return ctx.JSON(c.ErrorResponse(c.translation.Localization(constant.USER_NOT_FOUND)))
		}
	}
}
