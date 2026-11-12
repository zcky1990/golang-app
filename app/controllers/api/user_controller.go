package api

import (
	mdl "golang_app/golangApp/app/middlewares"
	m "golang_app/golangApp/app/models"
	"golang_app/golangApp/app/services"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	"golang_app/golangApp/config/session"
	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var _ BaseApiController = (*UserController)(nil)

type UserController struct {
	service     *services.UserService
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewUserController(database *gorm.DB, localize *localize.Localization, redis *redis.RedisClient) *UserController {
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
func (ctrl *UserController) getLanguange(ctx fiber.Ctx) string {
	locale := ctx.Get("Accept-Language")
	if locale == "" {
		return ""
	}
	if locale == c.LOCALE_ENGLISH {
		return c.LOCALE_ENGLISH
	}
	return c.LOCALE_INDONESIA
}

func (ctrl *UserController) Signup() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var params m.User
		locale := ctrl.getLanguange(ctx)
		// Fiber v3: BodyParser() became Bind().Body()
		if err := ctx.Bind().Body(&params); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(ctrl.ErrorResponse(err.Error()))
		}
		user := ctrl.service.GetUserByEmail(params.Email)
		if user != nil {
			return ctx.Status(fiber.StatusConflict).JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("EMAIL_TAKEN", locale)))
		}
		data, err := ctrl.service.CreateUser(params, locale)
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("FAILED_CREATE_USER", locale)))
		}
		return ctx.Status(fiber.StatusCreated).JSON(ctrl.SuccessResponse(data))
	}
}

type LoginResponse struct {
	Authorization mdl.Authorization `json:"authorization,omitempty"`
	Users         *m.User           `json:"users,omitempty"`
}

// Login authenticates the user and, in addition to the JSON body, stores the
// token in an httpOnly cookie so the SPA never has to hold it in JavaScript.
func (ctrl *UserController) Login() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var params m.User
		var authorization *mdl.Authorization
		locale := ctrl.getLanguange(ctx)

		// Retrieve session
		sesStore, ok := ctx.Locals("session").(*session.SessionStore)
		if !ok || sesStore == nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(ctrl.ErrorResponse("Session store is not available"))
		}
		// Set a session value
		sec, err := sesStore.Store.Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(ctrl.ErrorResponse(err.Error()))
		}

		if err := ctx.Bind().Body(&params); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(ctrl.ErrorResponse(err.Error()))
		}

		responseUser, err := ctrl.service.GetUserByEmailAndPassword(params.Email, params.Password)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(ctrl.ErrorResponse(err.Error()))
		}
		if responseUser == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("USER_NOT_FOUND", locale)))
		}

		authorization, err = mdl.GenerateToken(params.Email, params.Password)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(ctrl.ErrorResponse(err.Error()))
		}

		// create session in redis
		sec.Set(responseUser.Email, responseUser.Email)
		if err := sec.Save(); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(ctrl.ErrorResponse(err.Error()))
		}

		// Cookie auth for the SPA (httpOnly JWT + readable CSRF token).
		if err := mdl.SetAuthCookies(ctx, authorization); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(ctrl.ErrorResponse(err.Error()))
		}

		response := LoginResponse{
			Authorization: *authorization,
			Users:         responseUser,
		}
		return ctx.JSON(ctrl.SuccessResponse(response))
	}
}

// Logout clears the auth cookies (Rails: `reset_session` + redirect).
func (ctrl *UserController) Logout() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		mdl.ClearAuthCookies(ctx)
		return ctx.JSON(ctrl.SuccessResponse(nil))
	}
}

func (ctrl *UserController) UpdateUser() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var params m.User
		locale := ctrl.getLanguange(ctx)
		if err := ctx.Bind().Body(&params); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(ctrl.ErrorResponse(err.Error()))
		}
		user := ctrl.service.GetUserByEmail(params.Email)
		if user == nil {
			return ctx.Status(fiber.StatusNotFound).JSON(ctrl.ErrorResponse(ctrl.translation.GetMessage("USER_NOT_FOUND", locale)))
		}
		response, err := ctrl.service.UpdateUserById(user.ID, params, locale)
		if err != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(ctrl.ErrorResponse(err.Error()))
		}
		return ctx.JSON(ctrl.SuccessResponse(response))
	}
}
