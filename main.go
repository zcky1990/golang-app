package main

import (
	"golang_app/golangApp/config"
	controllers "golang_app/golangApp/controllers"
	"golang_app/golangApp/middlewares"
	"golang_app/golangApp/services"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"
	"log"

	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()
	api := app.Group("/api")

	v1 := api.Group("/v1")

	translation := localize.NewLocalization()
	redisClient := redis.NewRedisClient()

	userService := services.NewUserService(mongoDB, translation, redisClient)
	userController := controllers.NewUserController(userService, translation, redisClient)

	cloudinaryService := services.NewUCloudinaryService(translation, redisClient)
	imageController := controllers.NewCloudinaryController(cloudinaryService, translation, redisClient)

	v1.Post("/users/sign-up", userController.Signup())
	v1.Post("/users/login", userController.Login())
	v1.Post("/users/update-user", middlewares.JWTMiddleware(), userController.UpdateUser())

	v1.Post("/upload/image", middlewares.JWTMiddleware(), imageController.UploadFile())
	app.Listen(":10000")
}
