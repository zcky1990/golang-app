package config

import (
	"fmt"
	"context"
	"log"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//initializing the db object as a package-level variable
var db *mongo.Database
//initializing the client object as a package-level variable
var client *mongo.Client

func findRootDir(dir string) string {
	knownProjectItems := []string{"go.mod", "main.go"}
	for {
		for _, item := range knownProjectItems {
			if _, err := os.Stat(filepath.Join(dir, item)); err == nil {
				return dir
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rootDir := findRootDir(currentDir)
	fmt.Println(rootDir)
	err = godotenv.Load(rootDir + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func GetDB() *mongo.Database{
	return db
}

func GetClient() *mongo.Client {
	return client
}

func ConnectMongoDB(env string) ( error) {
	var mongoHost string
	var mongoPort string
	var mongoType string
	var mongoOption string
	var mongoUsername string
	var mongoPassword string

	if env != "test" {
		mongoHost = os.Getenv("MONGO_HOST")
		mongoPort = os.Getenv("MONGO_PORT")
		mongoType = os.Getenv("MONGO_TYPE")
		mongoOption = os.Getenv("MONGO_OPTION")
		mongoUsername = os.Getenv("MONGO_USERNAME")
		mongoPassword = os.Getenv("MONGO_PASSWORD")
	}else {
		mongoHost = os.Getenv("MONGO_HOST_TEST")
		mongoPort = os.Getenv("MONGO_PORT_TEST")
		mongoType = os.Getenv("MONGO_TYPE_TEST")
		mongoOption = os.Getenv("MONGO_OPTION_TEST")
		mongoUsername = os.Getenv("MONGO_USERNAME_TEST")
		mongoPassword = os.Getenv("MONGO_PASSWORD_TEST")
	}

    var url string
    if mongoPort != "" {
        url = mongoHost + ":" + mongoPort
    } else {
        url = mongoHost
    }

    if mongoUsername != "" && mongoPassword != "" {
        url = mongoUsername + ":" + mongoPassword + "@" + url
    }

    if mongoOption != "" {
        url = url + "/" + mongoOption
    }

    mongoURL := mongoType + "://" + url
    fmt.Println(mongoURL)

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    clientOptions := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)
	
	var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("Error creating MongoDB client:", err)
        return err
    }

    log.Println("MongoDB connected successfully!")

    db = client.Database(os.Getenv("MONGO_DATABASE_NAME"))
    log.Println("Connected to MongoDB!")
	return nil
}

func DisconnectMongoDB() {
    if client != nil {
        err := client.Disconnect(context.Background())
        if err != nil {
            log.Printf("Error disconnecting from MongoDB: %v", err)
        } else {
			client = nil
            log.Println("Disconnected from MongoDB")
        }
    }
}