package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	c "golang_app/golangApp/constant"

	s "strings"

	"github.com/redis/go-redis/v9"
)

// RedisClient is a struct that wraps the go-redis client
type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisClient initializes and returns a new RedisClient
func NewRedisClient() *RedisClient {
	var host string
	var port string
	var db int
	var username string
	var password string
	var address string
	var redisUrl string
	var rdb *redis.Client

	redisUrl = os.Getenv(c.REDIS_URL)
	username = os.Getenv(c.REDIS_USERNAME)
	password = os.Getenv(c.REDIS_PASSWORD)
	host = os.Getenv(c.REDIS_HOST)
	port = os.Getenv(c.REDIS_PORT)
	db, _ = strconv.Atoi(os.Getenv(c.REDIS_DB))

	if host != "" && port != "" {
		address = s.Join([]string{host, port}, ":")
	} else {
		address = "localhost:6379"
	}

	if redisUrl == "" {
		rdb = redis.NewClient(&redis.Options{
			Addr:     address,
			Username: username,
			Password: password,
			DB:       db,
		})
	} else {
		opt, err := redis.ParseURL(redisUrl)
		if err != nil {
			panic(err)
		}
		rdb = redis.NewClient(opt)
	}

	return &RedisClient{
		client: rdb,
		ctx:    context.Background(),
	}
}

// Set sets a key-value pair with an expiration time
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	err := r.client.Set(r.ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key: %w", err)
	}
	return nil
}

// Get retrieves the value for a given key
func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	} else if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}
	return val, nil
}

// Delete removes a key from Redis
func (r *RedisClient) Delete(key string) error {
	err := r.client.Del(r.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}
	return nil
}

// Close closes the Redis client connection
func (r *RedisClient) Close() error {
	err := r.client.Close()
	if err != nil {
		return fmt.Errorf("failed to close Redis client: %w", err)
	}
	return nil
}
