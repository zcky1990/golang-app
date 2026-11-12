// Command migrate is the Go equivalent of `bin/rails db:migrate`.
//
// Usage:
//
//	go run ./db/migrate/cmd -direction=up      # apply pending migrations
//	go run ./db/migrate/cmd -direction=down    # roll back the last migration
//	go run ./db/migrate/cmd -direction=status  # list applied/pending versions
package main

import (
	"context"
	"flag"
	"log"

	cfg "golang_app/golangApp/config"
	"golang_app/golangApp/config/postgres"
	"golang_app/golangApp/db/migrate"
)

func main() {
	direction := flag.String("direction", "up", "migration direction: up, down, or status")
	flag.Parse()

	// Same bootstrap as main.go so .env is loaded before connecting.
	cfg.GetApplicationInstance()

	database, err := postgres.NewPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer database.Close()

	runner := migrate.NewRunner(database.DB)
	ctx := context.Background()

	switch *direction {
	case "up":
		err = runner.Up(ctx)
	case "down":
		err = runner.Down(ctx)
	case "status":
		err = runner.Status(ctx)
	default:
		log.Fatalf("unknown direction %q, use up, down, or status", *direction)
	}
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("[migrate] done")
}
