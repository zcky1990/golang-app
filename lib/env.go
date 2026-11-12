package lib

import (
	"os"
	"strconv"
	"time"
)

// EnvDuration reads a duration from the environment and accepts both notations
// used across this project (and by Rails' session/secret configuration):
//
//	SESSION_EXPIRATION=72h     -> Go duration string
//	SESSION_EXPIRATION=259200  -> number of seconds
//
// The fallback is returned when the variable is missing or unparsable, so a
// misconfigured .env never results in a zero-length JWT or session.
func EnvDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	if duration, err := time.ParseDuration(value); err == nil {
		return duration
	}
	if seconds, err := strconv.ParseInt(value, 10, 64); err == nil {
		return time.Duration(seconds) * time.Second
	}
	return fallback
}
