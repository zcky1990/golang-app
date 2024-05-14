package services

import (
	"context"
	"golang_app/golangApp/constant"
	"golang_app/golangApp/utils/localize"
	"golang_app/golangApp/utils/redis"
	"mime/multipart"
	"os"

	s "strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadImageResponse struct {
	SecureUrl string `json:"secure_url"`
	PublicId  string `json:"public_id"`
}

type CloudinaryService struct {
	cld         *cloudinary.Cloudinary
	ctx         context.Context
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewUCloudinaryService(locale *localize.Localization, redis *redis.RedisClient) *CloudinaryService {
	var name string
	var api string
	var secret string

	name = os.Getenv(constant.CLOUD_NAME)
	api = os.Getenv(constant.CLOUD_API_KEY)
	secret = os.Getenv(constant.CLOUD_API_SECRET)

	cloudConfig, _ := cloudinary.NewFromParams(name, api, secret)
	contex := context.Background()

	return &CloudinaryService{cld: cloudConfig, ctx: contex, translation: locale, redis: redis}
}

// upload image to cloudinary to spesific folder
func (service *CloudinaryService) UploadImageToFolder(file multipart.File, filename string, folder string) (*UploadImageResponse, error) {
	resp, err := service.cld.Upload.Upload(service.ctx, file, uploader.UploadParams{PublicID: s.Join([]string{folder, filename}, "/")})
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
	resp, err := service.cld.Upload.Upload(service.ctx, file, uploader.UploadParams{PublicID: filename})
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
