package main

import (
	"fmt"
	"log"
	"golang_app/golangApp/config"
	"golang_app/golangApp/controllers"
	"golang_app/golangApp/middlewares"
	
	"net/http"
)


func main() {
	err := config.ConnectMongoDB("production")
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }
    defer config.DisconnectMongoDB()

	r := http.NewServeMux()

	indexHandler := http.HandlerFunc(controller.Index)
	r.Handle("GET /", middlewares.UserAuthenticate(indexHandler))

	r.HandleFunc("POST /login", controller.Login)

	http.ListenAndServe(":10000", r)
	fmt.Println("Run Server on : localhost:10000")
}
