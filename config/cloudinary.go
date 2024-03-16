package config

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	s "strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// initialize the context directly
var cloudinaryCtx = context.Background()
var cld *cloudinary.Cloudinary

type UploadImageResponse struct {
	SecureUrl string `json:"secure_url"`
	PublicId  string `json:"public_id"`
}

func InitializeCloudinary(env string) {
	var name string
	var api string
	var secret string

	if env != "test" {
		name = os.Getenv("CLOUD_NAME")
		api = os.Getenv("CLOUD_API_KEY")
		secret = os.Getenv("CLOUD_API_SECRET")
	} else {
		name = os.Getenv("CLOUD_NAME_TEST")
		api = os.Getenv("CLOUD_API_KEY_TEST")
		secret = os.Getenv("CLOUD_API_SECRET_TEST")
	}

	// log.Println("Initialize cloudinary")
	cld, _ = cloudinary.NewFromParams(name, api, secret)
}

// upload image to cloudinary to spesific folder
func UploadImageToFolder(file multipart.File, filename string, folder string) (*UploadImageResponse, error) {
	resp, err := cld.Upload.Upload(cloudinaryCtx, file, uploader.UploadParams{PublicID: s.Join([]string{folder, filename}, "/")})
	if err != nil {
		log.Printf("Error while Upload File, Reason: %v\n", err)
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
func UploadImage(file multipart.File, filename string) (*UploadImageResponse, error) {
	resp, err := cld.Upload.Upload(cloudinaryCtx, file, uploader.UploadParams{PublicID: filename})
	if err != nil {
		log.Printf("Error while Upload File, Reason: %v\n", err)
		return nil, err
	} else {
		responseUpload := &UploadImageResponse{
			SecureUrl: resp.SecureURL,
			PublicId:  resp.PublicID,
		}
		return responseUpload, nil
	}
}
