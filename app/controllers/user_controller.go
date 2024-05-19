package controllers

import (
	"fmt"
	mdl "golang_app/golangApp/app/middlewares"
	m "golang_app/golangApp/app/models"
	"golang_app/golangApp/app/services"
	c "golang_app/golangApp/constants"
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

func (ctrl *UserController) SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		c.STATUS: c.SUCCESS,
		c.DATA:   data,
	}
}

func (ctrl *UserController) ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}

func (ctrl *UserController) Signup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params m.User
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		user := ctrl.service.GetUserByEmail(params.Email)
		if user != nil {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.Localization(c.EMAIL_TAKEN)))
		}
		data, err := ctrl.service.CreateUser(params)
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(fmt.Sprintf("%s : %s", c.MESSAGE_ERROR_FAILED_CREATE_USER, err.Error())))
		}
		return ctx.JSON(ctrl.SuccessResponse(data))
	}
}

type LoginResponse struct {
	Authorization mdl.Authorization `json:"authorization,omitempty"`
	Users         *m.User           `json:"users,omitempty"`
}

func (ctrl *UserController) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params m.User
		var authorization *mdl.Authorization

		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		responseUser, err := ctrl.service.GetUserByEmailAndPassword(params.Email, params.Password)
		if err == nil {
			if responseUser != nil {
				authorization, err = mdl.GenerateToken(params.Email, params.Password)
				if err != nil {
					return ctx.JSON(ctrl.ErrorResponse(err.Error()))
				}
			} else {
				return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.Localization(c.USER_NOT_FOUND)))
			}
		} else {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		response := LoginResponse{
			Authorization: *authorization,
			Users:         responseUser,
		}
		return ctx.JSON(ctrl.SuccessResponse(response))
	}
}

func (ctrl *UserController) UpdateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params m.User
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		user := ctrl.service.GetUserByEmail(params.Email)
		if user != nil {
			response, err := ctrl.service.UpdateUserById(user.Id.Hex(), params)
			if err != nil {
				return ctx.JSON(ctrl.ErrorResponse(err.Error()))
			}
			return ctx.JSON(ctrl.SuccessResponse(response))
		} else {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.Localization(c.USER_NOT_FOUND)))
		}
	}
}
