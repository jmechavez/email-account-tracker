package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  *zap.Logger
	once sync.Once
)

func Initialize() {
	once.Do(func() {
		// Get environment - default to development if not specified
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "development"
		}

		var config zap.Config
		if env == "production" {
			config = zap.NewProductionConfig()
		} else {
			// More readable logging for development
			config = zap.NewDevelopmentConfig()
		}

		// Configure encoder
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		config.EncoderConfig = encoderConfig

		var err error
		log, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
	})
}

// GetLogger returns the singleton logger instance
func GetLogger() *zap.Logger {
	if log == nil {
		Initialize()
	}
	return log
}

// Info logs a message at info level
func Info(message string, fields ...zap.Field) {
	if log == nil {
		Initialize()
	}
	log.Info(message, fields...)
}

// Error logs a message at error level
func Error(message string, fields ...zap.Field) {
	if log == nil {
		Initialize()
	}
	log.Error(message, fields...)
}

// Warn logs a message at warn level
func Warn(message string, fields ...zap.Field) {
	if log == nil {
		Initialize()
	}
	log.Warn(message, fields...)
}

// Debug logs a message at debug level
func Debug(message string, fields ...zap.Field) {
	if log == nil {
		Initialize()
	}
	log.Debug(message, fields...)
}

// Fatal logs a message at fatal level then calls os.Exit(1)
func Fatal(message string, fields ...zap.Field) {
	if log == nil {
		Initialize()
	}
	log.Fatal(message, fields...)
}

// WithFields returns a new entry with the specified fields
func WithFields(fields ...zap.Field) *zap.Logger {
	if log == nil {
		Initialize()
	}
	return log.With(fields...)
}

// Sync flushes any buffered log entries
func Sync() error {
	if log == nil {
		return nil
	}
	return log.Sync()
}
