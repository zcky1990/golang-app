package models

import (
	"context"
	"log"

	"golang_app/golangApp/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

//use omitempty to automatically add id if we add empty id
type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Authtoken string             `json:"auth_token"`
}

//we add all query in models
func AddUser(user User) (*mongo.InsertOneResult, error) {
	result, err := config.Db.Collection("User").InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil, err
	}
	return result, nil
}

func GetUserByEmail(email string) User {
	result := User{}
	err := config.Db.Collection("User").FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func GetUserByEmailAndPassword(email string, password string) (User, error) {
	result := User{}
	err := config.Db.Collection("User").FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result, err
	}
	return result, nil
}

func GetAllUserList() []User {
	results := []User{}
	cursor, err := config.Db.Collection("User").Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return results
	}

	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		results = append(results, user)
	}
	return results
}

func SearchUser(search_type string, query string) []User {
	results := []User{}
	filter := bson.M{}
	if search_type == "name" {
		filter = bson.M{"$or": []interface{}{
			bson.M{"firstName": bson.M{"$regex": query, "$options": "im"}},
			bson.M{"lastName": bson.M{"$regex": query, "$options": "im"}},
		},
		}
	}

	if search_type == "email" {
		filter = bson.M{"email": query}
	}

	cursor, err := config.Db.Collection("User").Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return results
	}

	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		results = append(results, user)
	}
	return results
}