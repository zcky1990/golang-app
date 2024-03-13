package models

import (
	"golang_app/golangApp/config"
	"golang_app/golangApp/models"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	err := config.ConnectMongoDB("test")
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	config.DisconnectMongoDB()
	os.Exit(exitVal)
}

func TestCreateUser(t *testing.T) {
	user := models.User{
		Username:  "testuser",
		Email:     "test@example.com",
		Firstname: "Test",
		Lastname:  "User",
		Authtoken: "testtoken",
	}
	result, err := models.CreateUser(user)
	assert.Nil(t, err, "expected no error")
	assert.NotNil(t, result, "expected insert result not to be nil")

	result, err = models.CreateUser(user)
	assert.NotNil(t, err, "expected Got error")
}

func TestGetUserByEmail(t *testing.T) {
	//search user with correct email
	email := "test@example.com"
	result := models.GetUserByEmail(email)
	assert.NotNil(t, result, "expected Query result not to be nil")
	assert.Equal(t, result.Username, "testuser", "they should be equal")

	//search user with incorrect email
	email = "test@example.comassdasd"
	result = models.GetUserByEmail(email)
	assert.Nil(t, result, "expected Query result to be nil")
}

func TestGetAllUserList(t *testing.T) {
	//get user on page 1, max 10
	result, err := models.GetAllUserList(1, 10)
	assert.Nil(t, err, "expected no error")
	assert.NotEmpty(t, result, "expected Query result not to be empty")

	//get user on page 2, max 10
	result, err = models.GetAllUserList(2, 10)
	assert.Nil(t, err, "expected no error")
	assert.Nil(t, result, "expected Query result to be empty")
}

func TestSearchUser(t *testing.T) {
	//search by firstname
	result := models.SearchUser("name", "test")
	assert.NotEmpty(t, result, "expected Query result not to be empty")

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

func TestUpdateUser(t *testing.T) {
	var user models.User
	email := "test@example.com"
	result := models.GetUserByEmail(email)
	id := result.Id.Hex()

	user = *result
	user.Firstname = "Nurfadillah"

	data, err := models.UpdateUserById(id, user)
	assert.Nil(t, err, "expected error to be empty")
	assert.NotNil(t, data, "expected result not to be nil")
}

func TestDeleteUserById(t *testing.T) {
	//search user by email
	email := "test@example.com"
	result := models.GetUserByEmail(email)
	id := result.Id.Hex()

	//delete user by its id
	err := models.DeleteUserById(id)
	assert.Nil(t, err, "expected Error to be nil")

	//check if we can get user by email
	result = models.GetUserByEmail(email)
	assert.Nil(t, result, "expected Query result to be nil")
}
