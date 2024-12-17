package health

import (
	"fmt"
	"net/http"

	"github.com/andrew-chon/mneme/server/internal/logger"
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
	fmt.Fprintln(w, "status: available")

}
