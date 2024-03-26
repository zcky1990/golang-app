package controller

import (
	"fmt"
	"golang_app/golangApp/middlewares"
	"golang_app/golangApp/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love !")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var response []byte
	var err error
	var userParams *models.User

	w.Header().Set("Content-Type", "application/json")
	err = SetParams(r.Body, &userParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorResponse(err.Error()))
		return
	}

	user := models.GetUserByEmail(userParams.Email)
	if user != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(ErrorResponse("Email Sudah Terdaftar"))
		return
	}

	data, err := models.CreateUser(*userParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorResponse(err.Error()))
		return
	}
	response = SuccessResponse(data)

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var params *models.User
	var jsonResponse []byte
	var err error
	var token string
	var responseUser *models.User

	w.Header().Set("Content-Type", "application/json")
	err = SetParams(r.Body, &params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrorResponse(err.Error()))
		return
	}

	responseUser, err = models.GetUserByEmailAndPassword(params.Email, params.Password)
	if err == nil {
		if responseUser != nil {
			token = middlewares.GenerateToken(params.Email, params.Password)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(ErrorResponse("User Not Found"))
			return
		}

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrorResponse(err.Error()))
		return
	}

	// Create a struct to hold your JSON response
	type LoginResponse struct {
		Authorization string       `json:"auth_token,omitempty"`
		Users         *models.User `json:"users,omitempty"`
	}

	response := LoginResponse{
		Authorization: token,
		Users:         responseUser,
	}

	jsonResponse = SuccessResponse(response)

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
