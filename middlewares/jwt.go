package middlewares

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

func init() {
	jwtSecret = os.Getenv("SECRET")
	log.Println("secret : " + jwtSecret)
}

func JWTMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.JSON(fiber.Map{
				"status":  "Error",
				"message": "Missing Auth Token",
			})
		}
		tokenSlice := strings.Split(tokenString, " ")
		if len(tokenSlice) != 2 {
			return c.JSON(fiber.Map{
				"status":  "Error",
				"message": "Invalid/Malformed Auth Token",
			})
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		log.Println("tokenString : ", tokenString)
		log.Println("token : ", token)
		log.Println("valid : ", token.Valid)

		if err != nil || !token.Valid {
			return c.JSON(fiber.Map{
				"status":  "Error",
				"message": fiber.ErrUnauthorized,
			})
		}
		return c.Next()
	}
}

func GenerateToken(username string, password string) (string, error) {
	log.Println("jwtSecret : " + jwtSecret)
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
