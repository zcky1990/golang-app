package controllers

type BaseController interface {
	structToMap(data interface{}) map[string]interface{}
}
