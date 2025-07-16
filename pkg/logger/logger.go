package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// LogLevel represents the severity level of a log message
type LogLevel int

const (
	// DEBUG level for detailed debugging information
	DEBUG LogLevel = iota
	// INFO level for general operational information
	INFO
	// WARN level for warning messages
	WARN
	// ERROR level for error messages
	ERROR
)

var (
	// Default logger instance
	defaultLogger *Logger
	// Log level names for string conversion
	levelNames = map[LogLevel]string{
		DEBUG: "DEBUG",
		INFO:  "INFO",
		WARN:  "WARN",
		ERROR: "ERROR",
	}
	// Log level values for string parsing
	levelValues = map[string]LogLevel{
		"DEBUG": DEBUG,
		"INFO":  INFO,
		"WARN":  WARN,
		"ERROR": ERROR,
	}
)

// Logger is a simple leveled logger
type Logger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	level       LogLevel
}

// NewLogger creates a new logger with the specified level
func NewLogger(out io.Writer, level LogLevel) *Logger {
	return &Logger{
		debugLogger: log.New(out, colorCyan+"DEBUG: "+colorReset, log.Ldate|log.Ltime),
		infoLogger:  log.New(out, colorGreen+"INFO: "+colorReset, log.Ldate|log.Ltime),
		warnLogger:  log.New(out, colorYellow+"WARN: "+colorReset, log.Ldate|log.Ltime),
		errorLogger: log.New(out, colorRed+"ERROR: "+colorReset, log.Ldate|log.Ltime),
		level:       level,
	}
}

// ParseLevel parses a log level string into a LogLevel
func ParseLevel(level string) (LogLevel, error) {
	level = strings.ToUpper(level)
	if l, ok := levelValues[level]; ok {
		return l, nil
	}
	return INFO, fmt.Errorf("invalid log level: %s", level)
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.level <= DEBUG {
		l.debugLogger.Printf(format, v...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	if l.level <= INFO {
		l.infoLogger.Printf(format, v...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	if l.level <= WARN {
		l.warnLogger.Printf(format, v...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.errorLogger.Printf(format, v...)
	}
}

// Initialize the default logger
func init() {
	defaultLogger = NewLogger(os.Stdout, INFO)
}

// SetDefaultLevel sets the log level for the default logger
func SetDefaultLevel(level LogLevel) {
	defaultLogger.SetLevel(level)
}

// Debug logs a debug message using the default logger
func Debug(format string, v ...interface{}) {
	defaultLogger.Debug(format, v...)
}

// Info logs an info message using the default logger
func Info(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

// Warn logs a warning message using the default logger
func Warn(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

// Error logs an error message using the default logger
func Error(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}
