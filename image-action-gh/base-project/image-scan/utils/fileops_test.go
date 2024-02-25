package utils

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Setup: Create a temporary file with some content
	tempFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name()) // clean up

	content := []byte("hello world")
	if _, err := tempFile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test: Read the content back
	readContent, err := ReadFile(tempFile.Name())
	if err != nil {
		t.Errorf("ReadFile failed: %v", err)
	}
	if string(readContent) != string(content) {
		t.Errorf("Expected %s, got %s", string(content), string(readContent))
	}
}

//  write the code to test the WriteFile function

func TestWriteFile(t *testing.T) {
	// Create a temporary file for testing
	tempFile := "testfile.txt"

	// Defer removal of the temporary file
	defer func() {
		err := os.Remove(tempFile)
		if err != nil {
			t.Errorf("Error removing temporary file: %v", err)
		}
	}()

	// Test data
	data := []byte("Hello, World!")

	// Call the function being tested
	err := WriteFile(tempFile, data)
	if err != nil {
		t.Errorf("Error writing file: %v", err)
	}

	// Read the written file to verify its content
	content, err := os.ReadFile(tempFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	// Verify the content
	if string(content) != "Hello, World!" {
		t.Errorf("File content does not match expected value")
	}
}
