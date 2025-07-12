// Package downloader provides YouTube video downloading functionality.
// It supports various formats and quality options.
package downloader

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"

	"github.com/hidekingerz/drop-tube/internal/config"
)

// Downloader handles YouTube video downloads using yt-dlp.
type Downloader struct {
	config *config.Config
}

// New creates a new Downloader instance with the given configuration.
func New(cfg *config.Config) *Downloader {
	return &Downloader{
		config: cfg,
	}
}

// Download downloads a YouTube video from the given URL.
// It returns an error if the download fails.
func (d *Downloader) Download() error {
	if err := d.checkYtDlpInstalled(); err != nil {
		return fmt.Errorf("yt-dlp dependency check failed: %w", err)
	}

	if d.config.Verbose {
		log.Printf("starting download with config: %+v", d.config)
	}

	args := d.buildYtDlpArgs()

	if d.config.Verbose {
		log.Printf("executing: yt-dlp %s", strings.Join(args, " "))
	}

	cmd := exec.Command("yt-dlp", args...)
	cmd.Dir = d.config.OutputDir

	if d.config.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		bar := progressbar.Default(-1, "downloading video...")
		defer bar.Finish()
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yt-dlp execution failed: %w", err)
	}

	fmt.Printf("download completed successfully in %s\n", d.config.OutputDir)
	return nil
}

// checkYtDlpInstalled verifies that yt-dlp is installed and accessible.
func (d *Downloader) checkYtDlpInstalled() error {
	cmd := exec.Command("yt-dlp", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yt-dlp not found or not executable: %w", err)
	}
	return nil
}

// buildYtDlpArgs constructs the command line arguments for yt-dlp.
func (d *Downloader) buildYtDlpArgs() []string {
	args := []string{}

	if d.config.AudioOnly {
		args = append(args, "--extract-audio")
		args = append(args, "--audio-format", d.config.AudioFormat)
	} else {
		formatSpec := d.buildFormatSpec()
		if formatSpec != "" {
			args = append(args, "--format", formatSpec)
		}
	}

	if d.config.Playlist {
		args = append(args, "--yes-playlist")
	} else {
		args = append(args, "--no-playlist")
	}

	outputTemplate := filepath.Join(d.config.OutputDir, "%(title)s.%(ext)s")
	args = append(args, "--output", outputTemplate)

	if !d.config.Verbose {
		args = append(args, "--quiet", "--no-warnings")
	}

	cleanURL := d.cleanURL(d.config.URL)
	args = append(args, cleanURL)

	return args
}

// buildFormatSpec constructs the format specification for yt-dlp.
func (d *Downloader) buildFormatSpec() string {
	if d.config.Format != "best" && d.config.Quality != "best" {
		// Both format and quality specified
		height := d.extractHeight(d.config.Quality)
		return fmt.Sprintf("bestvideo[ext=%s][height<=%s]+bestaudio/best[ext=%s][height<=%s]", 
			d.config.Format, height, d.config.Format, height)
	} else if d.config.Quality != "best" {
		// Only quality specified
		height := d.extractHeight(d.config.Quality)
		return fmt.Sprintf("bestvideo[height<=%s]+bestaudio/best[height<=%s]", height, height)
	} else if d.config.Format != "best" {
		// Only format specified
		return fmt.Sprintf("best[ext=%s]", d.config.Format)
	}
	// Default to best
	return "best"
}

// extractHeight extracts the height value from quality string (e.g., "1080p" -> "1080").
func (d *Downloader) extractHeight(quality string) string {
	// Remove 'p' suffix if present
	if strings.HasSuffix(quality, "p") {
		return quality[:len(quality)-1]
	}
	return quality
}

// cleanURL removes shell escaping and normalizes the URL.
func (d *Downloader) cleanURL(rawURL string) string {
	// Remove shell escaping (backslashes before special characters)
	cleaned := strings.ReplaceAll(rawURL, "\\", "")
	
	// Try to parse and validate the URL
	if parsedURL, err := url.Parse(cleaned); err == nil {
		return parsedURL.String()
	}
	
	// If parsing fails, return the cleaned URL as-is
	return cleaned
}