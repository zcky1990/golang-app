package config

import (
	"golang_app/golangApp/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectMongoDB(t *testing.T) {
	// Initialize Cloudinary  with test env
	err := config.ConnectMongoDB("test")
	assert.Nil(t, err, "expected no error")
}

func TestDisconnectMongoDB(t *testing.T) {
	config.DisconnectMongoDB()
	assert.Nil(t, config.GetClient(), "expected client to be nil after disconnection")
}
