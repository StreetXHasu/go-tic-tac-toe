package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Logger structure for logging
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	file        *os.File
}

// NewLogger creates a new logger
func NewLogger(serviceName string) (*Logger, error) {
	// Create logs directory if it doesn't exist
	logsDir := "/app/logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Create log file
	logFile := filepath.Join(logsDir, fmt.Sprintf("%s.log", serviceName))
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Create loggers with different levels
	infoLogger := log.New(file, fmt.Sprintf("[INFO][%s] ", serviceName), log.Ldate|log.Ltime)
	errorLogger := log.New(file, fmt.Sprintf("[ERROR][%s] ", serviceName), log.Ldate|log.Ltime)
	debugLogger := log.New(file, fmt.Sprintf("[DEBUG][%s] ", serviceName), log.Ldate|log.Ltime)

	return &Logger{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
		debugLogger: debugLogger,
		file:        file,
	}, nil
}

// Info logs an informational message
func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	l.debugLogger.Printf(format, v...)
}

// Close closes the logger file
func (l *Logger) Close() error {
	return l.file.Close()
}

// LogRequest logs an HTTP request
func (l *Logger) LogRequest(method, path, remoteAddr string, statusCode int, duration time.Duration) {
	l.Info("%s %s from %s - %d - %v", method, path, remoteAddr, statusCode, duration)
}

// LogError logs an error with context
func (l *Logger) LogError(operation string, err error) {
	l.Error("Operation: %s, Error: %v", operation, err)
}

// LogDatabaseOperation logs database operations
func (l *Logger) LogDatabaseOperation(operation string, table string, duration time.Duration) {
	l.Info("DB %s on %s - %v", operation, table, duration)
}

// LogWebSocketEvent logs WebSocket events
func (l *Logger) LogWebSocketEvent(event, clientID, gameID string) {
	l.Info("WS %s - Client: %s, Game: %s", event, clientID, gameID)
} 