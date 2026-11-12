package session

import (
	"time"

	c "golang_app/golangApp/constants"
	"golang_app/golangApp/lib"

	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/redis/go-redis/v9"
)

// SessionStore bundles Fiber's store with the Redis adapter that backs it
// (Rails: config/initializers/session_store.rb).
type SessionStore struct {
	Store      *session.Store
	RedisStore *RedisStorage
}

// SessionStoreNew builds the cookie + Redis backed session store.
//
// The Redis client is injected instead of being created here so the whole
// application shares one connection, one set of REDIS_* variables and one
// driver version.
func SessionStoreNew(client *redis.Client) *SessionStore {
	storage := NewRedisStorage(client)
	store := session.New(session.Config{
		Expiration:     lib.EnvDuration(c.SESSION_EXPIRATION, 72*time.Hour),
		KeyLookup:      "cookie:session_id",
		CookieSameSite: "Lax",
		Storage:        storage,
	})
	return &SessionStore{
		Store:      store,
		RedisStore: storage,
	}
}
