package middlewares

import (
	c "golang_app/golangApp/constants"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

type Authorization struct {
	Token    string `json:"auth_token,omitempty"`
	AuthType string `json:"auth_type,omitempty"`
	ExpireAt int64  `json:"expire_at,omitempty"`
}

func init() {
	jwtSecret = []byte(os.Getenv("SECRET"))
}

func JWTMiddleware() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		const BEARER_SCHEMA = "Bearer "
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.JSON(generateErrorMessage("Missing Auth Token"))
		}
		if !strings.Contains(tokenString, BEARER_SCHEMA) {
			return ctx.JSON(generateErrorMessage("Invalid/Malformed Auth Token"))
		} else {
			tokenSlice := strings.Replace(tokenString, BEARER_SCHEMA, "", -1)
			token, err := jwt.Parse(tokenSlice, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})
			if err != nil {
				return ctx.JSON(generateErrorMessage(err.Error()))
			}
			if !token.Valid {
				return ctx.JSON(generateErrorMessage("Token Invalid"))
			}
		}
		return ctx.Next()
	}
}

func GenerateToken(username string, password string) (*Authorization, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expire_at := getExpireDate()
	claims["username"] = username
	claims["password"] = password
	claims["expire"] = expire_at

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return &Authorization{}, err
	}
	return &Authorization{
		Token:    tokenString,
		AuthType: "Bearer",
		ExpireAt: expire_at,
	}, nil
}

func generateErrorMessage(message string) fiber.Map {
	return fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: message,
	}
}

func getExpireDate() int64 {
	return time.Now().Add(time.Hour * 72).Unix()
}
