package config

import (
	"golang_app/golangApp/lib"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// log.Println("Initialize Environtment Variable")
	var env string
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error get current directory", err)
		return
	}
	rootDir := lib.FindRootDir(currentDir)

	if os.Getenv("ENV") == "test" {
		env = ".env." + os.Getenv("ENV")
	} else {
		env = ".env"
	}

	err = godotenv.Load(rootDir + "/" + env)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// log.Println(os.Getenv("ENV"))
}
