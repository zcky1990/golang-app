package middlewares

import (
	"fmt"
	"os"
	"strings"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang_app/golangApp/lib"
)

var secret string

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error get current directory", err)
		return
	}
	rootDir := lib.FindRootDir(currentDir)
	err = godotenv.Load(rootDir + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	secret = os.Getenv("SECRET")
	log.Println(secret)
}

func UserAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hi there, I love !")

        tokenHeader := r.Header.Get("Authorization")

        if tokenHeader == "" {
            w.WriteHeader(http.StatusForbidden)
            w.Write([]byte("Missing auth token"))
            return
        }

        tokenSlice := strings.Split(tokenHeader, " ")
        if len(tokenSlice) != 2 {
            w.WriteHeader(http.StatusForbidden)
            w.Write([]byte("Invalid/Malformed auth token"))
            return
        }

        tokenString := tokenSlice[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

        if err != nil {
            w.WriteHeader(http.StatusForbidden)
            w.Write([]byte(err.Error()))
            return
        }

        if !token.Valid {
            w.WriteHeader(http.StatusForbidden)
            w.Write([]byte("Invalid token"))
            return
        }

        next.ServeHTTP(w, r)
    })
}

func GenerateToken(username string, role string) string {
	claims := createClaim(username, role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Bearer %s", t)
}

//private method here
func createClaim(username string, role string) jwt.MapClaims {
	claims := jwt.MapClaims{
		"name": username,
		"role": role,
		"exp":  getExpireDate(),
	}
	return claims
}

func getExpireDate() int64 {
	return time.Now().Add(time.Hour * 72).Unix()
}