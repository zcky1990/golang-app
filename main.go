package main

import (
	"golang_app/golangApp/config"
	controllers "golang_app/golangApp/controllers"
	"golang_app/golangApp/middlewares"
	"golang_app/golangApp/services"
	"golang_app/golangApp/utils/localize"
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

	translation := localize.NewLocalization()

	userService := services.NewUserService(translation)
	userController := controllers.NewUserController(userService, translation)

	cloudinaryService := services.NewUCloudinaryService(translation)
	imageController := controllers.NewCloudinaryController(cloudinaryService, translation)

	v1.Post("/users/sign-up", userController.Signup())
	v1.Post("/users/login", userController.Login())
	v1.Post("/users/update-user", middlewares.JWTMiddleware(), userController.UpdateUser())

	v1.Post("/upload/image", middlewares.JWTMiddleware(), imageController.UploadFile())
	app.Listen(":10000")
}
