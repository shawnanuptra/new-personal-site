package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/shawnanuptra/new-personal-site/backend/handlers"
)

func main() {
	// get APP_ENV, default to local
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "local"
	}

	// attempt to load the env file, crash if fails
	if err := godotenv.Load(fmt.Sprintf(".env.%s", env)); err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("/ping", handlers.Ping)

	router.HandleFunc("GET /projects", handlers.GetAllProjects)
	router.HandleFunc("GET /projects/{project}", handlers.GetProject)

	log.Fatal(http.ListenAndServe(":8080", router))
}
