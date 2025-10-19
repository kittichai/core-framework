package logger

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type contextKey string

const LoggerKey contextKey = "logger"

// ECSLogger wraps zerolog.Logger with ECS fields
type ECSLogger struct {
	logger zerolog.Logger
}

// NewECSLogger creates a new ECS logger
func NewECSLogger() *ECSLogger {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("ecs.version", "1.12.0").
		Logger()

	return &ECSLogger{logger: logger}
}

// WithContext adds logger to context with request ID
func (l *ECSLogger) WithContext(ctx context.Context) context.Context {
	requestID := uuid.New().String()
	logger := l.logger.With().Str("request.id", requestID).Logger()
	return context.WithValue(ctx, LoggerKey, &logger)
}

// FromContext retrieves logger from context
func (l *ECSLogger) FromContext(ctx context.Context) *zerolog.Logger {
	if logger, ok := ctx.Value(LoggerKey).(*zerolog.Logger); ok {
		return logger
	}
	return &zerolog.Logger{}
}

// Log methods for different levels
func (l *ECSLogger) Info(ctx context.Context) *zerolog.Event {
	return l.FromContext(ctx).Info()
}

func (l *ECSLogger) Error(ctx context.Context) *zerolog.Event {
	return l.FromContext(ctx).Error()
}

func (l *ECSLogger) Debug(ctx context.Context) *zerolog.Event {
	return l.FromContext(ctx).Debug()
}
