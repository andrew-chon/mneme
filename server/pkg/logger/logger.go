package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func New() *Logger {
	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stderr, opts)
	logger := slog.New(handler)

	return &Logger{
		Logger: logger,
	}
}
