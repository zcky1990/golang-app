package main

import (
	cfg "golang_app/golangApp/config"
	"golang_app/golangApp/config/initializers"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/postgres"
	"golang_app/golangApp/config/redis"
	"log"
)

func main() {
	// we initiate app environment, and load env from env file
	appCnf := cfg.GetApplicationInstance()

	// run boot-time initializers (Rails: config/initializers)
	if err := initializers.Run(); err != nil {
		log.Fatalf("Failed to run initializers: %v", err)
	}

	// initiate database (PostgreSQL via GORM)
	postgresDB, err := postgres.NewPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer postgresDB.Close()

	// initiate localization
	translation := localize.NewLocalization()

	// initiate redis
	redisClient := redis.NewRedisClient()
	defer redisClient.Close()

	app := cfg.RoutesNew(appCnf, postgresDB.DB, translation, redisClient)

	// Fiber v3's Listen returns an error (bind failures, unexpected stop), and
	// ShutdownWithTimeout returns one too - both end up here.
	if err := app.StartServer(); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}
}
