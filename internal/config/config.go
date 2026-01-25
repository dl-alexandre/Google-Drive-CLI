package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/types"
)

const (
	// ConfigFileName is the name of the config file
	ConfigFileName = "config.json"
	// ConfigDirName is the directory where config is stored
	ConfigDirName = ".gdrv"
	// EnvPrefix is the prefix for environment variables
	EnvPrefix = "GDRV_"
)

// Config holds application configuration
type Config struct {
	// DefaultProfile is the default authentication profile to use
	DefaultProfile string `json:"defaultProfile"`

	// DefaultOutputFormat is the default output format (json, table)
	DefaultOutputFormat types.OutputFormat `json:"defaultOutputFormat"`

	// DefaultFields is the default field mask preset (minimal, standard, full)
	DefaultFields FieldMaskPreset `json:"defaultFields"`

	// CacheTTL is the default path cache TTL in seconds
	CacheTTL int `json:"cacheTTL"`

	// IncludeExportLinks controls whether to include export links by default
	IncludeExportLinks bool `json:"includeExportLinks"`

	// MaxRetries is the maximum number of retries for API calls
	MaxRetries int `json:"maxRetries"`

	// RetryBaseDelay is the base delay for exponential backoff in milliseconds
	RetryBaseDelay int `json:"retryBaseDelay"`

	// RequestTimeout is the default request timeout in seconds
	RequestTimeout int `json:"requestTimeout"`

	// LogLevel sets the logging verbosity (quiet, normal, verbose, debug)
	LogLevel string `json:"logLevel"`

	// ColorOutput enables color output for table format
	ColorOutput bool `json:"colorOutput"`
}

// FieldMaskPreset defines field mask presets
type FieldMaskPreset string

const (
	// FieldMaskMinimal returns only essential fields
	FieldMaskMinimal FieldMaskPreset = "minimal"
	// FieldMaskStandard returns commonly used fields
	FieldMaskStandard FieldMaskPreset = "standard"
	// FieldMaskFull returns all available fields
	FieldMaskFull FieldMaskPreset = "full"
)

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		DefaultProfile:      "default",
		DefaultOutputFormat: types.OutputFormatJSON,
		DefaultFields:       FieldMaskStandard,
		CacheTTL:            300, // 5 minutes
		IncludeExportLinks:  false,
		MaxRetries:          3,
		RetryBaseDelay:      1000, // 1 second
		RequestTimeout:      60,   // 60 seconds
		LogLevel:            "normal",
		ColorOutput:         true,
	}
}

// Load loads configuration with precedence: CLI flags > env vars > config file > defaults
func Load() (*Config, error) {
	// Start with defaults
	cfg := DefaultConfig()

	// Load from config file
	if err := cfg.loadFromFile(); err != nil {
		// Config file not existing is not an error
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}
	}

	// Override with environment variables
	cfg.loadFromEnv()

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

// loadFromFile loads configuration from the config file
func (c *Config) loadFromFile() error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}

// loadFromEnv loads configuration from environment variables
func (c *Config) loadFromEnv() {
	if v := os.Getenv(EnvPrefix + "DEFAULT_PROFILE"); v != "" {
		c.DefaultProfile = v
	}
	if v := os.Getenv(EnvPrefix + "OUTPUT_FORMAT"); v != "" {
		c.DefaultOutputFormat = types.OutputFormat(v)
	}
	if v := os.Getenv(EnvPrefix + "DEFAULT_FIELDS"); v != "" {
		c.DefaultFields = FieldMaskPreset(v)
	}
	if v := os.Getenv(EnvPrefix + "CACHE_TTL"); v != "" {
		if ttl, err := strconv.Atoi(v); err == nil {
			c.CacheTTL = ttl
		}
	}
	if v := os.Getenv(EnvPrefix + "INCLUDE_EXPORT_LINKS"); v != "" {
		c.IncludeExportLinks = parseBool(v)
	}
	if v := os.Getenv(EnvPrefix + "MAX_RETRIES"); v != "" {
		if retries, err := strconv.Atoi(v); err == nil {
			c.MaxRetries = retries
		}
	}
	if v := os.Getenv(EnvPrefix + "RETRY_BASE_DELAY"); v != "" {
		if delay, err := strconv.Atoi(v); err == nil {
			c.RetryBaseDelay = delay
		}
	}
	if v := os.Getenv(EnvPrefix + "REQUEST_TIMEOUT"); v != "" {
		if timeout, err := strconv.Atoi(v); err == nil {
			c.RequestTimeout = timeout
		}
	}
	if v := os.Getenv(EnvPrefix + "LOG_LEVEL"); v != "" {
		c.LogLevel = v
	}
	if v := os.Getenv(EnvPrefix + "COLOR_OUTPUT"); v != "" {
		c.ColorOutput = parseBool(v)
	}
}

