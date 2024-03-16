package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"

	"golang_app/golangApp/config"
)

var ctx context.Context
var cld *cloudinary.Cloudinary

func UploadFile(w http.ResponseWriter, r *http.Request) {
	var uploadResp *config.UploadImageResponse
	var response []byte
	var err error
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	fileName := handler.Filename
	folder := r.Form.Get("folder")

	defer file.Close()

	if folder != "" {
		uploadResp, err = config.UploadImageToFolder(file, fileName, folder)
	} else {
		uploadResp, err = config.UploadImage(file, fileName)
	}

	if err != nil {
		log.Printf("Error while Upload File, Reason: %v\n", err)
		response = ErrorResponse(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	} else {
		response = SuccessResponse(uploadResp)
	}
	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
