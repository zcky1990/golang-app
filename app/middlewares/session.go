package middlewares

import (
	"errors"

	"golang_app/golangApp/config/session"

	"github.com/gofiber/fiber/v3"
)

func CheckSession(ctx fiber.Ctx) error {
	// Retrieve session
	sesStore, ok := ctx.Locals("session").(*session.SessionStore)
	if !ok || sesStore == nil {
		return errors.New("Invalid Session, Please login again")
	}

	// GET a session value
	sec, err := sesStore.Store.Get(ctx)
	if err != nil {
		return errors.New("Invalid Session, Please login again")
	}

	// get current session ID from cookies and send to redis
	result, err := sesStore.RedisStore.Get(sec.ID())
	if err != nil {
		// Fiber v3: c.Redirect() returns a builder, so the old
		// c.Redirect("/login", fiber.StatusTemporaryRedirect) no longer compiles.
		// return ctx.Redirect().Status(fiber.StatusTemporaryRedirect).To("/login")
		return errors.New("Invalid Session, Please login again")
	}
	if len(result) == 0 {
		// return ctx.Redirect().Status(fiber.StatusTemporaryRedirect).To("/login")
		return errors.New("Invalid Session, Please login again")
	}
	return nil
}

// because we are using api, we are gonna check session value base on email
func CheckSessionClaim(ctx fiber.Ctx, email string) error {
	sesStore, ok := ctx.Locals("session").(*session.SessionStore)
	if !ok || sesStore == nil {
		return errors.New("Invalid Session, Please login again")
	}

	sec, err := sesStore.Store.Get(ctx)
	if err != nil {
		return errors.New("Invalid Session, Please login again")
	}

	stringSession := sec.Get(email)
	if stringSession == nil {
		return errors.New("Invalid Session, Please login again")
	}
	return nil
}
