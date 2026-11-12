package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	c "golang_app/golangApp/constants"
	"golang_app/golangApp/lib"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

const bearerSchema = "Bearer "

type Authorization struct {
	Token    string `json:"auth_token,omitempty"`
	AuthType string `json:"auth_type,omitempty"`
	ExpireAt int64  `json:"expire_at,omitempty"`
}

// secretKey returns the JWT signing key.
//
// It is read lazily on purpose: init() runs before main() loads the .env file,
// so reading SECRET there (the previous behaviour) always produced an empty
// key and tokens were signed with "".
func secretKey() ([]byte, error) {
	secret := os.Getenv(c.SECRET)
	if secret == "" {
		return nil, errors.New("SECRET is not configured")
	}
	return []byte(secret), nil
}

// keyFunc also pins the signing method, so a token can never downgrade the
// HMAC algorithm (alg confusion / "none" attacks).
func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return secretKey()
}

func GenerateToken(email string, password string) (*Authorization, error) {
	key, err := secretKey()
	if err != nil {
		return &Authorization{}, err
	}

	expireAt := getExpireDate()
	claims := jwt.MapClaims{
		"user_email": email,
		"password":   password,
		"expire":     expireAt,
		// exp makes jwt.Parse reject expired tokens for us (jwt/v5 validates it).
		"exp": expireAt,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return &Authorization{}, err
	}
	return &Authorization{
		Token:    tokenString,
		AuthType: "Bearer",
		ExpireAt: expireAt,
	}, nil
}

// ValidateJWTToken reports whether the request carries a valid token.
//
// It no longer writes the response itself: ctx.JSON returns nil on success, so
// the old implementation signalled "invalid token" by returning nil and the
// middleware happily continued (auth bypass). The caller now decides the status
// code and body.
func ValidateJWTToken(ctx fiber.Ctx) error {
	_, err := getToken(ctx)
	return err
}

func GetEmailFromToken(ctx fiber.Ctx) (string, error) {
	token, err := getToken(ctx)
	if err != nil {
		return "", err
	}
	claims, err := extractClaims(token)
	if err != nil {
		return "", err
	}

	email, ok := claims["user_email"].(string)
	if !ok || email == "" {
		return "", errors.New("Invalid session")
	}
	return email, nil
}

// private method here
//
// The token is read from the httpOnly cookie first (the SPA never sees it), with
// the Authorization header kept as a fallback so CLI tools, curl and the existing
// API tests keep working unchanged.
func getToken(ctx fiber.Ctx) (*jwt.Token, error) {
	tokenString := ""
	if cookie := ctx.Cookies(AuthCookieName); cookie != "" {
		tokenString = bearerSchema + cookie
	} else if header := ctx.Get("Authorization"); header != "" {
		tokenString = header
	} else {
		return nil, errors.New("Missing Auth Token")
	}

	if !strings.HasPrefix(tokenString, bearerSchema) {
		return nil, errors.New("Invalid/Malformed Auth Token")
	}

	tokenSlice := strings.TrimPrefix(tokenString, bearerSchema)
	token, err := jwt.Parse(tokenSlice, keyFunc)
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

func getExpireDate() int64 {
	return time.Now().Add(lib.EnvDuration(c.SESSION_EXPIRATION, 72*time.Hour)).Unix()
}
