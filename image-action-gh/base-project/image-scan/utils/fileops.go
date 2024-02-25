package utils

import (
	"os"
)

// ReadFile reads the file specified by filePath and returns its content as a byte slice.
// It takes a filePath string as a parameter and returns a byte slice and an error.
func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// WriteFile writes data to a file named by filename. If the file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile truncates it before writing.
//
// filePath string, data []byte
// error
func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644) // 0644 provides read/write permissions to the owner and read-only for others
}

func AppendToFile(filePath string, data []byte) error {
	// Open the file in append mode, or create the file if it does not exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
