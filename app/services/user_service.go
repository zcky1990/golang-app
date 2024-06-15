package services

import (
	"context"
	"errors"
	m "golang_app/golangApp/app/models"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Ensure UserService implements the BaseService interface
var _ BaseService = (*UserService)(nil)

type UserService struct {
	collection  *mongo.Collection
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewUserService(mongodb *mongo.Database, locale *localize.Localization, redis *redis.RedisClient) *UserService {
	collection := mongodb.Collection(m.USER_COLLECTION)
	return &UserService{collection: collection, translation: locale, redis: redis}
}

func (s *UserService) ConvertToBSON(data interface{}) (bson.M, error) {
	bsonData, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}
	var bsonMap bson.M
	err = bson.Unmarshal(bsonData, &bsonMap)
	if err != nil {
		return nil, err
	}
	for key, value := range bsonMap {
		if value == "" || value == nil {
			delete(bsonMap, key)
		}
	}
	return bsonMap, nil
}

func (s *UserService) CreateUser(user m.User, locale string) (string, error) {
	result, err := s.collection.InsertOne(context.Background(), user)
	if err != nil {
		// log.Printf("Error while inserting user: %v\n", err)
		return "", err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New(s.translation.GetMessage("FAILED_CREATE_USER", locale))
	}
	insertedIDString := insertedID.Hex()
	return insertedIDString, nil
}

func (s *UserService) UpdateUserById(id string, updates m.User, locale string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	data, _ := s.ConvertToBSON(updates)
	result, err := s.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": data},
	)
	if err != nil {
		return "", errors.New(s.translation.GetMessage("FAILED_UPDATE_USER", locale))
	}
	if result.ModifiedCount == 0 {
		return "", errors.New(s.translation.GetMessage("NO_DATA_USER_UPDATED", locale))
	}
	return s.translation.GetMessage("SUCCESS_UPDATE_USER", locale), nil
}

func (s *UserService) GetUserByEmail(email string) *m.User {
	result := m.User{}
	err := s.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func (s *UserService) DeleteUserById(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := s.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByEmailAndPassword(email string, password string) (*m.User, error) {
	var result m.User
	err := s.collection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (s *UserService) GetAllUserList(page, pageSize int) ([]m.User, error) {
	var results []m.User
	offset := (page - 1) * pageSize
	options := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(pageSize))
	cursor, err := s.collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user m.User
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

func (s *UserService) SearchUser(searchType string, query string) *[]m.User {
	results := []m.User{}
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
	cursor, err := s.collection.Find(context.TODO(), filter)
	if err != nil {
		// log.Printf("Error while searching for users: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user m.User
		if err := cursor.Decode(&user); err != nil {
			// log.Printf("Error decoding user: %v\n", err)
			continue
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil // No users found, return nil
	}
	return &results
}
