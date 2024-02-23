package controller

import (
	"fmt"
	"net/http"
	"golang_app/golangApp/models"
)

func Index(w http.ResponseWriter, r *http.Request) {

	userModel := models.User{
        Username:  "JohnDoe",
        Email: "john@example.com",
        Firstname:   "John",
		Lastname: "Doe",
    }

	data,err := models.AddUser(userModel)
	if err != nil {
		fmt.Fprintf(w, "Hi there, I love %s!", err.Error())
	}else {
		fmt.Fprintf(w, "Hi there, I love %s!", data)
	}
}