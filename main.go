package main

import (
	cfg "golang_app/golangApp/config"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/mongo"
	"golang_app/golangApp/config/redis"
	"log"
)

var env string

func init() {
	cfg.LoadEvirontment()
	env = cfg.GetEnv()
}

func main() {
	mongoDB, err := mongo.NewMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect()

	translation := localize.NewLocalization()
	redisClient := redis.NewRedisClient()
	defer redisClient.Close()

	app := cfg.RoutesNew(mongoDB.Db, translation, redisClient)
	app.SetUpRoutes()
	app.StartServer()
}
