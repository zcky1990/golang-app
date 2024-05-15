package config

import (
	"context"
	"fmt"
	"log"
	"os"

	c "golang_app/golangApp/constants"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB struct holds the client and database references
type MongoDB struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// NewMongoDB initializes a new MongoDB instance
func NewMongoDB() (*MongoDB, error) {
	mongoHost := os.Getenv(c.MONGO_HOST)
	mongoPort := os.Getenv(c.MONGO_PORT)
	mongoType := os.Getenv(c.MONGO_TYPE)
	mongoOption := os.Getenv(c.MONGO_OPTION)
	mongoUsername := os.Getenv(c.MONGO_USERNAME)
	mongoPassword := os.Getenv(c.MONGO_PASSWORD)
	databaseName := os.Getenv(c.MONGO_DATABASE_NAME)

	url := buildMongoURL(mongoHost, mongoPort, mongoType, mongoOption, mongoUsername, mongoPassword)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %w", err)
	}

	db := client.Database(databaseName)

	mdb := &MongoDB{
		Client: client,
		Db:     db,
	}

	if err := mdb.createUserIndex(); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return mdb, nil
}

// buildMongoURL constructs the MongoDB connection URL
func buildMongoURL(host, port, typ, option, username, password string) string {
	url := fmt.Sprintf("%s://", typ)
	if username != "" && password != "" {
		url += fmt.Sprintf("%s:%s@", username, password)
	}
	if port != "" {
		url += fmt.Sprintf("%s:%s", host, port)
	} else {
		url += host
	}
	if option != "" {
		url += "/" + option
	}
	return url
}

// Disconnect disconnects the MongoDB client
func (mdb *MongoDB) Disconnect() {
	if mdb.Client != nil {
		err := mdb.Client.Disconnect(context.Background())
		if err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		} else {
			log.Println("Disconnected from MongoDB")
		}
	}
}

// createUserIndex creates an index on the User collection
func (mdb *MongoDB) createUserIndex() error {
	indexOptions := options.Index().SetUnique(true)
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1, // Index the email field in ascending order
		},
		Options: indexOptions,
	}
	_, err := mdb.Db.Collection("User").Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return fmt.Errorf("error creating index: %w", err)
	}
	return nil
}
