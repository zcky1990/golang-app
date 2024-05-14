package models

import (
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
