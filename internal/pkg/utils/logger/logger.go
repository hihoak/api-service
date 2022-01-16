package logger

import (
	"context"
	"github.com/rs/zerolog"
	"hihoak/api-service/internal/pkg/utils"
	"os"
)

type LoggerKeyType uint8

const LoggerKey LoggerKeyType = 1

type Logger struct {
	logger *zerolog.Logger
}

func strStdOutputToFile(output string) *os.File {
	switch output {
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	default:
		return os.Stderr
	}
}

func strLevelToLevel(level string) zerolog.Level {
	switch level {
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "err":
		return zerolog.ErrorLevel
	case "debug":
		return zerolog.DebugLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

func NewDefault() Logger {
	logger := zerolog.New(os.Stderr).Level(zerolog.DebugLevel)
	return Logger{logger: &logger}
}

func FromConfigWithContext(ctx context.Context, config utils.Config) (*Logger, context.Context) {
	stdOutput := strStdOutputToFile(config.Logger.Output)
	logLevel := strLevelToLevel(config.Logger.Level)
	logger := zerolog.New(stdOutput).Level(logLevel)

	resLogger := Logger{logger: &logger}

	ctx = context.WithValue(ctx, LoggerKey, resLogger)

	return &resLogger, ctx
}

func FromContext(ctx context.Context) Logger {
	logger := ctx.Value(LoggerKey)
	if logger != nil {
		if resLogger, ok := logger.(Logger); ok {
			return resLogger
		}
	}

	return NewDefault()
}

func (l Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

func (l Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

func (l Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

func (l Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

func (l Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}
