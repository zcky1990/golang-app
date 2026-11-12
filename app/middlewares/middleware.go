package middlewares

import (
	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v3"
)

// CheckJWTandSessionMiddleware is the Fiber v3 replacement of the old
// `func(ctx *fiber.Ctx) error` middleware.
//
// Every failure is answered once, here, with a proper 401 status; the helpers in
// jwt.go/session.go only return errors now, so an invalid token can never fall
// through to ctx.Next().
func CheckJWTandSessionMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		if err := ValidateJWTToken(ctx); err != nil {
			return unauthorizedResponse(ctx, err)
		}

		email, err := GetEmailFromToken(ctx)
		if err != nil {
			return unauthorizedResponse(ctx, err)
		}

		if err := CheckSessionClaim(ctx, email); err != nil {
			return unauthorizedResponse(ctx, err)
		}

		return ctx.Next()
	}
}

func unauthorizedResponse(ctx fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(generateErrorMessage(err.Error()))
}

func generateErrorMessage(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}
