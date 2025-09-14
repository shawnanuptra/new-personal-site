package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// 1. Set the Content-Type header to "application/json"
		w.Header().Set("Content-Type", "application/json")

		// 2. Create a Go struct or map to hold your data.
		// A map is a simple way to return key-value pairs.
		response := map[string]string{"message": "pong"}

		// 3. Encode the Go data structure into a JSON string and write it to the response body.
		json.NewEncoder(w).Encode(response)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
