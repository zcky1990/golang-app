package middlewares

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"golang_app/golangApp/app/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func MockHandler(c *fiber.Ctx) error {
	// Return a mock response
	return c.SendString("This is a mock response")
}

func TestGenerateToken(t *testing.T) {
	token, err := middlewares.GenerateToken("testuser", "testpassword")
	// Check if the token is generated successfully
	assert.NoError(t, err)

	// Check if the generated token is in the expected format
	expectedPrefix := "Bearer "
	assert.Contains(t, token, expectedPrefix)
}

func TestJWTMiddleware(t *testing.T) {
	//see https://github.com/gofiber/fiber/blob/main/ctx_test.go to create mock context
	var body map[string]string
	app := fiber.New()
	//create mock endpoint to test JWT middleware
	app.Post("/", middlewares.JWTMiddleware(), MockHandler)

	// Set up a mock request
	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	// Set up a mock response
	//check if the response message is "Missing Auth Token
	resp, _ := app.Test(req)
	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, &body)
	assert.Equal(t, "Missing Auth Token", body["message"])

	//test if header not contain Bearer prefix in the request
	req.Header.Set("Authorization", "invalid_token")
	resp, _ = app.Test(req)
	data, _ = io.ReadAll(resp.Body)
	json.Unmarshal(data, &body)
	assert.Equal(t, "Invalid/Malformed Auth Token", body["message"])

	//test if header with invalid in the request when parse
	req.Header.Set("Authorization", "Bearer invalid_token")
	resp, _ = app.Test(req)
	data, _ = io.ReadAll(resp.Body)
	json.Unmarshal(data, &body)
	assert.Equal(t, "token is malformed: token contains an invalid number of segments", body["message"])

	//test if header with invalid token
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MTE5MDQ1OTcsInBhc3N3b3JkIjoiMTIzNDU2Njc3IiwidXNlcm5hbWUiOiJqb2huQGV4YW1wbGUuY29tIn0.qsApVZ4FNF4oQ_iVPywyBDcvfmEVFOuHB8wRW1n2IdM")
	resp, _ = app.Test(req)
	data, _ = io.ReadAll(resp.Body)
	json.Unmarshal(data, &body)
	assert.Equal(t, "token signature is invalid: signature is invalid", body["message"])

	//test using valid token
	//generate token to test and set to header
	token, _ := middlewares.GenerateToken("testuser", "testpassword")
	req.Header.Set("Authorization", token)
	resp, _ = app.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
}
