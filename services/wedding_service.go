package services

import (
	"context"
	"errors"
	"fmt"
	"golang_app/golangApp/config"
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

func (c *WeddingService) ConvertWeddingDataToBSON(data models.WeddingData) (bson.M, error) {
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

func (c *WeddingService) CreateWeddingData(data models.WeddingData) (string, error) {
	result, err := c.collection.InsertOne(context.Background(), data)
	if err != nil {
		return "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to extract inserted ID")
	}

	insertedIDString := insertedID.Hex()
	return insertedIDString, nil
}

func (c *WeddingService) UpdateWeddingDataById(id string, updates models.WeddingData) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	data, _ := c.ConvertWeddingDataToBSON(updates)
	result, err := c.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": data},
	)
	if err != nil {
		return "", err
	}
	if result.ModifiedCount == 0 {
		return "", errors.New("no wedding data updated")
	}
	return "success update wedding data", nil
}

func (c *WeddingService) GetWeddingDataById(id string) *models.WeddingData {
	result := models.WeddingData{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := c.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func (c *WeddingService) GetWeddingDataByUserId(userId string) *models.WeddingData {
	result := models.WeddingData{}
	objID, _ := primitive.ObjectIDFromHex(userId)
	err := c.collection.FindOne(context.TODO(), bson.M{"user_id": objID}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func (c *WeddingService) DeleteWeddingDataById(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := c.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
