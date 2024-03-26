package main

import (
	"golang_app/golangApp/config"
	controller "golang_app/golangApp/controllers"
	"golang_app/golangApp/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.ConnectMongoDB("production")
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer config.DisconnectMongoDB()

	config.InitializeCloudinary("production")
	config.InitializeRedis("production")

	app := fiber.New()
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Post("/users/sign-up", controller.Signup())
	v1.Post("/users/login", controller.Login())

	v1.Post("/upload/image", middlewares.JWTMiddleware(), controller.UploadFile())
	app.Listen(":10000")
}
