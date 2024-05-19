package services

import (
	"context"
	"errors"
	"fmt"
	m "golang_app/golangApp/app/models"
	c "golang_app/golangApp/constants"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Ensure WeddingService implements the BaseService interface
var _ BaseService = (*WeddingService)(nil)

type WeddingService struct {
	collection  *mongo.Collection
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewWeddingService(mongodb *mongo.Database, locale *localize.Localization, redis *redis.RedisClient) *WeddingService {
	collection := mongodb.Collection(m.WEDDING_COLLECTION)
	return &WeddingService{collection: collection, translation: locale, redis: redis}
}

func (s *WeddingService) ConvertToBSON(data interface{}) (bson.M, error) {
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

func (c *WeddingService) convertWeddingDataToBSON(data m.WeddingData) (bson.M, error) {
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
	// remove empty value from bson
	for key, value := range bsonMap {
		if value == "" || value == nil {
			delete(bsonMap, key)
		}
	}
	return bsonMap, nil
}

func (service *WeddingService) CreateWeddingData(data m.WeddingData) (string, error) {
	result, err := service.collection.InsertOne(context.Background(), data)
	if err != nil {
		return "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf(c.MESSAGE_ERROR_FAILED_EXTRACT_INSERTED_ID)
	}

	insertedIDString := insertedID.Hex()
	return insertedIDString, nil
}

func (service *WeddingService) UpdateWeddingDataById(id string, updates m.WeddingData) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	data, _ := service.convertWeddingDataToBSON(updates)
	result, err := service.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": data},
	)
	if err != nil {
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", errors.New(c.MESSAGE_ERROR_FAILED_UPDATE_DATA)
	}
	return c.MESSAGE_SUCCESS_UPDATE_WEDDING_DATA, nil
}

func (service *WeddingService) GetWeddingDataById(id string) (*m.WeddingData, error) {
	result := m.WeddingData{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := service.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (service *WeddingService) GetWeddingDataByUserId(userId string) (*m.WeddingData, error) {
	result := m.WeddingData{}
	objID, _ := primitive.ObjectIDFromHex(userId)
	err := service.collection.FindOne(context.TODO(), bson.M{"user_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (service *WeddingService) DeleteWeddingDataById(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := service.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
