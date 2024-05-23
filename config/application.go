package config

import (
	"os"
	"sync"

	e "golang_app/golangApp/config/environments"
)

type Application struct {
	EnvConfig e.EnvironmentConfiguration
}

var (
	instance *Application
	once     sync.Once
)

func initializeEnvirotmentConfigurationAndEnvirotment() e.EnvironmentConfiguration {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	var config e.EnvironmentConfiguration
	if env == "development" {
		config = e.NewDevEnvConfiguration(env)
	}
	if env == "test" {
		config = e.NewTestEnvConfiguration(env)
	}
	if env == "production" {
		config = e.NewProdEnvConfiguration(env)
	}
	return config
}

func GetApplicationInstance() *Application {
	once.Do(func() {
		config := initializeEnvirotmentConfigurationAndEnvirotment()
		config.LoadEvirontmentFile()
		instance = &Application{EnvConfig: config}
	})
	return instance
}
