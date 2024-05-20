package api

import (
	"fmt"
	mdl "golang_app/golangApp/app/middlewares"
	m "golang_app/golangApp/app/models"
	"golang_app/golangApp/app/services"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	c "golang_app/golangApp/constants"

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

func (cntrl *UserController) SuccessResponse(data interface{}) fiber.Map {
	return fiber.Map{
		c.STATUS: c.SUCCESS,
		c.DATA:   data,
	}
}

func (cntrl *UserController) ErrorResponse(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}

func (cntrl *UserController) Signup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params m.User
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(cntrl.ErrorResponse(err.Error()))
		}
		user := cntrl.service.GetUserByEmail(params.Email)
		if user != nil {
			return ctx.JSON(cntrl.ErrorResponse(cntrl.translation.Localization(c.EMAIL_TAKEN)))
		}
		data, err := cntrl.service.CreateUser(params)
		if err != nil {
			return ctx.JSON(cntrl.ErrorResponse(fmt.Sprintf("%s : %s", c.MESSAGE_ERROR_FAILED_CREATE_USER, err.Error())))
		}
		return ctx.JSON(cntrl.SuccessResponse(data))
	}
}

type LoginResponse struct {
	Authorization mdl.Authorization `json:"authorization,omitempty"`
	Users         *m.User           `json:"users,omitempty"`
}

func (cntrl *UserController) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params m.User
		var authorization *mdl.Authorization

		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(cntrl.ErrorResponse(err.Error()))
		}
		responseUser, err := cntrl.service.GetUserByEmailAndPassword(params.Email, params.Password)
		if err == nil {
			if responseUser != nil {
				authorization, err = mdl.GenerateToken(params.Email, params.Password)
				if err != nil {
					return ctx.JSON(cntrl.ErrorResponse(err.Error()))
				}
			} else {
				return ctx.JSON(cntrl.ErrorResponse(cntrl.translation.Localization(c.USER_NOT_FOUND)))
			}
		} else {
			return ctx.JSON(cntrl.ErrorResponse(err.Error()))
		}
		response := LoginResponse{
			Authorization: *authorization,
			Users:         responseUser,
		}
		return ctx.JSON(cntrl.SuccessResponse(response))
	}
}

func (cntrl *UserController) UpdateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params m.User
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(cntrl.ErrorResponse(err.Error()))
		}
		user := cntrl.service.GetUserByEmail(params.Email)
		if user != nil {
			response, err := cntrl.service.UpdateUserById(user.Id.Hex(), params)
			if err != nil {
				return ctx.JSON(cntrl.ErrorResponse(err.Error()))
			}
			return ctx.JSON(cntrl.SuccessResponse(response))
		} else {
			return ctx.JSON(cntrl.ErrorResponse(cntrl.translation.Localization(c.USER_NOT_FOUND)))
		}
	}
}
