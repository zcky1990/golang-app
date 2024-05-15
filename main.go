package main

import (
	"golang_app/golangApp/config"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"
	"log"
)

var env string

func init() {
	env = config.GetEnv()
}

func main() {
	mongoDB, err := config.NewMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect()

	translation := localize.NewLocalization()
	redisClient := redis.NewRedisClient()
	defer redisClient.Close()

	app := config.RoutesNew(mongoDB.Db, translation, redisClient)
	app.SetUpRoutes()
	app.StartServer()

}
