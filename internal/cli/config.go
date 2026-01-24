package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dl-alexandre/gdrive/internal/config"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration management",
	Long:  "Commands for managing gdrive configuration",
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long:  "Display the current configuration settings",
	RunE:  runConfigShow,
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Long:  "Set a configuration value. Use 'config show' to see available keys",
	Args:  cobra.ExactArgs(2),
	RunE:  runConfigSet,
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration to defaults",
	Long:  "Reset all configuration settings to their default values",
	RunE:  runConfigReset,
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configResetCmd)
}

func runConfigShow(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	cfg, err := config.Load()
	if err != nil {
		return out.WriteError("config.show", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("config.show", cfg)
}

func runConfigSet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	key := args[0]
	value := args[1]

	cfg, err := config.Load()
	if err != nil {
		return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	// Set the value based on key
	switch strings.ToLower(key) {
	case "defaultprofile":
		cfg.DefaultProfile = value
	case "defaultoutputformat":
		if value != string(types.OutputFormatJSON) && value != string(types.OutputFormatTable) {
			return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				"Invalid output format. Must be 'json' or 'table'").Build())
		}
		cfg.DefaultOutputFormat = types.OutputFormat(value)
	case "defaultfields":
		preset := config.FieldMaskPreset(value)
		if preset != config.FieldMaskMinimal && preset != config.FieldMaskStandard && preset != config.FieldMaskFull {
			return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				"Invalid field mask preset. Must be 'minimal', 'standard', or 'full'").Build())
		}
		cfg.DefaultFields = preset
	case "cachettl":
		ttl, err := strconv.Atoi(value)
		if err != nil || ttl < 0 {
			return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				"Cache TTL must be a non-negative integer").Build())
		}
		cfg.CacheTTL = ttl
	case "includeexportlinks":
		cfg.IncludeExportLinks = parseBool(value)
	case "maxretries":
		retries, err := strconv.Atoi(value)
		if err != nil || retries < 0 || retries > 10 {
			return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				"Max retries must be between 0 and 10").Build())
		}
		cfg.MaxRetries = retries
	case "retrybasedelay":
		delay, err := strconv.Atoi(value)
		if err != nil || delay < 100 || delay > 60000 {
			return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				"Retry base delay must be between 100 and 60000 ms").Build())
		}
		cfg.RetryBaseDelay = delay
	case "requesttimeout":
		timeout, err := strconv.Atoi(value)
		if err != nil || timeout < 1 || timeout > 3600 {
			return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				"Request timeout must be between 1 and 3600 seconds").Build())
		}
		cfg.RequestTimeout = timeout
	case "loglevel":
		validLevels := []string{"quiet", "normal", "verbose", "debug"}
		valid := false
		for _, level := range validLevels {
			if value == level {
				valid = true
				break
			}
		}
		if !valid {
			return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				fmt.Sprintf("Invalid log level. Must be one of: %s", strings.Join(validLevels, ", "))).Build())
		}
		cfg.LogLevel = value
	case "coloroutput":
		cfg.ColorOutput = parseBool(value)
	default:
		return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeInvalidArgument,
			fmt.Sprintf("Unknown configuration key: %s", key)).Build())
	}

	// Save the configuration
	if err := cfg.Save(); err != nil {
		return out.WriteError("config.set", utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to save configuration: %v", err)).Build())
	}

	out.Log("Configuration updated: %s = %s", key, value)
	return out.WriteSuccess("config.set", map[string]interface{}{
		"key":   key,
		"value": value,
	})
}

func runConfigReset(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)

	cfg := config.DefaultConfig()
	if err := cfg.Save(); err != nil {
		return out.WriteError("config.reset", utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to reset configuration: %v", err)).Build())
	}

	out.Log("Configuration reset to defaults")
	return out.WriteSuccess("config.reset", cfg)
}

// parseBool parses a boolean value from a string
func parseBool(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	return s == "true" || s == "1" || s == "yes" || s == "on"
}
