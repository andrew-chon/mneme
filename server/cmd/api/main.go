package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"time"

	"github.com/andrew-chon/mneme/server/internal/env"
	"github.com/andrew-chon/mneme/server/internal/logger"
	"github.com/andrew-chon/mneme/server/internal/server"
)

func main() {
	ctx := context.Background()
	logger := logger.New()
	if err := run(ctx, logger); err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *logger.Logger) error {
	var cfg server.Config

	cfg.Port = env.GetInt("PORT", 4444)

	server := server.NewServer(cfg, logger)

	shutdownErrorChan := make(chan error)
	go gracefulShutdown(ctx, server, shutdownErrorChan)

	logger.Info("Starting server", slog.Group("server", "addr", server.Addr))
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErrorChan
	if err != nil {
		return err
	}
	// Wait for the graceful shutdown to complete
	logger.Info("Gracefully shutdown error", slog.Group("server", "addr", server.Addr))
	return nil
}

func gracefulShutdown(ctx context.Context, server *http.Server, errChan chan<- error) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	errChan <- server.Shutdown(shutdownCtx)
}
