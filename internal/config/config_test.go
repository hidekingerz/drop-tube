package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()

	if cfg.OutputDir != DEFAULT_OUTPUT_DIR {
		t.Errorf("NewConfig() OutputDir = %v, want %v", cfg.OutputDir, DEFAULT_OUTPUT_DIR)
	}
	if cfg.Format != DEFAULT_FORMAT {
		t.Errorf("NewConfig() Format = %v, want %v", cfg.Format, DEFAULT_FORMAT)
	}
	if cfg.Quality != DEFAULT_QUALITY {
		t.Errorf("NewConfig() Quality = %v, want %v", cfg.Quality, DEFAULT_QUALITY)
	}
	if cfg.AudioOnly != DEFAULT_AUDIO_ONLY {
		t.Errorf("NewConfig() AudioOnly = %v, want %v", cfg.AudioOnly, DEFAULT_AUDIO_ONLY)
	}
	if cfg.AudioFormat != DEFAULT_AUDIO_FORMAT {
		t.Errorf("NewConfig() AudioFormat = %v, want %v", cfg.AudioFormat, DEFAULT_AUDIO_FORMAT)
	}
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				URL:       "https://youtube.com/watch?v=123",
				OutputDir: ".",
			},
			wantErr: false,
		},
		{
			name: "missing URL",
			config: &Config{
				OutputDir: ".",
			},
			wantErr: true,
		},
		{
			name: "invalid output directory",
			config: &Config{
				URL:       "https://youtube.com/watch?v=123",
				OutputDir: "/root/invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_ensureOutputDir(t *testing.T) {
	tmpDir := os.TempDir()
	testDir := filepath.Join(tmpDir, "drop-tube-test")

	defer os.RemoveAll(testDir)

	cfg := &Config{
		OutputDir: testDir,
	}

	err := cfg.ensureOutputDir()
	if err != nil {
		t.Errorf("ensureOutputDir() error = %v", err)
	}

	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Errorf("ensureOutputDir() failed to create directory")
	}
}