package logging

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDefaultLogConfig(t *testing.T) {
	config := DefaultLogConfig()

	if config.Level != INFO {
		t.Errorf("Expected Level=INFO, got %v", config.Level)
	}
	if !config.EnableConsole {
		t.Error("Expected EnableConsole=true")
	}
	if !config.RedactSensitive {
		t.Error("Expected RedactSensitive=true")
	}
	if config.MaxFileSize != 100*1024*1024 {
		t.Errorf("Expected MaxFileSize=104857600, got %v", config.MaxFileSize)
	}
}

func TestNewLogger_ConsoleOnly(t *testing.T) {
	config := LogConfig{
		Level:         INFO,
		EnableConsole: true,
		OutputFile:    "",
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("NewLogger() error = %v", err)
	}
	t.Cleanup(func() {
		logger.Close()
	})

	if logger == nil {
		t.Fatal("Logger is nil")
	}

	// Should be a ConsoleLogger
	if _, ok := logger.(*ConsoleLogger); !ok {
		t.Errorf("Expected ConsoleLogger, got %T", logger)
	}
}

func TestNewLogger_FileOnly(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := LogConfig{
		Level:         INFO,
		EnableConsole: false,
		OutputFile:    logPath,
		MaxFileSize:   1024,
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("NewLogger() error = %v", err)
	}
	t.Cleanup(func() {
		logger.Close()
	})

	if logger == nil {
		t.Fatal("Logger is nil")
	}

	// Should be a FileLogger
	if _, ok := logger.(*FileLogger); !ok {
		t.Errorf("Expected FileLogger, got %T", logger)
	}

	// Verify file was created
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}
}

func TestNewLogger_Both(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := LogConfig{
		Level:         INFO,
		EnableConsole: true,
		OutputFile:    logPath,
		MaxFileSize:   1024,
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("NewLogger() error = %v", err)
	}
	t.Cleanup(func() {
		logger.Close()
	})

	if logger == nil {
		t.Fatal("Logger is nil")
	}

	// Should be a MultiLogger
	if _, ok := logger.(*MultiLogger); !ok {
		t.Errorf("Expected MultiLogger, got %T", logger)
	}
}

func TestNewLogger_NoOp(t *testing.T) {
	config := LogConfig{
		Level:         INFO,
		EnableConsole: false,
		OutputFile:    "",
	}

	logger, err := NewLogger(config)
	if err != nil {
		t.Fatalf("NewLogger() error = %v", err)
	}
	t.Cleanup(func() {
		logger.Close()
	})

	if logger == nil {
		t.Fatal("Logger is nil")
	}

	// Should be a NoOpLogger
	if _, ok := logger.(*NoOpLogger); !ok {
		t.Errorf("Expected NoOpLogger, got %T", logger)
	}
}

func TestNewLogger_InvalidPath(t *testing.T) {
	var invalidPath string
	if runtime.GOOS == "windows" {
		invalidPath = `Z:\nonexistent\path\that\does\not\exist\test.log`
	} else {
		invalidPath = "/invalid/path/that/does/not/exist/test.log"
	}

	config := LogConfig{
		Level:         INFO,
		EnableConsole: false,
		OutputFile:    invalidPath,
	}

	_, err := NewLogger(config)
	if err == nil {
		t.Error("Expected error for invalid path, got nil")
	}
}

func TestNewDebugLoggerWithTransport(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	config := LogConfig{
		Level:         DEBUG,
		EnableConsole: false,
		OutputFile:    logPath,
		EnableDebug:   true,
	}

	logger, transport, err := NewDebugLoggerWithTransport(config)
	if err != nil {
		t.Fatalf("NewDebugLoggerWithTransport() error = %v", err)
	}
	t.Cleanup(func() {
		logger.Close()
	})

	if logger == nil {
		t.Fatal("Logger is nil")
	}

	if transport == nil {
		t.Fatal("DebugTransport is nil")
	}
}

func TestNewDebugLoggerWithTransport_NoDebug(t *testing.T) {
	config := LogConfig{
		Level:         INFO,
		EnableConsole: true,
		EnableDebug:   false,
	}

	logger, transport, err := NewDebugLoggerWithTransport(config)
	if err != nil {
		t.Fatalf("NewDebugLoggerWithTransport() error = %v", err)
	}
	t.Cleanup(func() {
		logger.Close()
	})

	if logger == nil {
		t.Fatal("Logger is nil")
	}

	if transport != nil {
		t.Error("Expected nil DebugTransport when EnableDebug=false")
	}
}
