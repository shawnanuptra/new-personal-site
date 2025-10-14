package handlers

import (
	"encoding/json"
	"net/http"
)

type APISuccess[T any] struct {
	Data T `json:"data"`
}

type APIError[T any] struct {
	Error T `json:"error"`
}

func writeJSON[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APISuccess[T]{Data: data})
}

func writeJSONError[T any](w http.ResponseWriter, status int, error T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIError[T]{Error: error})
}
