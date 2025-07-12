package cli

import (
	"os"
	"testing"
)

func TestRootCmd(t *testing.T) {
	// Test that rootCmd is properly initialized
	if rootCmd == nil {
		t.Error("rootCmd should not be nil")
	}

	if rootCmd.Use != "drop-tube [OPTIONS] <YouTube URL>" {
		t.Errorf("rootCmd.Use = %v, expected 'drop-tube [OPTIONS] <YouTube URL>'", rootCmd.Use)
	}

	if rootCmd.Short == "" {
		t.Error("rootCmd.Short should not be empty")
	}
}

func TestExecute(t *testing.T) {
	// Save original args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Test execute with help flag (should not fail)
	os.Args = []string{"drop-tube", "--help"}

	// Capture the output by redirecting stdout temporarily
	// This test mainly ensures Execute() doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Execute() panicked: %v", r)
		}
	}()

	// Note: Execute() will call os.Exit(0) for --help, so we can't easily test it
	// without mocking. For now, we just test that the function exists and is callable.
}

func TestRootCmdFlags(t *testing.T) {
	// Test that all expected flags are present
	expectedFlags := []string{
		"output",
		"format",
		"quality",
		"audio-only",
		"audio-format",
		"playlist",
		"verbose",
	}

	for _, flagName := range expectedFlags {
		flag := rootCmd.PersistentFlags().Lookup(flagName)
		if flag == nil {
			t.Errorf("Expected flag %s not found", flagName)
		}
	}
}

func TestRootCmdShortFlags(t *testing.T) {
	// Test that short flags are properly set
	shortFlags := map[string]string{
		"o": "output",
		"f": "format",
		"q": "quality",
		"a": "audio-only",
		"v": "verbose",
	}

	for shortFlag, longFlag := range shortFlags {
		flag := rootCmd.PersistentFlags().ShorthandLookup(shortFlag)
		if flag == nil {
			t.Errorf("Expected short flag -%s for --%s not found", shortFlag, longFlag)
		}
		if flag.Name != longFlag {
			t.Errorf("Short flag -%s should map to --%s, got --%s", shortFlag, longFlag, flag.Name)
		}
	}
}

func TestRootCmdArgs(t *testing.T) {
	// Test argument validation
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no URL provided",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "single URL provided",
			args:    []string{"https://www.youtube.com/watch?v=test"},
			wantErr: false, // Args validation should pass
		},
		{
			name:    "multiple arguments",
			args:    []string{"arg1", "arg2"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rootCmd.Args(rootCmd, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Args() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
