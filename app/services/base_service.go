package services

import (
	"encoding/json"

	"github.com/google/uuid"
)

// BaseService is the small contract shared by every service object, the Go
// equivalent of the concerns Rails models/services usually share.
type BaseService interface {
	// ConvertToMap turns a model into a column -> value map for partial
	// updates, dropping empty fields. It replaces the old ConvertToBSON helper
	// now that the application talks to PostgreSQL through GORM.
	ConvertToMap(data interface{}) (map[string]interface{}, error)
}

// StructToMap marshals a struct to JSON and keeps only the non-empty keys, so
// GORM's Updates() behaves like Rails' partial update (`record.update(params)`)
// and never overwrites columns the caller did not send.
func StructToMap(data interface{}) (map[string]interface{}, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	values := map[string]interface{}{}
	if err := json.Unmarshal(payload, &values); err != nil {
		return nil, err
	}
	for key, value := range values {
		if value == "" || value == nil {
			delete(values, key)
		}
	}
	return values, nil
}

// ValidUUID reports whether the given string is a well formed UUID. PostgreSQL
// rejects any other value for a uuid column, so services validate first.
func ValidUUID(value string) bool {
	_, err := uuid.Parse(value)
	return err == nil
}

// NormalizeUUID returns nil for an empty or invalid UUID so the column is
// stored as NULL instead of raising a cast error.
func NormalizeUUID(value string) *string {
	if !ValidUUID(value) {
		return nil
	}
	return &value
}
