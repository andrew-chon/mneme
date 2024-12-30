package response

import (
	"net/http"

	"github.com/andrew-chon/mneme/server/pkg/logger"
)

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func newErrorResponse(statusCode int, message string) errorResponse {
	return errorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}
}

func NotFound(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	message := "The requested resource was not found"
	response := newErrorResponse(http.StatusNotFound, message)

	err := Encode(w, int(http.StatusNotFound), response)
	if err != nil {
		logger.ErrorContext(r.Context(), "Error encoding response", "err", err, "requestMethod", r.Method, "requestUrl", r.URL.String)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func InternalServerError(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	message := "The requested resource was not found"
	response := newErrorResponse(http.StatusInternalServerError, message)

	err := Encode(w, int(http.StatusInternalServerError), response)
	if err != nil {
		logger.ErrorContext(r.Context(), "Error encoding response", "err", err, "requestMethod", r.Method, "requestUrl", r.URL.String)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
