// Package downloader provides YouTube video downloading functionality.
// It supports various formats and quality options.
package downloader

import (
	"fmt"
	"log"
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
		if d.config.Format != "best" {
			args = append(args, "--format", d.config.Format)
		}
	}

	if d.config.Quality != "best" && !d.config.AudioOnly {
		args = append(args, "--format-sort", fmt.Sprintf("height:%s", d.config.Quality))
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

	args = append(args, d.config.URL)

	return args
}