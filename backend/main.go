package main

import (
	"log"
	"net/http"

	"github.com/shawnanuptra/new-personal-site/backend/handlers"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/ping", handlers.Ping)

	log.Fatal(http.ListenAndServe(":8080", router))
}
