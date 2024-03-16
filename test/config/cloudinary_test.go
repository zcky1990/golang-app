package config

import (
	"bytes"
	"mime/multipart"
	"testing"

	"golang_app/golangApp/config"

	"github.com/stretchr/testify/assert"
)

// MockFile implements the multipart.File interface
type MockFile struct {
	*bytes.Reader
}

func (m *MockFile) Close() error {
	return nil
}

// Function to create a mock multipart file with image content
func createMockMultipartFile(content []byte) multipart.File {
	return &MockFile{bytes.NewReader(content)}
}

func TestUploadImageToFolder(t *testing.T) {
	expectedPublicId := "test_folder/test.jpg"
	expectedSecureUrl := "https://res.cloudinary.com/dzzp8z3ad/raw/upload/v1710605591/test_folder/test.jpg"

	// Initialize Cloudinary
	config.InitializeCloudinary("test")

	// Create a mock multipart file with image content
	mockFile := createMockMultipartFile([]byte("mock image content"))

	filename := "test.jpg"
	folder := "test_folder"

	// Call the UploadImageToFolder function
	resp, err := config.UploadImageToFolder(mockFile, filename, folder)

	//expected to not have error
	assert.Nil(t, err, "expected no error")

	//expected to have equal value for public_id and secure_id
	assert.EqualValues(t, expectedPublicId, resp.PublicId, "public_id should be equal")
	assert.EqualValues(t, expectedSecureUrl, resp.SecureUrl, "secure_url should be equal")

}
func TestUploadImage(t *testing.T) {

	expectedPublicId := "test.jpg"
	expectedSecureUrl := "https://res.cloudinary.com/dzzp8z3ad/raw/upload/v1710605591/test.jpg"

	// Initialize Cloudinary
	config.InitializeCloudinary("test")

	// Create a mock multipart file with image content
	mockFile := createMockMultipartFile([]byte("mock image content"))

	filename := "test.jpg"

	// Call the UploadImage function
	resp, err := config.UploadImage(mockFile, filename)

	//expected to not have error
	assert.Nil(t, err, "expected no error")

	//expected to have equal value for public_id and secure_id
	assert.EqualValues(t, expectedPublicId, resp.PublicId, "public_id should be equal")
	assert.EqualValues(t, expectedSecureUrl, resp.SecureUrl, "secure_url should be equal")
}
