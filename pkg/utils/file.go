package utils

import "os"

func ReadFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	return data, err
}

func WriteFile(filePath string, content []byte) bool {
	err := os.WriteFile(filePath, content, 0644)
	return err == nil
}