// Save saves the configuration to the config file
func (c *Config) Save() error {
	// Validate before saving
	if err := c.Validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// Ensure config directory exists
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Marshal config to JSON
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write to file with restricted permissions
	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	// Validate output format
	if c.DefaultOutputFormat != types.OutputFormatJSON && 
		c.DefaultOutputFormat != types.OutputFormatTable {
		return fmt.Errorf("invalid output format: %s (must be 'json' or 'table')", c.DefaultOutputFormat)
	}

	// Validate field mask preset
	if c.DefaultFields != FieldMaskMinimal && 
		c.DefaultFields != FieldMaskStandard && 
		c.DefaultFields != FieldMaskFull {
		return fmt.Errorf("invalid field mask preset: %s (must be 'minimal', 'standard', or 'full')", c.DefaultFields)
	}

	// Validate cache TTL
	if c.CacheTTL < 0 {
		return fmt.Errorf("cache TTL must be non-negative, got: %d", c.CacheTTL)
	}

	// Validate max retries
	if c.MaxRetries < 0 || c.MaxRetries > 10 {
		return fmt.Errorf("max retries must be between 0 and 10, got: %d", c.MaxRetries)
	}

	// Validate retry base delay
	if c.RetryBaseDelay < 100 || c.RetryBaseDelay > 60000 {
		return fmt.Errorf("retry base delay must be between 100ms and 60000ms, got: %d", c.RetryBaseDelay)
	}

	// Validate request timeout
	if c.RequestTimeout < 1 || c.RequestTimeout > 3600 {
		return fmt.Errorf("request timeout must be between 1 and 3600 seconds, got: %d", c.RequestTimeout)
	}

	// Validate log level
	validLogLevels := []string{"quiet", "normal", "verbose", "debug"}
	isValid := false
	for _, level := range validLogLevels {
		if c.LogLevel == level {
			isValid = true
			break
		}
	}
	if !isValid {
		return fmt.Errorf("invalid log level: %s (must be one of: %s)", c.LogLevel, strings.Join(validLogLevels, ", "))
	}

	return nil
}

// GetCacheTTL returns the cache TTL as a duration
func (c *Config) GetCacheTTL() time.Duration {
	return time.Duration(c.CacheTTL) * time.Second
}

// GetRetryBaseDelay returns the retry base delay as a duration
func (c *Config) GetRetryBaseDelay() time.Duration {
	return time.Duration(c.RetryBaseDelay) * time.Millisecond
}

// GetRequestTimeout returns the request timeout as a duration
func (c *Config) GetRequestTimeout() time.Duration {
	return time.Duration(c.RequestTimeout) * time.Second
}

// GetConfigPath returns the path to the config file
func GetConfigPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, ConfigFileName), nil
}

// GetConfigDir returns the path to the config directory
func GetConfigDir() (string, error) {
	if dir := os.Getenv(EnvPrefix + "CONFIG_DIR"); dir != "" {
		return dir, nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".config", "gdrv"), nil
}

// parseBool parses a boolean value from a string
func parseBool(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	return s == "true" || s == "1" || s == "yes" || s == "on"
}

// GetFieldMask returns the field mask for the given preset
func GetFieldMask(preset FieldMaskPreset, includeExportLinks bool) string {
	var fields []string

	switch preset {
	case FieldMaskMinimal:
		// Minimal: only essential fields for identification
		fields = []string{
			"id",
			"name",
			"mimeType",
		}
	case FieldMaskStandard:
		// Standard: commonly used fields for most operations
		fields = []string{
			"id",
			"name",
			"mimeType",
			"size",
			"md5Checksum",
			"createdTime",
			"modifiedTime",
			"parents",
			"trashed",
			"webViewLink",
			"webContentLink",
			"resourceKey",
			"capabilities(canDownload,canEdit,canShare,canDelete,canTrash,canReadRevisions)",
		}
	case FieldMaskFull:
		// Full: comprehensive field set for detailed information
		fields = []string{
			"id",
			"name",
			"mimeType",
			"size",
			"md5Checksum",
			"createdTime",
			"modifiedTime",
			"modifiedByMe",
			"modifiedByMeTime",
			"parents",
			"properties",
			"appProperties",
			"trashed",
			"trashedTime",
			"starred",
			"shared",
			"sharedWithMeTime",
			"sharingUser",
			"owners",
			"lastModifyingUser",
			"webViewLink",
			"webContentLink",
			"iconLink",
			"thumbnailLink",
			"viewedByMe",
			"viewedByMeTime",
			"resourceKey",
			"capabilities",
			"folderColorRgb",
			"originalFilename",
			"fullFileExtension",
			"fileExtension",
			"md5Checksum",
			"headRevisionId",
			"copyRequiresWriterPermission",
			"writersCanShare",
			"permissions",
			"hasAugmentedPermissions",
			"driveId",
			"shortcutDetails",
			"contentRestrictions",
			"labelInfo",
			"linkShareMetadata",
		}
	}

	// Add exportLinks if requested
	if includeExportLinks {
		fields = append(fields, "exportLinks")
	}

	return "files(" + strings.Join(fields, ",") + ")"
}
