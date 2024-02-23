package main

import (
	"fmt"
	"log"
	"golang_app/golangApp/config"
	"golang_app/golangApp/controllers"
	
	"net/http"
)


func main() {
	err := config.ConnectMongoDB("production")
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }
    defer config.DisconnectMongoDB()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controller.Index)

	http.ListenAndServe(":10000", mux)
	fmt.Println("Run Server on : localhost:10000")
}
