package health

import (
	"net/http"

	"github.com/andrew-chon/mneme/server/pkg/json"
	"github.com/andrew-chon/mneme/server/pkg/logger"
)

type Handler struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) Handler {
	return Handler{
		logger: logger,
	}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("healthHandler")

	status := Status{
		Status: "available",
	}

	err := json.Encode(w, 200, status)
	if err != nil {
		h.logger.Error("failed to encode health status", "err", err)
	}
}
