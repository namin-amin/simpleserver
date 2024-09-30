package logger

import (
	"context"
	"log/slog"
	"os"
)

type Logger interface {
	Info(msg string, v ...any)
	Error(msg string, v ...any)
	Debug(msg string, v ...any)
	Warn(msg string, v ...any)
	Fatal(msg string)
	InfoWithContext(ctx context.Context, msg string, v ...any)
	ErrorWithContext(ctx context.Context, msg string, v ...any)
	DebugWithContext(ctx context.Context, msg string, v ...any)
	WarnWithContext(ctx context.Context, msg string, v ...any)
	Log(logLevel LogLevel, msg string, v ...any)
}

type LogHandler func(loglevel LogLevel, msg string, v ...any)

type Config struct {
	Handler slog.Handler
}

type LogLevel int

const (
	INFO  LogLevel = iota
	DEBUG LogLevel = iota
	WARN  LogLevel = iota
	ERROR LogLevel = iota
	FATAL LogLevel = iota
)

type logger struct {
	logger *slog.Logger
}

func (l logger) Info(msg string, v ...any) {
	l.logger.Info(msg, v...)
}

func (l logger) Error(msg string, v ...any) {
	l.logger.Error(msg, v...)
}

func (l logger) Debug(msg string, v ...any) {
	l.logger.Debug(msg, v...)
}

func (l logger) Warn(msg string, v ...any) {
	l.logger.Warn(msg, v...)
}

func (l logger) Fatal(msg string) {
	panic(msg)
}

func (l logger) InfoWithContext(ctx context.Context, msg string, v ...any) {
	l.logger.InfoContext(ctx, msg, v...)
}

func (l logger) ErrorWithContext(ctx context.Context, msg string, v ...any) {
	l.logger.ErrorContext(ctx, msg, v...)
}

func (l logger) DebugWithContext(ctx context.Context, msg string, v ...any) {
	l.logger.DebugContext(ctx, msg, v...)
}

func (l logger) WarnWithContext(ctx context.Context, msg string, v ...any) {
	l.logger.WarnContext(ctx, msg, v...)
}

func (l logger) Log(level LogLevel, msg string, v ...any) {
	switch level {
	case INFO:
		l.logger.Info(msg, v...)
	case DEBUG:
		l.logger.Debug(msg, v...)
	case WARN:
		l.logger.Warn(msg, v...)
	case ERROR:
		l.logger.Error(msg, v...)
	case FATAL:
		panic(msg)
	}
}

func NewWithConfig(config Config) Logger {
	return logger{
		logger: slog.New(config.Handler),
	}
}

func New() Logger {
	config := Config{
		Handler: slog.NewJSONHandler(os.Stderr, nil),
	}
	return NewWithConfig(config)
}
