package middlewares

import (
	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v2"
)

func CheckJWTandSessionMiddleware() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		err := ValidateJWTToken(ctx)

		if err != nil {
			return ctx.JSON(generateErrorMessage(err.Error()))
		}

		email, err := GetEmailFromToken(ctx)
		if err != nil {
			return ctx.JSON(generateErrorMessage(err.Error()))
		}

		err = CheckSessionClaim(ctx, email)

		if err != nil {
			return ctx.JSON(generateErrorMessage(err.Error()))
		}

		return ctx.Next()
	}
}

func generateErrorMessage(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}
