package models

import (
	"context"
	"fmt"
	"log"

	"golang_app/golangApp/config"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const USER_COLLECTION = "User"

// use omitempty to automatically add id if we add empty id
type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty"`
	Email     string             `json:"email,omitempty"`
	Firstname string             `json:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty"`
	Authtoken string             `json:"auth_token,omitempty"`
	Password  string             `json:"password,omitempty"`
}

func ConvertUserToBSON(data User) (bson.M, error) {
	// Marshal the struct to BSON
	bsonData, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}
	// Unmarshal the BSON to bson.M
	var bsonMap bson.M
	err = bson.Unmarshal(bsonData, &bsonMap)
	if err != nil {
		return nil, err
	}
	//remove empty value from bson
	for key, value := range bsonMap {
		if value == "" || value == nil {
			delete(bsonMap, key)
		}
	}
	return bsonMap, nil
}

func CreateUser(user User) (string, error) {
	result, err := config.GetDB().Collection(USER_COLLECTION).InsertOne(context.Background(), user)
	if err != nil {
		// log.Printf("Error while inserting user: %v\n", err)
		return "", err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to extract inserted ID")
	}
	insertedIDString := insertedID.Hex()
	return insertedIDString, nil
}

func UpdateUserById(id string, updates User) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	data, _ := ConvertUserToBSON(updates)
	result, err := config.GetDB().Collection(USER_COLLECTION).UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": data},
	)
	if err != nil {
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", errors.New("No User Updated")
	}
	return "Success Update User", nil
}

func GetUserByEmail(email string) *User {
	result := User{}
	err := config.GetDB().Collection(USER_COLLECTION).FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func DeleteUserById(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := config.GetDB().Collection(USER_COLLECTION).DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmailAndPassword(email string, password string) (*User, error) {
	var result User
	err := config.GetDB().Collection(USER_COLLECTION).FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func GetAllUserList(page, pageSize int) ([]User, error) {
	var results []User
	offset := (page - 1) * pageSize
	options := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(pageSize))
	cursor, err := config.GetDB().Collection(USER_COLLECTION).Find(context.TODO(), bson.M{}, options)
	if err != nil {
		// log.Printf("Error while getting all users: %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			// log.Printf("Error decoding user: %v\n", err)
			continue
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results, nil
}

func SearchUser(searchType string, query string) *[]User {
	results := []User{}
	filter := bson.M{}
	if searchType == "name" {
		filter = bson.M{"$or": []interface{}{
			bson.M{"firstname": bson.M{"$regex": query, "$options": "im"}},
			bson.M{"lastname": bson.M{"$regex": query, "$options": "im"}},
		}}
	}
	if searchType == "email" {
		filter = bson.M{"email": query}
	}
	cursor, err := config.GetDB().Collection(USER_COLLECTION).Find(context.TODO(), filter)
	if err != nil {
		// log.Printf("Error while searching for users: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Error decoding user: %v\n", err)
			continue
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil // No users found, return nil
	}
	return &results
}
