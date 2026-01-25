package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.DefaultProfile != "default" {
		t.Errorf("Expected default profile 'default', got '%s'", cfg.DefaultProfile)
	}

	if cfg.DefaultOutputFormat != types.OutputFormatJSON {
		t.Errorf("Expected default output format 'json', got '%s'", cfg.DefaultOutputFormat)
	}

	if cfg.DefaultFields != FieldMaskStandard {
		t.Errorf("Expected default fields 'standard', got '%s'", cfg.DefaultFields)
	}

	if cfg.CacheTTL != 300 {
		t.Errorf("Expected cache TTL 300, got %d", cfg.CacheTTL)
	}

	if cfg.MaxRetries != 3 {
		t.Errorf("Expected max retries 3, got %d", cfg.MaxRetries)
	}

	if cfg.LogLevel != "normal" {
		t.Errorf("Expected log level 'normal', got '%s'", cfg.LogLevel)
	}
}

func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name      string
		config    *Config
		wantError bool
		errorMsg  string
	}{
		{
			name:      "valid default config",
			config:    DefaultConfig(),
			wantError: false,
		},
		{
			name: "invalid output format",
			config: &Config{
				DefaultProfile:      "default",
				DefaultOutputFormat: types.OutputFormat("invalid"),
				DefaultFields:       FieldMaskStandard,
				CacheTTL:            300,
				MaxRetries:          3,
				RetryBaseDelay:      1000,
				RequestTimeout:      60,
				LogLevel:            "normal",
			},
			wantError: true,
			errorMsg:  "invalid output format",
		},
		{
			name: "invalid field mask preset",
			config: &Config{
				DefaultProfile:      "default",
				DefaultOutputFormat: types.OutputFormatJSON,
				DefaultFields:       FieldMaskPreset("invalid"),
				CacheTTL:            300,
				MaxRetries:          3,
				RetryBaseDelay:      1000,
				RequestTimeout:      60,
				LogLevel:            "normal",
			},
			wantError: true,
			errorMsg:  "invalid field mask preset",
		},
		{
			name: "negative cache TTL",
			config: &Config{
				DefaultProfile:      "default",
				DefaultOutputFormat: types.OutputFormatJSON,
				DefaultFields:       FieldMaskStandard,
				CacheTTL:            -1,
				MaxRetries:          3,
				RetryBaseDelay:      1000,
				RequestTimeout:      60,
				LogLevel:            "normal",
			},
			wantError: true,
			errorMsg:  "cache TTL must be non-negative",
		},
		{
			name: "max retries too high",
			config: &Config{
				DefaultProfile:      "default",
				DefaultOutputFormat: types.OutputFormatJSON,
				DefaultFields:       FieldMaskStandard,
				CacheTTL:            300,
				MaxRetries:          11,
				RetryBaseDelay:      1000,
				RequestTimeout:      60,
				LogLevel:            "normal",
			},
			wantError: true,
			errorMsg:  "max retries must be between 0 and 10",
		},
		{
			name: "retry base delay too low",
			config: &Config{
				DefaultProfile:      "default",
				DefaultOutputFormat: types.OutputFormatJSON,
				DefaultFields:       FieldMaskStandard,
				CacheTTL:            300,
				MaxRetries:          3,
				RetryBaseDelay:      50,
				RequestTimeout:      60,
				LogLevel:            "normal",
			},
			wantError: true,
			errorMsg:  "retry base delay must be between 100ms and 60000ms",
		},
		{
			name: "request timeout out of range",
			config: &Config{
				DefaultProfile:      "default",
				DefaultOutputFormat: types.OutputFormatJSON,
				DefaultFields:       FieldMaskStandard,
				CacheTTL:            300,
				MaxRetries:          3,
				RetryBaseDelay:      1000,
				RequestTimeout:      3700,
				LogLevel:            "normal",
			},
			wantError: true,
			errorMsg:  "request timeout must be between 1 and 3600 seconds",
		},
		{
			name: "invalid log level",
			config: &Config{
				DefaultProfile:      "default",
				DefaultOutputFormat: types.OutputFormatJSON,
				DefaultFields:       FieldMaskStandard,
				CacheTTL:            300,
				MaxRetries:          3,
				RetryBaseDelay:      1000,
				RequestTimeout:      60,
				LogLevel:            "invalid",
			},
			wantError: true,
			errorMsg:  "invalid log level",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantError {
				if err == nil {
					t.Errorf("Expected error containing '%s', got nil", tt.errorMsg)
				} else if !contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			}
		})
	}
}

