package logging

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorGray   = "\033[90m"
)

// ConsoleLogger implements Logger interface for console output
type ConsoleLogger struct {
	mu               sync.Mutex
	writer           io.Writer
	level            LogLevel
	traceID          string
	colorEnabled     bool
	timestampEnabled bool
	redactSensitive  bool
}

// ConsoleLoggerConfig contains configuration for console logger
type ConsoleLoggerConfig struct {
	Writer           io.Writer
	Level            LogLevel
	ColorEnabled     bool
	TimestampEnabled bool
	RedactSensitive  bool
}

// NewConsoleLogger creates a new console logger
func NewConsoleLogger(config ConsoleLoggerConfig) *ConsoleLogger {
	if config.Writer == nil {
		config.Writer = os.Stderr
	}

	return &ConsoleLogger{
		writer:           config.Writer,
		level:            config.Level,
		colorEnabled:     config.ColorEnabled,
		timestampEnabled: config.TimestampEnabled,
		redactSensitive:  config.RedactSensitive,
	}
}

// Patterns for sensitive data redaction
var (
	// Bearer tokens
	bearerTokenPattern = regexp.MustCompile(`Bearer\s+[A-Za-z0-9\-._~+/]+=*`)
	// OAuth tokens
	oauthTokenPattern = regexp.MustCompile(`(access_token|refresh_token|id_token)["']?\s*[:=]\s*["']?[A-Za-z0-9\-._~+/]+=*`)
	// API keys
	apiKeyPattern = regexp.MustCompile(`(?i)(api[_-]?key|apikey)["']?\s*[:=]\s*["']?[A-Za-z0-9\-._~+/]+=*`)
	// Authorization headers
	authHeaderPattern = regexp.MustCompile(`(?i)authorization["']?\s*[:=]\s*["']?[^\s"']+`)
)

// redactSensitiveData redacts sensitive information from log messages
func redactSensitiveData(s string) string {
	s = bearerTokenPattern.ReplaceAllString(s, "Bearer [REDACTED]")
	s = oauthTokenPattern.ReplaceAllString(s, "$1=[REDACTED]")
	s = apiKeyPattern.ReplaceAllString(s, "$1=[REDACTED]")
	s = authHeaderPattern.ReplaceAllString(s, "Authorization: [REDACTED]")
	return s
}

// formatMessage formats a log message with colors and fields
func (l *ConsoleLogger) formatMessage(level LogLevel, msg string, fields ...Field) string {
	var sb strings.Builder

	// Add timestamp if enabled
	if l.timestampEnabled {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		if l.colorEnabled {
			sb.WriteString(colorGray)
		}
		sb.WriteString(timestamp)
		sb.WriteString(" ")
		if l.colorEnabled {
			sb.WriteString(colorReset)
		}
	}

	// Add log level with color
	levelStr := level.String()
	if l.colorEnabled {
		switch level {
		case DEBUG:
			sb.WriteString(colorBlue)
		case INFO:
			sb.WriteString(colorReset)
		case WARN:
			sb.WriteString(colorYellow)
		case ERROR:
			sb.WriteString(colorRed)
		}
	}
	sb.WriteString(fmt.Sprintf("%-5s", levelStr))
	if l.colorEnabled {
		sb.WriteString(colorReset)
	}
	sb.WriteString(" ")

	// Add trace ID if present
	if l.traceID != "" {
		if l.colorEnabled {
			sb.WriteString(colorGray)
		}
		sb.WriteString(fmt.Sprintf("[%s] ", l.traceID[:8])) // Show first 8 chars
		if l.colorEnabled {
			sb.WriteString(colorReset)
		}
	}

	// Add message
	if l.redactSensitive {
		msg = redactSensitiveData(msg)
	}
	sb.WriteString(msg)

	// Add fields
	if len(fields) > 0 {
		sb.WriteString(" ")
		for i, field := range fields {
			if i > 0 {
				sb.WriteString(", ")
			}
			value := fmt.Sprintf("%v", field.Value)
			if l.redactSensitive {
				value = redactSensitiveData(value)
			}
			sb.WriteString(fmt.Sprintf("%s=%s", field.Key, value))
		}
	}

	return sb.String()
}

// log writes a log message to the console
func (l *ConsoleLogger) log(level LogLevel, msg string, fields ...Field) {
	if level < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	formatted := l.formatMessage(level, msg, fields...)
	if _, err := fmt.Fprintln(l.writer, formatted); err != nil {
		return
	}
}

// Debug logs a debug-level message
func (l *ConsoleLogger) Debug(msg string, fields ...Field) {
	l.log(DEBUG, msg, fields...)
}

// Info logs an info-level message
func (l *ConsoleLogger) Info(msg string, fields ...Field) {
	l.log(INFO, msg, fields...)
}

// Warn logs a warning-level message
func (l *ConsoleLogger) Warn(msg string, fields ...Field) {
	l.log(WARN, msg, fields...)
}

// Error logs an error-level message
func (l *ConsoleLogger) Error(msg string, fields ...Field) {
	l.log(ERROR, msg, fields...)
}

// WithTraceID returns a new logger with the trace ID set
func (l *ConsoleLogger) WithTraceID(traceID string) Logger {
	return &ConsoleLogger{
		writer:           l.writer,
		level:            l.level,
		traceID:          traceID,
		colorEnabled:     l.colorEnabled,
		timestampEnabled: l.timestampEnabled,
		redactSensitive:  l.redactSensitive,
	}
}

// WithContext returns a new logger that extracts trace ID from context
func (l *ConsoleLogger) WithContext(ctx context.Context) Logger {
	traceID := TraceIDFromContext(ctx)
	if traceID == "" {
		return l
	}
	return l.WithTraceID(traceID)
}

// SetLevel sets the minimum log level
func (l *ConsoleLogger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// Close closes the logger (no-op for console logger)
func (l *ConsoleLogger) Close() error {
	return nil
}
