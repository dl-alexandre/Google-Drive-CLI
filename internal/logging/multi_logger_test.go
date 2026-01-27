package logging

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestMultiLogger_Creation(t *testing.T) {
	var buf bytes.Buffer
	console := NewConsoleLogger(ConsoleLoggerConfig{
		Writer: &buf,
		Level:  INFO,
	})

	multi := NewMultiLogger(console)
	if multi == nil {
		t.Fatal("NewMultiLogger() returned nil")
	}
}

func TestMultiLogger_LogsToAll(t *testing.T) {
	var buf1, buf2 bytes.Buffer

	logger1 := NewConsoleLogger(ConsoleLoggerConfig{
		Writer:           &buf1,
		Level:            INFO,
		ColorEnabled:     false,
		TimestampEnabled: false,
	})

	logger2 := NewConsoleLogger(ConsoleLoggerConfig{
		Writer:           &buf2,
		Level:            INFO,
		ColorEnabled:     false,
		TimestampEnabled: false,
	})

	multi := NewMultiLogger(logger1, logger2)
	multi.Info("test message")

	output1 := buf1.String()
	output2 := buf2.String()

	if output1 == "" {
		t.Error("First logger didn't receive message")
	}
	if output2 == "" {
		t.Error("Second logger didn't receive message")
	}
	if output1 != output2 {
		t.Errorf("Loggers produced different output:\n%s\n%s", output1, output2)
	}
}

func TestMultiLogger_AllLevels(t *testing.T) {
	var buf bytes.Buffer
	console := NewConsoleLogger(ConsoleLoggerConfig{
		Writer:           &buf,
		Level:            DEBUG,
		ColorEnabled:     false,
		TimestampEnabled: false,
	})

	multi := NewMultiLogger(console)

	multi.Debug("debug message")
	multi.Info("info message")
	multi.Warn("warn message")
	multi.Error("error message")

	output := buf.String()
	if output == "" {
		t.Error("MultiLogger didn't log anything")
	}
}

func TestMultiLogger_WithTraceID(t *testing.T) {
	var buf bytes.Buffer
	console := NewConsoleLogger(ConsoleLoggerConfig{
		Writer:           &buf,
		Level:            INFO,
		ColorEnabled:     false,
		TimestampEnabled: false,
	})

	multi := NewMultiLogger(console)
	traceID := "trace-123"

	traced := multi.WithTraceID(traceID)
	traced.Info("test message")

	output := buf.String()
	if output == "" {
		t.Error("Traced logger didn't log anything")
	}
}

func TestMultiLogger_WithContext(t *testing.T) {
	var buf bytes.Buffer
	console := NewConsoleLogger(ConsoleLoggerConfig{
		Writer:           &buf,
		Level:            INFO,
		ColorEnabled:     false,
		TimestampEnabled: false,
	})

	multi := NewMultiLogger(console)
	ctx := context.Background()
	traceID := "ctx-trace"
	ctx = ContextWithTraceID(ctx, traceID)

	traced := multi.WithContext(ctx)
	traced.Info("test message")

	output := buf.String()
	if output == "" {
		t.Error("Context logger didn't log anything")
	}
}

func TestMultiLogger_SetLevel(t *testing.T) {
	var buf bytes.Buffer
	console := NewConsoleLogger(ConsoleLoggerConfig{
		Writer:           &buf,
		Level:            DEBUG,
		ColorEnabled:     false,
		TimestampEnabled: false,
	})

	multi := NewMultiLogger(console)

	multi.Debug("debug 1") // Should log

	multi.SetLevel(ERROR)

	multi.Debug("debug 2") // Should not log
	multi.Info("info 2")   // Should not log
	multi.Error("error 1") // Should log

	output := buf.String()
	if output == "" {
		t.Error("MultiLogger didn't log anything")
	}
}

func TestMultiLogger_Close(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	fileLogger, err := NewFileLogger(FileLoggerConfig{
		FilePath: logPath,
		Level:    INFO,
	})
	if err != nil {
		t.Fatalf("Failed to create file logger: %v", err)
	}

	multi := NewMultiLogger(fileLogger)

	if err := multi.Close(); err != nil {
		t.Errorf("Close() returned error: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}
}

func TestMultiLogger_FileAndConsole(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "test.log")

	var buf bytes.Buffer

	fileLogger, err := NewFileLogger(FileLoggerConfig{
		FilePath: logPath,
		Level:    INFO,
	})
	if err != nil {
		t.Fatalf("Failed to create file logger: %v", err)
	}

	consoleLogger := NewConsoleLogger(ConsoleLoggerConfig{
		Writer:           &buf,
		Level:            INFO,
		ColorEnabled:     false,
		TimestampEnabled: false,
	})

	multi := NewMultiLogger(fileLogger, consoleLogger)

	multi.Info("test message", F("key", "value"))

	if err := multi.Close(); err != nil {
		t.Fatalf("Failed to close multi logger: %v", err)
	}

	// Verify console output
	consoleOutput := buf.String()
	if consoleOutput == "" {
		t.Error("Console didn't receive message")
	}

	// Verify file output
	fileData, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}
	if len(fileData) == 0 {
		t.Error("Log file is empty")
	}
}
