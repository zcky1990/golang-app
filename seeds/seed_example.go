// Seeder example, the Go equivalent of Rails' db/seeds.rb.
//
// Run it with:
//
//	go run ./seeds
package main

import (
	"log"

	m "golang_app/golangApp/app/models"
	cfg "golang_app/golangApp/config"
	"golang_app/golangApp/config/postgres"
)

func main() {
	// Load .env the same way main.go does before touching the database.
	cfg.GetApplicationInstance()

	database, err := postgres.NewPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer database.Close()

	// FirstOrCreate keeps the seed idempotent, so it is safe to run repeatedly.
	user := m.User{
		Username:  "demo",
		Email:     "demo@example.com",
		Firstname: "Demo",
		Lastname:  "User",
		Password:  "password",
	}
	if err := database.DB.Where("email = ?", user.Email).FirstOrCreate(&user).Error; err != nil {
		log.Fatalf("Failed to seed user: %v", err)
	}

	log.Println("seed running: demo user ready with id", user.ID)
}
