package session

import (
	"os"
	"strconv"
	"time"

	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

type sessionStore struct {
	Session *session.Store
}

var RedisStore *redis.Storage

func SessionStoreNew() *sessionStore {
	initRedisStore()
	expiration, _ := time.ParseDuration(os.Getenv("SESSION_EXPIRATION"))
	store := session.New(session.Config{
		Expiration:     expiration,
		KeyLookup:      "cookie:session_id",
		CookieSameSite: "Lax",
		Storage:        RedisStore,
	})
	return &sessionStore{
		Session: store,
	}
}

func initRedisStore() {
	portnum, _ := strconv.Atoi(os.Getenv(c.REDIS_PORT))
	redisDB, _ := strconv.Atoi(os.Getenv(c.REDIS_DB))
	RedisStore = redis.New(redis.Config{
		Host:     os.Getenv(c.REDIS_HOST),
		Port:     portnum,
		Username: os.Getenv(c.REDIS_USERNAME),
		Password: os.Getenv(c.REDIS_PASSWORD),
		Database: redisDB,
	})
}
