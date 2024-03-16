package lib

import (
	"os"
	"path/filepath"
	"testing"

	"golang_app/golangApp/lib"

	"github.com/stretchr/testify/assert"
)

func TestFindRootDir(t *testing.T) {
	// Create a temporary directory structure for testing
	tmpDir := createTempDir(t)
	defer os.RemoveAll(tmpDir) // Clean up temporary directory

	// Test case 1: Root directory contains go.mod file
	rootDir := filepath.Join(tmpDir, "project")
	createFile(t, filepath.Join(rootDir, "go.mod"))
	expectedRoot := rootDir
	testFindRootDir(t, rootDir, expectedRoot)

	// Test case 2: Root directory contains main.go file
	rootDir = filepath.Join(tmpDir, "project2")
	createFile(t, filepath.Join(rootDir, "main.go"))
	expectedRoot = rootDir
	testFindRootDir(t, rootDir, expectedRoot)

	// Test case 3: Root directory contains both go.mod and main.go files
	rootDir = filepath.Join(tmpDir, "project3")
	createFile(t, filepath.Join(rootDir, "go.mod"))
	createFile(t, filepath.Join(rootDir, "main.go"))
	expectedRoot = rootDir
	testFindRootDir(t, rootDir, expectedRoot)

	// Test case 4: No known project items found (root directory not found)
	rootDir = "/"
	expectedRoot = ""
	testFindRootDir(t, rootDir, expectedRoot)
}

// Helper function to test FindRootDir
func testFindRootDir(t *testing.T, inputDir, expectedRoot string) {
	t.Helper()
	result := lib.FindRootDir(inputDir)
	// log.Println("RESULT :", result)
	//expected to have equal stored redis value
	assert.EqualValues(t, expectedRoot, result, "directory should be equal")
}

// Helper function to create a temporary directory for testing
func createTempDir(t *testing.T) string {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	return tmpDir
}

// Helper function to create a file for testing
func createFile(t *testing.T, filePath string) {
	t.Helper()

	// Create the directory structure if it doesn't exist
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file %s: %v", filePath, err)
	}
	defer file.Close()
}
