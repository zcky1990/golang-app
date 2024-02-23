package models

import (
	"context"
	"log"
	"fmt"

	"golang_app/golangApp/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func AddUser(user User) (string, error) {
    result, err := config.GetDB().Collection("User").InsertOne(context.Background(), user)
    if err != nil {
        log.Printf("Error while inserting user: %v\n", err)
        return "", err
    }

    insertedID, ok := result.InsertedID.(primitive.ObjectID)
    if !ok {
        return "", fmt.Errorf("failed to extract inserted ID")
    }

    insertedIDString := insertedID.Hex()
    return insertedIDString, nil
}

func GetUserByEmail(email string) *User {
	result := User{}
	err := config.GetDB().Collection("User").FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil
	}
	return &result
}

func DeleteUserById(id string){
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error convert string Id to primitive Object ID: %v\n", err)
	}
	result, err := config.GetDB().Collection("User").DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Printf("Error while Deleting single document, Reason: %v\n",err)
	}
	log.Printf("Number of Deleted document : %v\n",result.DeletedCount)
}

func GetUserByEmailAndPassword(email string, password string) (User, error) {
	result := User{}
	err := config.GetDB().Collection("User").FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result, err
	}
	return result, nil
}

func GetAllUserList(page, pageSize int) ([]User, error) {
    var results []User
    offset := (page - 1) * pageSize

    options := options.Find().
        SetSkip(int64(offset)).
        SetLimit(int64(pageSize))

    cursor, err := config.GetDB().Collection("User").Find(context.TODO(), bson.M{}, options)
    if err != nil {
        log.Printf("Error while getting all users: %v\n", err)
        return nil, err
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

	cursor, err := config.GetDB().Collection("User").Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Error while searching for users: %v\n", err)
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