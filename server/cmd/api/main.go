package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andrew-chon/mneme/server/internal/server"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s/n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	server := server.NewServer()

	// Create a done channel to signal when the shutdown is complete
	done := make(chan struct{}, 1)

	go gracefulShutdown(ctx, server, done)

	go func() {
		log.Printf("listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("Error listening and servering: %s", err))
		}
	}()

	// Wait for the graceful shutdown to complete
	<-done
	fmt.Fprint(os.Stderr, "\nGraceful shutdown complete\n")

	return nil
}

func gracefulShutdown(ctx context.Context, server *http.Server, done chan<- struct{}) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
	}

	done <- struct{}{}
}
