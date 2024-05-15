package services

import (
	"context"
	c "golang_app/golangApp/constants"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"
	"mime/multipart"
	"os"

	s "strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.mongodb.org/mongo-driver/bson"
)

// Ensure CloudinaryService implements the BaseService interface
var _ BaseService = (*CloudinaryService)(nil)

type UploadImageResponse struct {
	SecureUrl string `json:"secure_url"`
	PublicId  string `json:"public_id"`
}

type CloudinaryService struct {
	cloudinary  *cloudinary.Cloudinary
	ctx         context.Context
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewCloudinaryService(locale *localize.Localization, redis *redis.RedisClient) *CloudinaryService {
	var name string
	var api string
	var secret string

	name = os.Getenv(c.CLOUD_NAME)
	api = os.Getenv(c.CLOUD_API_KEY)
	secret = os.Getenv(c.CLOUD_API_SECRET)

	cloudinary, _ := cloudinary.NewFromParams(name, api, secret)
	contex := context.Background()

	return &CloudinaryService{cloudinary: cloudinary, ctx: contex, translation: locale, redis: redis}
}

func (s *CloudinaryService) ConvertToBSON(data interface{}) (bson.M, error) {
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

// upload image to cloudinary to spesific directory
func (service *CloudinaryService) UploadImageToFolder(file multipart.File, filename string, directory string) (*UploadImageResponse, error) {
	resp, err := service.cloudinary.Upload.Upload(service.ctx, file, uploader.UploadParams{PublicID: s.Join([]string{directory, filename}, "/")})
	if err != nil {
		return nil, err
	} else {
		responseUpload := &UploadImageResponse{
			SecureUrl: resp.SecureURL,
			PublicId:  resp.PublicID,
		}
		return responseUpload, nil
	}
}

// upload image to cloudinary
func (service *CloudinaryService) UploadImage(file multipart.File, filename string) (*UploadImageResponse, error) {
	resp, err := service.cloudinary.Upload.Upload(service.ctx, file, uploader.UploadParams{PublicID: filename})
	if err != nil {
		return nil, err
	} else {
		responseUpload := &UploadImageResponse{
			SecureUrl: resp.SecureURL,
			PublicId:  resp.PublicID,
		}
		return responseUpload, nil
	}
}
