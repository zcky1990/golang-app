package models

import (
	"context"
	"errors"
	"fmt"
	"golang_app/golangApp/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const WEDDING_COLLECTION = "wedding_data"

type SocialMedia struct {
	Username string `json:"username"`
	Link     string `json:"link"`
	Platform string `json:"platform"`
}

type Parent struct {
	Father struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"father"`
	Mother struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"mother"`
}

type AkadResepsi struct {
	Date    string `json:"date"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Address string `json:"address"`
	MapURL  string `json:"mapUrl"`
}

type LoveStory struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}

type Envelop struct {
	BackgroundImage string `json:"backgroundImage"`
	BackgroundColor string `json:"backgroundColor"`
}

type Guest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Streaming struct {
	Platform   string `json:"platform"`
	StreamLink string `json:"streamLink"`
}

type Gift struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	BankName string `json:"bankname"`
	Link     string `json:"link"`
}

type Wedding struct {
	Date    string      `json:"date"`
	Akad    AkadResepsi `json:"akad"`
	Resepsi AkadResepsi `json:"resepsi"`
}

type BrideGroom struct {
	Firstname   string        `json:"firstname"`
	Lastname    string        `json:"lastname"`
	ProfileURL  string        `json:"profileUrl"`
	SocialMedia []SocialMedia `json:"socialMedia"`
	Parent      Parent        `json:"parent"`
}

type WeddingData struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty"`
	Bride     BrideGroom         `json:"bride"`
	Groom     BrideGroom         `json:"groom"`
	Wedding   Wedding            `json:"wedding"`
	LoveStory []LoveStory        `json:"loveStory"`
	Envelop   Envelop            `json:"envelop"`
	Gallery   []string           `json:"gallery"`
	Guest     Guest              `json:"guest"`
	Streaming []Streaming        `json:"streaming"`
	Gifts     []Gift             `json:"gifts"`
}

func ConvertWeddingDataToBSON(data WeddingData) (bson.M, error) {
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

	return bsonMap, nil
}

func CreateWeddingData(data WeddingData) (string, error) {
	result, err := config.GetDB().Collection(WEDDING_COLLECTION).InsertOne(context.Background(), data)
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

func UpdateWeddingDataById(id string, updates WeddingData) (string, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	data, _ := ConvertWeddingDataToBSON(updates)
	result, err := config.GetDB().Collection(WEDDING_COLLECTION).UpdateOne(
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

func GetWeddingDataById(id string) *WeddingData {
	result := WeddingData{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := config.GetDB().Collection(WEDDING_COLLECTION).FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func GetWeddingDataByUserId(userId string) *WeddingData {
	result := WeddingData{}
	objID, _ := primitive.ObjectIDFromHex(userId)
	err := config.GetDB().Collection(WEDDING_COLLECTION).FindOne(context.TODO(), bson.M{"user_id": objID}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}

func DeleteWeddingDataById(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := config.GetDB().Collection(WEDDING_COLLECTION).DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
