package middlewares

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(os.Getenv("SECRET"))
}

func JWTMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		const BEARER_SCHEMA = "Bearer "
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.JSON(fiber.Map{
				"status":  "Error",
				"message": "Missing Auth Token",
			})
		}
		if !strings.Contains(tokenString, BEARER_SCHEMA) {
			return c.JSON(fiber.Map{
				"status":  "Error",
				"message": "Invalid/Malformed Auth Token",
			})
		} else {
			tokenSlice := strings.Replace(tokenString, BEARER_SCHEMA, "", -1)
			token, err := jwt.Parse(tokenSlice, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})
			if err != nil {
				return c.JSON(fiber.Map{
					"status":  "Error",
					"message": err.Error(),
				})
			}
			if !token.Valid {
				return c.JSON(fiber.Map{
					"status":  "Error",
					"message": "Token Invalid",
				})
			}

		}

		return c.Next()
	}
}

func GenerateToken(username string, password string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["password"] = password
	claims["expire"] = getExpireDate()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	stringtoken := fmt.Sprintf("Bearer %s", tokenString)
	return stringtoken, nil
}

func getExpireDate() int64 {
	return time.Now().Add(time.Hour * 72).Unix()
}
