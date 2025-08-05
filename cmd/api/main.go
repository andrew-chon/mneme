package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/andrew-chon/mneme/pkg/logger"
)

func main() {
	ctx := context.Background()

	logger.SetGlobalContextLogger()

	if err := run(ctx); err != nil {
		trace := string(debug.Stack())
		slog.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	slog.InfoContext(ctx, "Starting")

	return errors.New("hello world")
}
