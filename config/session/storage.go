package session

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisStorage adapts go-redis to fiber.Storage, the interface Fiber's session
// middleware persists sessions through.
//
// The gofiber/storage/redis v1 driver is built against Fiber v2 (and pulls in
// the v2 module), so instead of adding a second Redis driver we wrap the client
// the application already owns (config/redis). Only the five methods of
// fiber.Storage are needed.
type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisStorage wraps an existing go-redis client. A nil client is allowed:
// the adapter then behaves like an empty store instead of panicking, which
// keeps the app usable when Redis is not configured.
func NewRedisStorage(client *redis.Client) *RedisStorage {
	return &RedisStorage{client: client, ctx: context.Background()}
}

// Get returns the raw session bytes. Fiber expects (nil, nil) when the key does
// not exist - the Go equivalent of a Rails session that is simply gone.
func (s *RedisStorage) Get(key string) ([]byte, error) {
	if s == nil || s.client == nil {
		return nil, nil
	}
	value, err := s.client.Get(s.ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	return value, err
}

// Set stores the session bytes, with exp == 0 meaning "no expiry".
func (s *RedisStorage) Set(key string, val []byte, exp time.Duration) error {
	if s == nil || s.client == nil {
		return nil
	}
	return s.client.Set(s.ctx, key, val, exp).Err()
}

// Delete removes a single session.
func (s *RedisStorage) Delete(key string) error {
	if s == nil || s.client == nil {
		return nil
	}
	return s.client.Del(s.ctx, key).Err()
}

// Reset drops every session, the equivalent of `rake db:sessions:clear`.
func (s *RedisStorage) Reset() error {
	if s == nil || s.client == nil {
		return nil
	}
	return s.client.FlushDB(s.ctx).Err()
}

// Close is a no-op: the client is owned (and closed) by config/redis, exactly
// like a Rails connection pool is owned by ActiveRecord.
func (s *RedisStorage) Close() error {
	return nil
}
