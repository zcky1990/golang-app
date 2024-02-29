package controller

import (
	"fmt"
	"net/http"
	"golang_app/golangApp/models"
	"golang_app/golangApp/middlewares"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {

	// userModel := models.User{
    //     Username:  "JohnDoe",
    //     Email: "john@example.com",
    //     Firstname:   "John",
	// 	Lastname: "Doe",
    // }

	// data,err := models.AddUser(userModel)
	// if err != nil {
	// 	fmt.Fprintf(w, "Hi there, I love %s!", err.Error())
	// }else {
	// 	fmt.Fprintf(w, "Hi there, I love %s!", data)
	// }
	fmt.Fprintf(w, "Hi there, I love !")
}


type LoginParams struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var params LoginParams
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
	if err == nil  {
		if (responseUser != nil) {
			token = middlewares.GenerateToken(params.Email, params.Password)
		}else {
			w.WriteHeader(http.StatusOK)
			w.Write(ErrorResponse("User Not Found"))
			return
		}
		
	}else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrorResponse(err.Error()))
		return
	}

	// Create a struct to hold your JSON response
    type Response struct {
        Authorization string `json:"auth_token,omitempty"`
		Users *models.User `json:"users,omitempty"`
    }

	response := Response {
        Authorization: token,
        Users: responseUser,
    }

	jsonResponse, err = json.Marshal(response)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
    }

    // Write the JSON response
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}