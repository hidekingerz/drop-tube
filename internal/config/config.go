// Package config provides configuration management for DropTube.
// It contains all the necessary parameters for customizing the download process.
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	DEFAULT_OUTPUT_DIR   = "."
	DEFAULT_FORMAT       = "best"
	DEFAULT_AUDIO_FORMAT = "mp3"
	DEFAULT_QUALITY      = "best"
	DEFAULT_VERBOSE      = false
	DEFAULT_AUDIO_ONLY   = false
	DEFAULT_PLAYLIST     = false
)

// Config represents the configuration for video downloading.
// It contains all the necessary parameters for customizing the download process.
type Config struct {
	OutputDir   string
	Format      string
	Quality     string
	AudioOnly   bool
	AudioFormat string
	Playlist    bool
	Verbose     bool
	URL         string
}

// NewConfig creates a new configuration with default values.
func NewConfig() *Config {
	return &Config{
		OutputDir:   DEFAULT_OUTPUT_DIR,
		Format:      DEFAULT_FORMAT,
		Quality:     DEFAULT_QUALITY,
		AudioOnly:   DEFAULT_AUDIO_ONLY,
		AudioFormat: DEFAULT_AUDIO_FORMAT,
		Playlist:    DEFAULT_PLAYLIST,
		Verbose:     DEFAULT_VERBOSE,
	}
}

// Validate validates the configuration parameters.
func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("youtube URL is required")
	}

	if c.OutputDir != "" {
		absPath, err := filepath.Abs(c.OutputDir)
		if err != nil {
			return fmt.Errorf("invalid output directory path: %w", err)
		}
		c.OutputDir = absPath

		if err := c.ensureOutputDir(); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	return nil
}

// ensureOutputDir creates the output directory if it doesn't exist.
func (c *Config) ensureOutputDir() error {
	if _, err := os.Stat(c.OutputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(c.OutputDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", c.OutputDir, err)
		}
	}
	return nil
}
