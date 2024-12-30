package health_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrew-chon/mneme/server/internal/health"
	"github.com/andrew-chon/mneme/server/pkg/assert"
	"github.com/andrew-chon/mneme/server/pkg/logger"
)

func TestHealthHandler(t *testing.T) {
	logger := logger.New()
	h := health.NewHandler(logger)
	server := httptest.NewServer(http.HandlerFunc(h.HealthHandler))
	defer server.Close()

	resp, err := http.Get(server.URL) //nolint:all
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	var got health.Status
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		t.Errorf("error decoding body. Err: %v", err)
	}

	want := health.Status{
		Status: "available",
	}
	assert.Equal(t, got, want)
}
