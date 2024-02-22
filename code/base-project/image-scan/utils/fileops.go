package utils

import (
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644) // 0644 provides read/write permissions to the owner and read-only for others
}
