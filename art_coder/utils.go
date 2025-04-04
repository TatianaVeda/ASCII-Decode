package main

import (
	"os"
)

// ReadFile reads the contents of a file and returns it as a string.

func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
