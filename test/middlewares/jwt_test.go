package middlewares

import (
	"testing"

	"golang_app/golangApp/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestGenerateToken(t *testing.T) {
	token, err := middlewares.GenerateToken("testuser", "testpassword")
	// Check if the token is generated successfully
	if err != nil {

		assert.NotNil(t, err, "expected error not to be nil")
	}

	// Check if the generated token is in the expected format
	expectedPrefix := "Bearer "
	if !startsWith(token, expectedPrefix) {
		t.Errorf("Expected token to start with '%s', got: %s", expectedPrefix, token)
	}
}

func TestJWTMiddleware(t *testing.T) {
	//see https://github.com/gofiber/fiber/blob/main/ctx_test.go to create mock context
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	// Set up a mock request with a valid JWT token in the Authorization header
	c.Request().Header.Set("Authorization", "Bearer valid_token")

	// Call the JWT middleware function
	err := middlewares.JWTMiddleware()(c)
	// Check if the middleware returns an error (it should not for a valid token)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
