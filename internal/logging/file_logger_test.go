package logging

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFileLogger_Creation(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := FileLoggerConfig{
		FilePath:      logPath,
		Level:         INFO,
		MaxFileSize:   1024,
		RotateEnabled: true,
	}

	logger, err := NewFileLogger(config)
	if err != nil {
		t.Fatalf("NewFileLogger() error = %v", err)
	}
	t.Cleanup(func() {
		if closeErr := logger.Close(); closeErr != nil {
			t.Fatalf("Failed to close logger: %v", closeErr)
		}
	})

	// Verify file was created
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}
}

func TestFileLogger_Logging(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := FileLoggerConfig{
		FilePath:      logPath,
		Level:         DEBUG,
		MaxFileSize:   0, // No rotation
		RotateEnabled: false,
	}

	logger, err := NewFileLogger(config)
	if err != nil {
		t.Fatalf("NewFileLogger() error = %v", err)
	}
	t.Cleanup(func() {
		if closeErr := logger.Close(); closeErr != nil {
			t.Fatalf("Failed to close logger: %v", closeErr)
		}
	})

	// Log various levels
	logger.Debug("debug message", F("key1", "value1"))
	logger.Info("info message", F("key2", 123))
	logger.Warn("warn message")
	logger.Error("error message", F("key3", true))

	if err := logger.Close(); err != nil {
		t.Fatalf("Failed to close logger: %v", err)
	}

	// Read log file
	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Verify log entries
	lines := splitLogLines(data)
	if len(lines) != 4 {
		t.Errorf("Expected 4 log entries, got %d", len(lines))
	}

	// Parse first entry
	var entry LogEntry
	if err := json.Unmarshal([]byte(lines[0]), &entry); err != nil {
		t.Fatalf("Failed to parse log entry: %v", err)
	}

	if entry.Level != "DEBUG" {
		t.Errorf("Entry.Level = %v, want DEBUG", entry.Level)
	}
	if entry.Message != "debug message" {
		t.Errorf("Entry.Message = %v, want 'debug message'", entry.Message)
	}
	if entry.Fields["key1"] != "value1" {
		t.Errorf("Entry.Fields[key1] = %v, want 'value1'", entry.Fields["key1"])
	}
}

func TestFileLogger_LevelFiltering(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := FileLoggerConfig{
		FilePath:      logPath,
		Level:         WARN, // Only WARN and ERROR
		MaxFileSize:   0,
		RotateEnabled: false,
	}

	logger, err := NewFileLogger(config)
	if err != nil {
		t.Fatalf("NewFileLogger() error = %v", err)
	}

	logger.Debug("debug message") // Should be filtered
	logger.Info("info message")   // Should be filtered
	logger.Warn("warn message")   // Should be logged
	logger.Error("error message") // Should be logged

	logger.Close()

	// Read log file
	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	lines := splitLogLines(data)
	if len(lines) != 2 {
		t.Errorf("Expected 2 log entries, got %d", len(lines))
	}
}

func TestFileLogger_WithTraceID(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := FileLoggerConfig{
		FilePath:      logPath,
		Level:         INFO,
		MaxFileSize:   0,
		RotateEnabled: false,
	}

	logger, err := NewFileLogger(config)
	if err != nil {
		t.Fatalf("NewFileLogger() error = %v", err)
	}

	traceID := "trace-123-456"
	tracedLogger := logger.WithTraceID(traceID)
	tracedLogger.Info("test message")

	logger.Close()

	// Read and verify trace ID
	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	var entry LogEntry
	lines := splitLogLines(data)
	if err := json.Unmarshal([]byte(lines[0]), &entry); err != nil {
		t.Fatalf("Failed to parse log entry: %v", err)
	}

	if entry.TraceID != traceID {
		t.Errorf("Entry.TraceID = %v, want %v", entry.TraceID, traceID)
	}
}

func TestFileLogger_WithContext(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := FileLoggerConfig{
		FilePath:      logPath,
		Level:         INFO,
		MaxFileSize:   0,
		RotateEnabled: false,
	}

	logger, err := NewFileLogger(config)
	if err != nil {
		t.Fatalf("NewFileLogger() error = %v", err)
	}

	ctx := context.Background()
	traceID := "ctx-trace-789"
	ctx = ContextWithTraceID(ctx, traceID)

	tracedLogger := logger.WithContext(ctx)
	tracedLogger.Info("test message")

	logger.Close()

	// Read and verify trace ID
	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	var entry LogEntry
	lines := splitLogLines(data)
	if err := json.Unmarshal([]byte(lines[0]), &entry); err != nil {
		t.Fatalf("Failed to parse log entry: %v", err)
	}

	if entry.TraceID != traceID {
		t.Errorf("Entry.TraceID = %v, want %v", entry.TraceID, traceID)
	}
}

func TestFileLogger_Rotation(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := FileLoggerConfig{
		FilePath:      logPath,
		Level:         INFO,
		MaxFileSize:   100, // Very small for testing
		RotateEnabled: true,
	}

	logger, err := NewFileLogger(config)
	if err != nil {
		t.Fatalf("NewFileLogger() error = %v", err)
	}

	// Write enough data to trigger rotation
	for i := 0; i < 20; i++ {
		logger.Info("This is a test message that should trigger rotation")
		time.Sleep(1 * time.Millisecond)
	}

	logger.Close()

	// Check for rotated files
	files, err := filepath.Glob(filepath.Join(tempDir, "test.log*"))
	if err != nil {
		t.Fatalf("Failed to glob log files: %v", err)
	}

	// Should have original file plus at least one rotated file
	if len(files) < 2 {
		t.Errorf("Expected at least 2 log files (original + rotated), got %d", len(files))
	}
}

func TestFileLogger_SetLevel(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := FileLoggerConfig{
		FilePath:      logPath,
		Level:         DEBUG,
		MaxFileSize:   0,
		RotateEnabled: false,
	}

	logger, err := NewFileLogger(config)
	if err != nil {
		t.Fatalf("NewFileLogger() error = %v", err)
	}

	logger.Debug("debug 1") // Should be logged

	// Change level to ERROR
	logger.SetLevel(ERROR)

	logger.Debug("debug 2") // Should be filtered
	logger.Info("info 2")   // Should be filtered
	logger.Error("error 1") // Should be logged

	logger.Close()

	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	lines := splitLogLines(data)
	if len(lines) != 2 {
		t.Errorf("Expected 2 log entries, got %d", len(lines))
	}
}

// Helper function to split log lines
func splitLogLines(data []byte) []string {
	var lines []string
	current := ""
	for _, b := range data {
		if b == '\n' {
			if current != "" {
				lines = append(lines, current)
				current = ""
			}
		} else {
			current += string(b)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}
