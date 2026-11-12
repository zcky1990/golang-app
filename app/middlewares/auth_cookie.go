package middlewares

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

// Cookie names. Authentication lives in an httpOnly cookie so the SPA can never
// read (or leak) the JWT; CSRF is prevented with the double-submit pattern.
const (
	// AuthCookieName holds the JWT.
	AuthCookieName = "auth_token"

	// CSRFCookieName is readable by JavaScript on purpose: app/javascript/lib/api.js
	// echoes it back in the CSRFHeaderName header.
	CSRFCookieName = "csrf_token"

	// CSRFHeaderName is the header every non-GET request must carry.
	CSRFHeaderName = "X-CSRF-Token"
)

func secureCookie() bool {
	return os.Getenv("ENV") == "production"
}

func randomToken() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

// SetAuthCookies stores the JWT and a fresh CSRF token after a successful login.
func SetAuthCookies(ctx fiber.Ctx, auth *Authorization) error {
	if auth == nil || auth.Token == "" {
		return errors.New("empty authorization")
	}

	csrf, err := randomToken()
	if err != nil {
		return err
	}

	expires := time.Unix(auth.ExpireAt, 0)
	if auth.ExpireAt == 0 {
		expires = time.Now().Add(72 * time.Hour)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     AuthCookieName,
		Value:    auth.Token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   secureCookie(),
		SameSite: "Lax",
		Expires:  expires,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     CSRFCookieName,
		Value:    csrf,
		Path:     "/",
		HTTPOnly: false,
		Secure:   secureCookie(),
		SameSite: "Lax",
		Expires:  expires,
	})

	return nil
}

// ClearAuthCookies removes both cookies (Rails: `reset_session`).
func ClearAuthCookies(ctx fiber.Ctx) {
	past := time.Unix(0, 0)
	for _, name := range []string{AuthCookieName, CSRFCookieName} {
		ctx.Cookie(&fiber.Cookie{
			Name:     name,
			Value:    "",
			Path:     "/",
			HTTPOnly: name == AuthCookieName,
			Secure:   secureCookie(),
			SameSite: "Lax",
			Expires:  past,
		})
	}
}

// CSRFToken returns the current double-submit token.
func CSRFToken(ctx fiber.Ctx) string {
	return ctx.Cookies(CSRFCookieName)
}

// CSRF verifies the double-submit token for cookie-authenticated requests.
//
// Requests with no auth/CSRF cookie at all are let through: those are either
// pre-authentication calls (login, sign-up) or stateless bearer-token clients,
// which a third-party site cannot forge anyway.
func CSRF() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		switch ctx.Method() {
		case "GET", "HEAD", "OPTIONS":
			return ctx.Next()
		}

		csrfCookie := ctx.Cookies(CSRFCookieName)
		if csrfCookie == "" && ctx.Cookies(AuthCookieName) == "" {
			return ctx.Next()
		}

		header := ctx.Get(CSRFHeaderName)
		if header == "" || subtle.ConstantTimeCompare([]byte(header), []byte(csrfCookie)) != 1 {
			return ctx.Status(fiber.StatusForbidden).JSON(generateErrorMessage("Invalid CSRF token"))
		}

		return ctx.Next()
	}
}
