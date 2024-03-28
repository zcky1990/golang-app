package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initializing the db object as a package-level variable
var db *mongo.Database

// initializing the client object as a package-level variable
var client *mongo.Client

func GetDB() *mongo.Database {
	return db
}

func GetClient() *mongo.Client {
	return client
}

func ConnectMongoDB(env string) error {
	var mongoHost string
	var mongoPort string
	var mongoType string
	var mongoOption string
	var mongoUsername string
	var mongoPassword string
	var databaseName string

	mongoHost = os.Getenv("MONGO_HOST")
	mongoPort = os.Getenv("MONGO_PORT")
	mongoType = os.Getenv("MONGO_TYPE")
	mongoOption = os.Getenv("MONGO_OPTION")
	mongoUsername = os.Getenv("MONGO_USERNAME")
	mongoPassword = os.Getenv("MONGO_PASSWORD")
	databaseName = os.Getenv("MONGO_DATABASE_NAME")

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

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
		return err
	}

	// log.Println("MongoDB connected successfully!")

	db = client.Database(databaseName)
	// log.Println("Connected to MongoDB!")
	createUserIndex()
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

func createUserIndex() error {
	indexOptions := options.Index().SetUnique(true)
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1, // Index the email field in ascending order
		},
		Options: indexOptions,
	}
	_, err := db.Collection("User").Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("Error creating index: %v\n", err)
		return err
	}
	return nil
}
