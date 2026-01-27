package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// FileLogger implements Logger interface for file-based logging
type FileLogger struct {
	mu            sync.Mutex
	file          *os.File
	filePath      string
	level         LogLevel
	traceID       string
	maxFileSize   int64
	currentSize   int64
	rotateEnabled bool
}

// FileLoggerConfig contains configuration for file logger
type FileLoggerConfig struct {
	FilePath      string
	Level         LogLevel
	MaxFileSize   int64 // in bytes, 0 means no rotation
	RotateEnabled bool
}

// NewFileLogger creates a new file logger
func NewFileLogger(config FileLoggerConfig) (*FileLogger, error) {
	// Create directory if it doesn't exist
	dir := filepath.Dir(config.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Open or create log file
	file, err := os.OpenFile(config.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Get current file size
	info, err := file.Stat()
	if err != nil {
		if closeErr := file.Close(); closeErr != nil {
			return nil, fmt.Errorf("failed to close log file after stat error: %w", closeErr)
		}
		return nil, fmt.Errorf("failed to stat log file: %w", err)
	}

	return &FileLogger{
		file:          file,
		filePath:      config.FilePath,
		level:         config.Level,
		maxFileSize:   config.MaxFileSize,
		currentSize:   info.Size(),
		rotateEnabled: config.RotateEnabled && config.MaxFileSize > 0,
	}, nil
}

// log writes a log entry to the file
func (l *FileLogger) log(level LogLevel, msg string, fields ...Field) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Check if rotation is needed
	if l.rotateEnabled && l.currentSize >= l.maxFileSize {
		if err := l.rotate(); err != nil {
			// If rotation fails, log to stderr and continue
			fmt.Fprintf(os.Stderr, "Failed to rotate log file: %v\n", err)
		}
	}

	entry := LogEntry{
		Timestamp: time.Now().UTC(),
		Level:     level.String(),
		Message:   msg,
		TraceID:   l.traceID,
		Fields:    make(map[string]interface{}),
	}

	// Add fields to entry
	for _, field := range fields {
		entry.Fields[field.Key] = field.Value
	}

	// Marshal to JSON
	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal log entry: %v\n", err)
		return
	}

	// Write to file
	data = append(data, '\n')
	n, err := l.file.Write(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write log entry: %v\n", err)
		return
	}

	l.currentSize += int64(n)
}

// rotate rotates the log file
func (l *FileLogger) rotate() error {
	// Close current file
	if err := l.file.Close(); err != nil {
		return fmt.Errorf("failed to close log file: %w", err)
	}

	// Rename current file with timestamp
	timestamp := time.Now().UTC().Format("20060102-150405")
	rotatedPath := fmt.Sprintf("%s.%s", l.filePath, timestamp)
	if err := os.Rename(l.filePath, rotatedPath); err != nil {
		// Try to reopen the original file
		file, _ := os.OpenFile(l.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		l.file = file
		return fmt.Errorf("failed to rename log file: %w", err)
	}

	// Open new file
	file, err := os.OpenFile(l.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to create new log file: %w", err)
	}

	l.file = file
	l.currentSize = 0

	return nil
}

// Debug logs a debug-level message
func (l *FileLogger) Debug(msg string, fields ...Field) {
	l.log(DEBUG, msg, fields...)
}

// Info logs an info-level message
func (l *FileLogger) Info(msg string, fields ...Field) {
	l.log(INFO, msg, fields...)
}

// Warn logs a warning-level message
func (l *FileLogger) Warn(msg string, fields ...Field) {
	l.log(WARN, msg, fields...)
}

// Error logs an error-level message
func (l *FileLogger) Error(msg string, fields ...Field) {
	l.log(ERROR, msg, fields...)
}

// WithTraceID returns a new logger with the trace ID set
func (l *FileLogger) WithTraceID(traceID string) Logger {
	return &FileLogger{
		file:          l.file,
		filePath:      l.filePath,
		level:         l.level,
		traceID:       traceID,
		maxFileSize:   l.maxFileSize,
		currentSize:   l.currentSize,
		rotateEnabled: l.rotateEnabled,
	}
}

// WithContext returns a new logger that extracts trace ID from context
func (l *FileLogger) WithContext(ctx context.Context) Logger {
	traceID := TraceIDFromContext(ctx)
	if traceID == "" {
		return l
	}
	return l.WithTraceID(traceID)
}

// SetLevel sets the minimum log level
func (l *FileLogger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// Close closes the log file
func (l *FileLogger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
