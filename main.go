package main

import (
	cfg "golang_app/golangApp/config"
	e "golang_app/golangApp/config/environments"
)

func initializeEnvirotmentConfigurationAndEnvirotment(env string) e.EnvironmentConfiguration {
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

func main() {
	app := cfg.RoutesNew()
	app.StartServer()
}
