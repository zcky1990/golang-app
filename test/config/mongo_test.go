package config

import (
    "testing"
	"golang_app/golangApp/config"
	"github.com/stretchr/testify/assert"
)

func TestConnectMongoDB(t *testing.T) {
    env := "test"
	err := config.ConnectMongoDB(env)
	assert.Nil(t, err, "expected no error")
}

func TestDisconnectMongoDB(t *testing.T) {
	config.DisconnectMongoDB()
	assert.Nil(t, config.GetClient(), "expected client to be nil after disconnection")
}