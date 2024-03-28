package config

import (
	"golang_app/golangApp/lib"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// this init method will called first when we import config package
// it will load all the env variables from.env file
// and set the env variable to global variables
// we read different env file for test and production
func init() {
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
}