func TestConfigDurationGetters(t *testing.T) {
	cfg := &Config{
		CacheTTL:       300,
		RetryBaseDelay: 1000,
		RequestTimeout: 60,
	}

	if d := cfg.GetCacheTTL(); d != 300*time.Second {
		t.Errorf("Expected cache TTL 300s, got %v", d)
	}

	if d := cfg.GetRetryBaseDelay(); d != 1000*time.Millisecond {
		t.Errorf("Expected retry base delay 1000ms, got %v", d)
	}

	if d := cfg.GetRequestTimeout(); d != 60*time.Second {
		t.Errorf("Expected request timeout 60s, got %v", d)
	}
}

func TestConfigSaveAndLoad(t *testing.T) {
	tempDir := t.TempDir()

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", originalHome)

	if runtime.GOOS == "windows" {
		originalUserProfile := os.Getenv("USERPROFILE")
		os.Setenv("USERPROFILE", tempDir)
		defer os.Setenv("USERPROFILE", originalUserProfile)
	}

	// Create a config with custom values
	cfg := &Config{
		DefaultProfile:      "test-profile",
		DefaultOutputFormat: types.OutputFormatTable,
		DefaultFields:       FieldMaskFull,
		CacheTTL:            600,
		IncludeExportLinks:  true,
		MaxRetries:          5,
		RetryBaseDelay:      2000,
		RequestTimeout:      120,
		LogLevel:            "verbose",
		ColorOutput:         false,
	}

	// Ensure config directory exists
	configDir, err := GetConfigDir()
	if err != nil {
		t.Fatalf("Failed to get config dir: %v", err)
	}
	if err := os.MkdirAll(configDir, 0700); err != nil {
		t.Fatalf("Failed to create config dir: %v", err)
	}

	// Save the config
	fullConfigPath := filepath.Join(configDir, ConfigFileName)
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(fullConfigPath, data, 0600); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	// Load the config
	loadedCfg := DefaultConfig()
	if err := loadedCfg.loadFromFile(); err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify loaded values
	if loadedCfg.DefaultProfile != cfg.DefaultProfile {
		t.Errorf("Expected profile '%s', got '%s'", cfg.DefaultProfile, loadedCfg.DefaultProfile)
	}

	if loadedCfg.DefaultOutputFormat != cfg.DefaultOutputFormat {
		t.Errorf("Expected output format '%s', got '%s'", cfg.DefaultOutputFormat, loadedCfg.DefaultOutputFormat)
	}

	if loadedCfg.CacheTTL != cfg.CacheTTL {
		t.Errorf("Expected cache TTL %d, got %d", cfg.CacheTTL, loadedCfg.CacheTTL)
	}

	if loadedCfg.IncludeExportLinks != cfg.IncludeExportLinks {
		t.Errorf("Expected include export links %v, got %v", cfg.IncludeExportLinks, loadedCfg.IncludeExportLinks)
	}
}

