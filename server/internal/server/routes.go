package server

import (
	"net/http"

	"github.com/andrew-chon/mneme/server/internal/health"
	"github.com/andrew-chon/mneme/server/pkg/logger"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	health *health.Handler
}

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	handlers := getHandlers(s.logger)

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", handlers.health.HealthHandler)

	return r
}

func getHandlers(logger *logger.Logger) handler {
	health := health.New(logger)

	return handler{
		health: &health,
	}
}
