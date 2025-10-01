package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shawnanuptra/new-personal-site/backend/handlers"
)

func TestPing(t *testing.T) {
	tests := []struct {
		name string
		w    *httptest.ResponseRecorder
		r    *http.Request
	}{
		{
			name: "ping returns 'pong'",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/ping", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlers.Ping(tt.w, tt.r)
			if tt.w.Code != http.StatusOK {
				t.Errorf("Expected status code %d, but got %d", http.StatusOK, tt.w.Code)
			}

			var responseBody map[string]string
			err := json.NewDecoder(tt.w.Body).Decode(&responseBody)
			if err != nil {
				t.Fatalf("Failed to decode JSON response: %v", err)
			}

			if responseBody["message"] != "pong" {
				t.Errorf("Expected message 'pong', but got '%s'", responseBody["message"])
			}
		})
	}
}
