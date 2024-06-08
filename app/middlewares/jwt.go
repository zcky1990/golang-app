package middlewares

import (
	"errors"
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
		tokenString := ctx.Get("Authorization")
		token, err := validateToken(tokenString)
		if err != nil {
			return ctx.JSON(generateErrorMessage(err.Error()))
		}

		claims, err := extractClaims(token)
		if err != nil {
			return ctx.JSON(generateErrorMessage(err.Error()))
		}

		if err := checkSession(ctx, claims); err != nil {
			return ctx.JSON(generateErrorMessage(err.Error()))
		}

		return ctx.Next()
	}
}

func validateToken(tokenString string) (*jwt.Token, error) {
	const bearerSchema = "Bearer "
	if tokenString == "" {
		return nil, errors.New("Missing Auth Token")
	}
	if !strings.HasPrefix(tokenString, bearerSchema) {
		return nil, errors.New("Invalid/Malformed Auth Token")
	}

	tokenSlice := strings.TrimPrefix(tokenString, bearerSchema)
	token, err := jwt.Parse(tokenSlice, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Token Invalid")
	}
	return token, nil
}

func extractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Failed to extract claims from token")
	}

	if _, ok := claims["user_email"].(string); !ok {
		return nil, errors.New("Invalid claims in token")
	}

	return claims, nil
}

func checkSession(ctx *fiber.Ctx, claims jwt.MapClaims) error {
	email := claims["user_email"].(string)
	sesStore := ctx.Locals("session").(*session.SessionStore)
	sec, _ := sesStore.Store.Get(ctx)
	stringSession := sec.Get(email)

	if stringSession == nil {
		return errors.New("Invalid Session, Please login again")
	}
	return nil
}

func GenerateToken(email string, password string) (*Authorization, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expireAt := getExpireDate()
	claims["user_email"] = email
	claims["password"] = password
	claims["expire"] = expireAt

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return &Authorization{}, err
	}
	return &Authorization{
		Token:    tokenString,
		AuthType: "Bearer",
		ExpireAt: expireAt,
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
