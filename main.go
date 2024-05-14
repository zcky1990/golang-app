package main

import (
	"golang_app/golangApp/config"
	controllers "golang_app/golangApp/controllers"
	"golang_app/golangApp/middlewares"
	"golang_app/golangApp/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

var env string

func init() {
	env = config.GetEnv()
	err := config.ConnectMongoDB(env)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	config.InitializeRedis(config.GetEnv())
}

func DisconnectDB() {
	config.DisconnectMongoDB()
}

func main() {
	defer DisconnectDB()
	app := fiber.New()
	api := app.Group("/api")

	v1 := api.Group("/v1")

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	cloudinaryService := services.NewUCloudinaryService()
	imageController := controllers.NewCloudinaryController(cloudinaryService)

	v1.Post("/users/sign-up", userController.Signup())
	v1.Post("/users/login", userController.Login())
	v1.Post("/users/update-user", middlewares.JWTMiddleware(), userController.UpdateUser())

	v1.Post("/upload/image", middlewares.JWTMiddleware(), imageController.UploadFile())
	app.Listen(":10000")
}
