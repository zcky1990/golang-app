package services

import (
	"context"
	"errors"
	"fmt"
	"golang_app/golangApp/config"
	c "golang_app/golangApp/constant"
	"golang_app/golangApp/models"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const WEDDING_COLLECTION = "wedding_data"

type WeddingService struct {
	collection  *mongo.Collection
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewWeddingService(mongodb *config.MongoDB, locale *localize.Localization, redis *redis.RedisClient) *WeddingService {
	collection := mongodb.GetDB().Collection(WEDDING_COLLECTION)
	return &WeddingService{collection: collection, translation: locale, redis: redis}
}

func (c *WeddingService) convertWeddingDataToBSON(data models.WeddingData) (bson.M, error) {
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

func (service *WeddingService) CreateWeddingData(data models.WeddingData) (string, error) {
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

func (service *WeddingService) UpdateWeddingDataById(id string, updates models.WeddingData) (string, error) {
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

func (service *WeddingService) GetWeddingDataById(id string) (*models.WeddingData, error) {
	result := models.WeddingData{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := service.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (service *WeddingService) GetWeddingDataByUserId(userId string) (*models.WeddingData, error) {
	result := models.WeddingData{}
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
