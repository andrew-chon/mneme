package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andrew-chon/mneme/server/pkg/logger"
)

type Config struct {
	Port int
}
type Server struct {
	config Config
	logger *logger.Logger
}

func NewServer(cfg Config, logger *logger.Logger) *http.Server {
	NewServer := &Server{
		config: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
