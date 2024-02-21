package config

import (
	"fmt"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//initializing the db object as a package-level variable
var Db *mongo.Database
//initializing the client object as a package-level variable
var client *mongo.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func ConnectMongoDB() ( error) {
    mongoHost := os.Getenv("MONGO_HOST")
    mongoPort := os.Getenv("MONGO_PORT")
    mongoType := os.Getenv("MONGO_TYPE")
    mongoOption := os.Getenv("MONGO_OPTION")
    mongoUsername := os.Getenv("MONGO_USERNAME")
    mongoPassword := os.Getenv("MONGO_PASSWORD")

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

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("Error creating MongoDB client:", err)
        return err
    }

    log.Println("MongoDB connected successfully!")

    Db = client.Database(os.Getenv("MONGO_DATABASE_NAME"))

    log.Println("Connected to MongoDB!")
	return nil
}

func DisconnectMongoDB() {
    if client != nil {
        err := client.Disconnect(context.Background())
        if err != nil {
            log.Printf("Error disconnecting from MongoDB: %v", err)
        } else {
            log.Println("Disconnected from MongoDB")
        }
    }
}