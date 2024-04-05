package config

import (
	"golang_app/golangApp/lib"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var env string

// this init method will called first when we import config package
// it will load all the env variables from.env file
// and set the env variable to global variables
// we read different env file for test and production
func init() {
	var filePath string
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error get current directory", err)
		return
	}
	rootDir := lib.FindRootDir(currentDir)
	if isRunningTests() {
		filePath = ".env.test"
	} else {
		filePath = ".env"
	}
	err = godotenv.Load(rootDir + "/" + filePath)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	env = os.Getenv("ENV")
}

func GetEnv() string {
	return env
}

func isRunningTests() bool {
	for _, arg := range os.Args {
		if strings.Contains(arg, "test") {
			return true
		}
	}
	return false
}
