package main

import (
	"golang_app/golangApp/config"
	controller "golang_app/golangApp/controllers"
	"golang_app/golangApp/middlewares"
	"log"

	"net/http"
)

func main() {
	err := config.ConnectMongoDB("production")
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer config.DisconnectMongoDB()

	config.InitializeCloudinary("production")
	config.InitializeRedis("production")

	r := http.NewServeMux()

	indexHandler := http.HandlerFunc(controller.Index)
	r.Handle("GET /", middlewares.UserAuthenticate(indexHandler))

	r.HandleFunc("POST /api/sign-up", controller.Signup)
	r.HandleFunc("POST /api/login", controller.Login)

	uploadImageHandler := http.HandlerFunc(controller.UploadFile)
	r.Handle("POST /api/upload-image", middlewares.UserAuthenticate(uploadImageHandler))

	http.ListenAndServe(":10000", r)
	log.Println("Run Server on : localhost:10000")
}
