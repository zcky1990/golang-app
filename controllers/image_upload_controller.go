package controller

import (
	"context"
	"golang_app/golangApp/lib"
	"log"
	"net/http"
	"os"

	s "strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

var ctx context.Context
var cld *cloudinary.Cloudinary

type UploadImageResponse struct {
	SecureUrl string `json:"secure_url"`
	PublicId  string `json:"public_id"`
}

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error get current directory", err)
		return
	}
	rootDir := lib.FindRootDir(currentDir)
	err = godotenv.Load(rootDir + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cld, _ = cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_API_KEY"), os.Getenv("CLOUD_API_SECRET"))
	ctx = context.Background()
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	imageType := r.Form.Get("type")
	var response []byte

	if err != nil || imageType == "" {
		if imageType == "" {
			response = ErrorResponse("Key Type not found")
		} else {
			response = ErrorResponse(err.Error())
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	} else {
		defer file.Close()
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: s.Join([]string{imageType, handler.Filename}, "/")})
		if err != nil {
			log.Printf("Error while Upload File, Reason: %v\n", err)
			response = ErrorResponse(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		} else {
			responseUpload := UploadImageResponse{
				SecureUrl: resp.SecureURL,
				PublicId:  resp.PublicID,
			}
			response = SuccessResponse(responseUpload)
		}
	}
	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
