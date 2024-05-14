package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
