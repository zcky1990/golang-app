package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	s "strings"

	"github.com/redis/go-redis/v9"
)

var redisCtx = context.Background()
var redisClient *redis.Client

func InitializeRedis(env string) {
	var host string
	var port string
	var db int
	var username string
	var password string
	var address string
	var err error
	var redisUrl string

	if env != "test" {
		redisUrl = os.Getenv("REDIS_URL")
		username = os.Getenv("REDIS_USERNAME")
		password = os.Getenv("REDIS_PASSWORD")
		host = os.Getenv("REDIS_HOST")
		port = os.Getenv("REDIS_PORT")
		db, err = strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	} else {
		redisUrl = os.Getenv("REDIS_URL_TEST")
		username = os.Getenv("REDIS_USERNAME_TEST")
		password = os.Getenv("REDIS_PASSWORD_TEST")
		host = os.Getenv("REDIS_HOST_TEST")
		port = os.Getenv("REDIS_PORT_TEST")
		db, err = strconv.Atoi(os.Getenv("REDIS_DB_TEST"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	// log.Println("Initialize Redis")
	if host != "" && port != "" {
		address = s.Join([]string{host, port}, ":")
	} else {
		address = "localhost:6379"
	}

	if redisUrl == "" {
		redisClient = redis.NewClient(&redis.Options{
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
		redisClient = redis.NewClient(opt)
	}
}

func SetRedis(key string, value string, expiration time.Duration) error {
	err := redisClient.Set(redisCtx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis(key string) (string, error) {
	val, err := redisClient.Get(redisCtx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
