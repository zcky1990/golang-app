package middlewares

import (
	"golang_app/golangApp/config/session"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(ssn *session.SessionStore) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sess, err := ssn.Session.Get(ctx)
		if err != nil {
			return err
		}
		defer sess.Save()
		return ctx.Next()
	}
}
