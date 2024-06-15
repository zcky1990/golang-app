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

func (s *WeddingService) CreateWeddingData(data m.WeddingData, locale string) (string, error) {
	result, err := s.collection.InsertOne(context.Background(), data)
	if err != nil {
		return "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New(s.translation.GetMessage("FAILED_CREATE_WEDDDING", locale))
	}

	insertedIDString := insertedID.Hex()
	return insertedIDString, nil
}

func (s *WeddingService) UpdateWeddingDataById(id string, updates m.WeddingData, locale string) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	data, _ := s.ConvertToBSON(updates)
	result, err := s.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": data},
	)
	if err != nil {
		return "", errors.New(s.translation.GetMessage("FAILED_UPDATE_WEDDDING", locale))
	}
	if result.ModifiedCount == 0 {
		return "", errors.New(s.translation.GetMessage("NO_DATA_WEDDDING_UPDATED", locale))
	}
	return s.translation.GetMessage("SUCCESS_UPDATE_WEDDING", locale), nil
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
