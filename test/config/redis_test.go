package config

import (
	"testing"

	"golang_app/golangApp/config"

	"github.com/stretchr/testify/assert"
)

func TestStoreAndGetRedis(t *testing.T) {
	expectedValue := "test"

	// Initialize Cloudinary with test env
	config.InitializeRedis("test")

	//store data to redis with no expired
	err := config.SetRedis("test", "test", 0)

	//expected to not have error when store data
	assert.Nil(t, err, "expected no error")

	val, err := config.GetRedis("test")

	//expected to not have error when get data
	assert.Nil(t, err, "expected no error")

	//expected to have equal stored redis value
	assert.EqualValues(t, expectedValue, val, "value should be equal")

}