func TestLoadFromEnv(t *testing.T) {
	// Save original environment
	originalEnv := map[string]string{
		"GDRV_DEFAULT_PROFILE":      os.Getenv("GDRV_DEFAULT_PROFILE"),
		"GDRV_OUTPUT_FORMAT":        os.Getenv("GDRV_OUTPUT_FORMAT"),
		"GDRV_DEFAULT_FIELDS":       os.Getenv("GDRV_DEFAULT_FIELDS"),
		"GDRV_CACHE_TTL":            os.Getenv("GDRV_CACHE_TTL"),
		"GDRV_INCLUDE_EXPORT_LINKS": os.Getenv("GDRV_INCLUDE_EXPORT_LINKS"),
		"GDRV_MAX_RETRIES":          os.Getenv("GDRV_MAX_RETRIES"),
		"GDRV_LOG_LEVEL":            os.Getenv("GDRV_LOG_LEVEL"),
	}

	// Restore environment after test
	defer func() {
		for key, value := range originalEnv {
			if value == "" {
				os.Unsetenv(key)
			} else {
				os.Setenv(key, value)
			}
		}
	}()

	// Set test environment variables
	os.Setenv("GDRV_DEFAULT_PROFILE", "env-profile")
	os.Setenv("GDRV_OUTPUT_FORMAT", "table")
	os.Setenv("GDRV_DEFAULT_FIELDS", "full")
	os.Setenv("GDRV_CACHE_TTL", "900")
	os.Setenv("GDRV_INCLUDE_EXPORT_LINKS", "true")
	os.Setenv("GDRV_MAX_RETRIES", "7")
	os.Setenv("GDRV_LOG_LEVEL", "debug")

	// Load config (which should apply env vars)
	cfg := DefaultConfig()
	cfg.loadFromEnv()

	// Verify values from environment
	if cfg.DefaultProfile != "env-profile" {
		t.Errorf("Expected profile 'env-profile', got '%s'", cfg.DefaultProfile)
	}

	if cfg.DefaultOutputFormat != types.OutputFormatTable {
		t.Errorf("Expected output format 'table', got '%s'", cfg.DefaultOutputFormat)
	}

	if cfg.DefaultFields != FieldMaskFull {
		t.Errorf("Expected fields 'full', got '%s'", cfg.DefaultFields)
	}

	if cfg.CacheTTL != 900 {
		t.Errorf("Expected cache TTL 900, got %d", cfg.CacheTTL)
	}

	if !cfg.IncludeExportLinks {
		t.Error("Expected include export links to be true")
	}

	if cfg.MaxRetries != 7 {
		t.Errorf("Expected max retries 7, got %d", cfg.MaxRetries)
	}

	if cfg.LogLevel != "debug" {
		t.Errorf("Expected log level 'debug', got '%s'", cfg.LogLevel)
	}
}

func TestGetFieldMask(t *testing.T) {
	tests := []struct {
		name               string
		preset             FieldMaskPreset
		includeExportLinks bool
		wantContains       []string
	}{
		{
			name:   "minimal preset",
			preset: FieldMaskMinimal,
			wantContains: []string{
				"id",
				"name",
				"mimeType",
			},
		},
		{
			name:   "standard preset",
			preset: FieldMaskStandard,
			wantContains: []string{
				"id",
				"name",
				"mimeType",
				"size",
				"parents",
				"capabilities",
			},
		},
		{
			name:   "full preset",
			preset: FieldMaskFull,
			wantContains: []string{
				"id",
				"name",
				"mimeType",
				"size",
				"owners",
				"permissions",
				"shortcutDetails",
			},
		},
		{
			name:               "standard with export links",
			preset:             FieldMaskStandard,
			includeExportLinks: true,
			wantContains: []string{
				"id",
				"name",
				"exportLinks",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mask := GetFieldMask(tt.preset, tt.includeExportLinks)

			for _, field := range tt.wantContains {
				if !contains(mask, field) {
					t.Errorf("Expected field mask to contain '%s', got: %s", field, mask)
				}
			}

			// Verify exportLinks is only included when requested
			hasExportLinks := contains(mask, "exportLinks")
			if hasExportLinks != tt.includeExportLinks {
				t.Errorf("Expected exportLinks=%v, got %v in mask: %s", 
					tt.includeExportLinks, hasExportLinks, mask)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"true", true},
		{"True", true},
		{"TRUE", true},
		{"1", true},
		{"yes", true},
		{"YES", true},
		{"on", true},
		{"ON", true},
		{"false", false},
		{"False", false},
		{"0", false},
		{"no", false},
		{"off", false},
		{"", false},
		{"invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseBool(tt.input)
			if got != tt.want {
				t.Errorf("parseBool(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && 
		(s == substr || len(s) >= len(substr) && 
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || 
		len(s) > len(substr) && containsInner(s, substr)))
}

func containsInner(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
