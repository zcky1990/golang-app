package controller

import (
	"encoding/json"
	"io"
)

type successResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(data interface{}) []byte {
	response := successResponse{
		Status: "success",
		Data:   data,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return nil
	}

	return jsonData
}

type failedResponse struct {
	Status       string      `json:"status"`
	ErrorMessage interface{} `json:"error_message"`
}

func ErrorResponse(erroMessage string) []byte {
	response := failedResponse{
		Status:       "Error",
		ErrorMessage: erroMessage,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return nil
	}

	return jsonData
}

func SetParams(body io.Reader, v interface{}) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
