package middlewares

import (
	"errors"
	"golang_app/golangApp/config/session"

	"github.com/gofiber/fiber/v2"
)

func CheckSession(ctx *fiber.Ctx) error {
	// Retrieve session
	sesStore := ctx.Locals("session").(*session.SessionStore)

	// GET a session value
	sec, _ := sesStore.Store.Get(ctx)

	//get current session ID from cookies and send to redis
	sec.ID()

	result, err := sesStore.RedisStore.Get(sec.ID())
	if err != nil {
		return errors.New("Invalid Session, Please login again")

		// Redirect to the login route
		// ctx.Redirect("/login", fiber.StatusTemporaryRedirect)
		// Stop the middleware chain
		// return nil
	}
	if result == nil {
		return errors.New("Invalid Session, Please login again")

		// Redirect to the login route
		// ctx.Redirect("/login", fiber.StatusTemporaryRedirect)
		// Stop the middleware chain
		// return nil
	}
	return nil
}

// becasue we are using api, we are gonna check session value base on email
func CheckSessionClaim(ctx *fiber.Ctx, email string) error {
	sesStore := ctx.Locals("session").(*session.SessionStore)
	sec, _ := sesStore.Store.Get(ctx)
	stringSession := sec.Get(email)

	if stringSession == nil {
		return errors.New("Invalid Session, Please login again")
	}
	return nil
}
