package services

import "go.mongodb.org/mongo-driver/bson"

type BaseService interface {
	ConvertToBSON(data interface{}) (bson.M, error)
}
