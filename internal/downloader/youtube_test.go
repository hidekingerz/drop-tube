package downloader

import (
	"testing"

	"github.com/hidekingerz/drop-tube/internal/config"
)

func TestNew(t *testing.T) {
	cfg := config.NewConfig()
	cfg.URL = "https://www.youtube.com/watch?v=test"

	downloader := New(cfg)

	if downloader == nil {
		t.Error("New() returned nil")
	}

	if downloader.config != cfg {
		t.Error("New() did not set config correctly")
	}
}

func TestBuildFormatSpec(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		quality  string
		expected string
	}{
		{
			name:     "default best",
			format:   "best",
			quality:  "best",
			expected: "best",
		},
		{
			name:     "quality only",
			format:   "best",
			quality:  "1080p",
			expected: "bestvideo[height<=1080]+bestaudio/best[height<=1080]",
		},
		{
			name:     "format only",
			format:   "mp4",
			quality:  "best",
			expected: "best[ext=mp4]",
		},
		{
			name:     "both format and quality",
			format:   "mp4",
			quality:  "720p",
			expected: "bestvideo[ext=mp4][height<=720]+bestaudio/best[ext=mp4][height<=720]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewConfig()
			cfg.Format = tt.format
			cfg.Quality = tt.quality
			cfg.URL = "https://www.youtube.com/watch?v=test"

			d := New(cfg)
			result := d.buildFormatSpec()

			if result != tt.expected {
				t.Errorf("buildFormatSpec() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestExtractHeight(t *testing.T) {
	tests := []struct {
		name     string
		quality  string
		expected string
	}{
		{
			name:     "quality with p suffix",
			quality:  "1080p",
			expected: "1080",
		},
		{
			name:     "quality without p suffix",
			quality:  "720",
			expected: "720",
		},
		{
			name:     "best quality",
			quality:  "best",
			expected: "best",
		},
	}

	cfg := config.NewConfig()
	cfg.URL = "https://www.youtube.com/watch?v=test"
	d := New(cfg)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := d.extractHeight(tt.quality)
			if result != tt.expected {
				t.Errorf("extractHeight() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestCleanURL(t *testing.T) {
	tests := []struct {
		name     string
		rawURL   string
		expected string
	}{
		{
			name:     "clean URL",
			rawURL:   "https://www.youtube.com/watch?v=test",
			expected: "https://www.youtube.com/watch?v=test",
		},
		{
			name:     "URL with backslashes",
			rawURL:   "https://www.youtube.com/watch\\?v\\=test",
			expected: "https://www.youtube.com/watch?v=test",
		},
		{
			name:     "URL with multiple backslashes",
			rawURL:   "https://www.youtube.com/watch\\?v\\=P0YWWyeUTII",
			expected: "https://www.youtube.com/watch?v=P0YWWyeUTII",
		},
	}

	cfg := config.NewConfig()
	d := New(cfg)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := d.cleanURL(tt.rawURL)
			if result != tt.expected {
				t.Errorf("cleanURL() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestBuildYtDlpArgs(t *testing.T) {
	tests := []struct {
		name        string
		config      func() *config.Config
		contains    []string
		notContains []string
	}{
		{
			name: "default config",
			config: func() *config.Config {
				cfg := config.NewConfig()
				cfg.URL = "https://www.youtube.com/watch?v=test"
				return cfg
			},
			contains:    []string{"--no-playlist", "--newline", "--no-warnings"},
			notContains: []string{"--extract-audio", "--yes-playlist"},
		},
		{
			name: "audio only",
			config: func() *config.Config {
				cfg := config.NewConfig()
				cfg.URL = "https://www.youtube.com/watch?v=test"
				cfg.AudioOnly = true
				cfg.AudioFormat = "mp3"
				return cfg
			},
			contains:    []string{"--extract-audio", "--audio-format", "mp3"},
			notContains: []string{},
		},
		{
			name: "playlist",
			config: func() *config.Config {
				cfg := config.NewConfig()
				cfg.URL = "https://www.youtube.com/watch?v=test"
				cfg.Playlist = true
				return cfg
			},
			contains:    []string{"--yes-playlist"},
			notContains: []string{"--no-playlist"},
		},
		{
			name: "verbose",
			config: func() *config.Config {
				cfg := config.NewConfig()
				cfg.URL = "https://www.youtube.com/watch?v=test"
				cfg.Verbose = true
				return cfg
			},
			contains:    []string{"--newline"},
			notContains: []string{"--quiet", "--no-warnings"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(tt.config())
			args := d.buildYtDlpArgs()

			// Convert args to string for easier checking
			argsStr := ""
			for _, arg := range args {
				argsStr += arg + " "
			}

			for _, contain := range tt.contains {
				found := false
				for _, arg := range args {
					if arg == contain {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("buildYtDlpArgs() should contain %v, args: %v", contain, args)
				}
			}

			for _, notContain := range tt.notContains {
				for _, arg := range args {
					if arg == notContain {
						t.Errorf("buildYtDlpArgs() should not contain %v, args: %v", notContain, args)
					}
				}
			}
		})
	}
}
