package main

import (
	cfg "golang_app/golangApp/config"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	"log"

	"golang_app/golangApp/config/mongo"
)

func main() {
	// we initiate app environment, and load env from env file
	appCnf := cfg.GetApplicationInstance()

	// initiate database
	mongoDB, err := mongo.NewMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect()

	// initiate localization
	translation := localize.NewLocalization()

	// initiate redis
	redisClient := redis.NewRedisClient()
	defer redisClient.Close()

	app := cfg.RoutesNew(appCnf, mongoDB.Db, translation, redisClient)
	app.StartServer()
}
