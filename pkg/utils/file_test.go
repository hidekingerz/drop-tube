package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureDir(t *testing.T) {
	tmpDir := os.TempDir()
	testDir := filepath.Join(tmpDir, "drop-tube-utils-test")

	defer os.RemoveAll(testDir)

	err := EnsureDir(testDir)
	if err != nil {
		t.Errorf("EnsureDir() error = %v", err)
	}

	if !IsValidDir(testDir) {
		t.Errorf("EnsureDir() failed to create directory")
	}
}

func TestIsValidDir(t *testing.T) {
	tests := []struct {
		name     string
		dirPath  string
		expected bool
	}{
		{
			name:     "valid directory",
			dirPath:  os.TempDir(),
			expected: true,
		},
		{
			name:     "non-existent directory",
			dirPath:  "/non/existent/path",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidDir(tt.dirPath)
			if result != tt.expected {
				t.Errorf("IsValidDir() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetAbsolutePath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "valid relative path",
			path:    ".",
			wantErr: false,
		},
		{
			name:    "valid absolute path",
			path:    "/tmp",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetAbsolutePath(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAbsolutePath() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !filepath.IsAbs(result) {
				t.Errorf("GetAbsolutePath() returned non-absolute path: %v", result)
			}
		})
	}
}