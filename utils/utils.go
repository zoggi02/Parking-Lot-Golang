package utils

import (
	"fmt"
	"os"
)

// ReadFile reads the file content for commands
func ReadFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return file, nil
}
