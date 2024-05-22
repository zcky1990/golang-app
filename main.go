package main

import (
	cfg "golang_app/golangApp/config"
	e "golang_app/golangApp/config/environments"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/mongo"
	"golang_app/golangApp/config/redis"
	"log"
	"os"
)

func initializeEnvirotmentConfigurationAndEnvirotment(env string) e.EnvironmentConfiguration {
	var config e.EnvironmentConfiguration
	if env == "development" {
		config = e.NewDevEnvConfiguration(env)
	}
	if env == "test" {
		config = e.NewTestEnvConfiguration(env)
	}
	if env == "production" {
		config = e.NewProdEnvConfiguration(env)
	}
	return config
}

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	envConfig := initializeEnvirotmentConfigurationAndEnvirotment(env)
	envConfig.LoadEvirontmentFile()

	mongoDB, err := mongo.NewMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect()

	translation := localize.NewLocalization()
	redisClient := redis.NewRedisClient()
	defer redisClient.Close()

	app := cfg.RoutesNew(envConfig, mongoDB.Db, translation, redisClient)
	app.SetUpRoutes()
	app.StartServer()
}
