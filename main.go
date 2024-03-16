package main

import (
	"fmt"
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

	r.HandleFunc("POST /login", controller.Login)

	r.HandleFunc("POST /upload-image", controller.UploadFile)

	http.ListenAndServe(":10000", r)
	fmt.Println("Run Server on : localhost:10000")
}
