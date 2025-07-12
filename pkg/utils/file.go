// Package utils provides utility functions for file operations.
package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// EnsureDir creates the directory if it doesn't exist.
func EnsureDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}
	}
	return nil
}

// IsValidDir checks if the given path is a valid directory.
func IsValidDir(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// GetAbsolutePath returns the absolute path of the given path.
func GetAbsolutePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for %s: %w", path, err)
	}
	return absPath, nil
}