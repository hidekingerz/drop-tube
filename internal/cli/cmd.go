// Package cli provides command-line interface functionality for DropTube.
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hidekingerz/drop-tube/internal/config"
	"github.com/hidekingerz/drop-tube/internal/downloader"
)

var cfg *config.Config

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "drop-tube [OPTIONS] <YouTube URL>",
	Short: "Download YouTube videos to local storage",
	Long: `DropTube is a command-line tool for downloading YouTube videos.
It supports various formats and quality options while respecting YouTube's terms of service.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg.URL = args[0]

		if err := cfg.Validate(); err != nil {
			return fmt.Errorf("configuration validation failed: %w", err)
		}

		downloader := downloader.New(cfg)
		return downloader.Download()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	cfg = config.NewConfig()

	rootCmd.PersistentFlags().StringVarP(&cfg.OutputDir, "output", "o", cfg.OutputDir, "output directory")
	rootCmd.PersistentFlags().StringVarP(&cfg.Format, "format", "f", cfg.Format, "video format (mp4, webm, best)")
	rootCmd.PersistentFlags().StringVarP(&cfg.Quality, "quality", "q", cfg.Quality, "video quality (720p, 1080p, best)")
	rootCmd.PersistentFlags().BoolVarP(&cfg.AudioOnly, "audio-only", "a", cfg.AudioOnly, "download audio only")
	rootCmd.PersistentFlags().StringVar(&cfg.AudioFormat, "audio-format", cfg.AudioFormat, "audio format (mp3, m4a)")
	rootCmd.PersistentFlags().BoolVar(&cfg.Playlist, "playlist", cfg.Playlist, "download entire playlist")
	rootCmd.PersistentFlags().BoolVarP(&cfg.Verbose, "verbose", "v", cfg.Verbose, "verbose output")
}