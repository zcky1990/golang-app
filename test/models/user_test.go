package models

import (
	"testing"
    "github.com/stretchr/testify/assert"
	"fmt"
	"log"
    "golang_app/golangApp/models"
	"golang_app/golangApp/config"
)

func init() {
	err := config.ConnectMongoDB("test")
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }
}

func TestAddUser(t *testing.T) {
    user :=models.User{
        Username:  "testuser",
        Email:     "test@example.com",
        Firstname: "Test",
        Lastname:  "User",
        Authtoken: "testtoken",
    }
	result,err :=models.AddUser(user)
    assert.Nil(t, err, "expected no error")
    assert.NotNil(t, result, "expected insert result not to be nil")
}

func TestGetUserByEmail(t *testing.T) {
	email := "test@example.com"
	result := models.GetUserByEmail(email)
	assert.Equal(t, result.Username, "testuser", "they should be equal")
	assert.NotNil(t, result, "expected Query result not to be nil")
}

func TestGetAllUserList(t *testing.T) {
	result,err := models.GetAllUserList(1,10)
	assert.Nil(t, err, "expected no error")
	assert.NotEmpty(t, result,  "expected Query result not to be empty")
}	

func TestSearchUser(t *testing.T) {
	//search by firstname
	result := models.SearchUser("name", "test")
	assert.NotEmpty(t, result,  "expected Query result not to be empty")
	//search by lastname
	result = models.SearchUser("name", "User")
	assert.NotEmpty(t, result, "expected Query result not to be nil")
	//search random name
	result = models.SearchUser("name", "asdw")
	assert.Nil(t, result, "expected Query result to be empty")
	//search by email
	result = models.SearchUser("email", "test@example.com")
	assert.NotNil(t, result, "expected Query result not to be nil")
}

func TestDeleteUserById(t *testing.T) {
	email := "test@example.com"
	result := models.GetUserByEmail(email)
	id := result.Id.Hex()
	fmt.Println(id)
	models.DeleteUserById(id)
	result = models.GetUserByEmail(email)
	assert.Nil(t, result, "expected Query result to be nil")
}