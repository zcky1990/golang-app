package middlewares

import (
	"golang_app/golangApp/config/session"
	c "golang_app/golangApp/constants"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte
var expiration int64

type Authorization struct {
	Token    string `json:"auth_token,omitempty"`
	AuthType string `json:"auth_type,omitempty"`
	ExpireAt int64  `json:"expire_at,omitempty"`
}

func init() {
	jwtSecret = []byte(os.Getenv("SECRET"))
	expiration, _ = strconv.ParseInt(os.Getenv("SESSION_EXPIRATION"), 10, 64)
}

func JWTMiddleware() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		//check bearer and validate token
		const BEARER_SCHEMA = "Bearer "
		tokenString := ctx.Get("Authorization")
		var token *jwt.Token
		var err error
		if tokenString == "" {
			return ctx.JSON(generateErrorMessage("Missing Auth Token"))
		}
		if !strings.Contains(tokenString, BEARER_SCHEMA) {
			return ctx.JSON(generateErrorMessage("Invalid/Malformed Auth Token"))
		} else {
			tokenSlice := strings.Replace(tokenString, BEARER_SCHEMA, "", -1)
			token, err = jwt.Parse(tokenSlice, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})
			if err != nil {
				return ctx.JSON(generateErrorMessage(err.Error()))
			}
			if !token.Valid {
				return ctx.JSON(generateErrorMessage("Token Invalid"))
			}
		}

		//to make it more save, we also check if stil have session login in redis
		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.JSON(generateErrorMessage("Failed to extract claims from token"))
		}

		// Access specific claims
		name, nameOk := claims["user_email"].(string)
		if !nameOk {
			return ctx.JSON(generateErrorMessage("Invalid claims in token"))
		}

		// Retrieve session
		sesStore := ctx.Locals("session").(*session.SessionStore)
		// GET a session value
		sec, _ := sesStore.Store.Get(ctx)
		stringSession := sec.Get(name)

		// if session not present we make jwt token invalid
		if stringSession == nil {
			return ctx.JSON(generateErrorMessage("Invalid Session, Please login again"))
		}
		return ctx.Next()
	}
}

func GenerateToken(email string, password string) (*Authorization, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expire_at := getExpireDate()
	claims["user_email"] = email
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
	if expiration <= 0 {
		return time.Now().Add(time.Hour * 72).Unix()
	}
	return expiration
}
