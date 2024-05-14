package services

import (
	"context"
	"golang_app/golangApp/config"
	"golang_app/golangApp/constant"
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
	cld           *cloudinary.Cloudinary
	cloudinaryCtx context.Context
	Locale        *config.Localization
}

func NewUCloudinaryService() *CloudinaryService {
	var name string
	var api string
	var secret string

	name = os.Getenv(constant.CLOUD_NAME)
	api = os.Getenv(constant.CLOUD_API_KEY)
	secret = os.Getenv(constant.CLOUD_API_SECRET)

	cloudConfig, _ := cloudinary.NewFromParams(name, api, secret)
	contex := context.Background()

	locale := config.NewLocalization()

	return &CloudinaryService{cld: cloudConfig, cloudinaryCtx: contex, Locale: locale}
}

// upload image to cloudinary to spesific folder
func (service *CloudinaryService) UploadImageToFolder(file multipart.File, filename string, folder string) (*UploadImageResponse, error) {
	resp, err := service.cld.Upload.Upload(service.cloudinaryCtx, file, uploader.UploadParams{PublicID: s.Join([]string{folder, filename}, "/")})
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
	resp, err := service.cld.Upload.Upload(service.cloudinaryCtx, file, uploader.UploadParams{PublicID: filename})
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
