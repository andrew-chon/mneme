package health

import (
	"net/http"

	"github.com/andrew-chon/mneme/server/pkg/logger"
	"github.com/andrew-chon/mneme/server/pkg/response"
)

type Handler struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) Handler {
	return Handler{
		logger: logger,
	}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("healthHandler")

	status := Status{
		Status: "available",
	}

	err := response.Encode(w, 200, status)
	if err != nil {
		h.logger.Error("failed to encode health status", "err", err)
		response.InternalServerError(w, r, h.logger)
	}
}
