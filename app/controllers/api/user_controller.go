package api

import (
	mdl "golang_app/golangApp/app/middlewares"
	m "golang_app/golangApp/app/models"
	"golang_app/golangApp/app/services"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	"golang_app/golangApp/config/session"
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
func (ctrl *UserController) getLanguange(ctx *fiber.Ctx) string {
	locale := ctx.Get("Accept-Language")
	if locale == "" {
		return c.LOCALE_ENGLISH
	}
	if locale == c.LOCALE_ENGLISH {
		return c.LOCALE_ENGLISH
	} else {
		return c.LOCALE_INDONESIA
	}
}

func (ctrl *UserController) Signup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params m.User
		locale := ctrl.getLanguange(ctx)
		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		user := ctrl.service.GetUserByEmail(params.Email)
		if user != nil {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("EMAIL_TAKEN", locale)))
		}
		data, err := ctrl.service.CreateUser(params)
		if err != nil {
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("FAILED_CREATE_USER", locale)))
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
		locale := ctrl.getLanguange(ctx)
		// Retrieve session
		sesStore := ctx.Locals("session").(*session.SessionStore)
		// Set a session value
		sec, _ := sesStore.Store.Get(ctx)

		if err := ctx.BodyParser(&params); err != nil {
			return ctx.JSON(ctrl.ErrorResponse(err.Error()))
		}
		responseUser, err := ctrl.service.GetUserByEmailAndPassword(params.Email, params.Password)
		if err == nil {
			if responseUser != nil {
				authorization, err = mdl.GenerateToken(params.Email, params.Password)
				//create session in redis
				sec.Set(responseUser.Email, responseUser.Email)
				sec.Save()

				if err != nil {
					return ctx.JSON(ctrl.ErrorResponse(err.Error()))
				}
			} else {
				return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("USER_NOT_FOUND", locale)))
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
		locale := ctrl.getLanguange(ctx)
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
			return ctx.JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("USER_NOT_FOUND", locale)))
		}
	}
}
