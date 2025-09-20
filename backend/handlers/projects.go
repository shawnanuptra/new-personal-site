package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shawnanuptra/new-personal-site/backend/sanity"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projects, err := sanity.GetAllProjects()
	if err != nil {
		// check for sanity type error, and return the struct if it is
		if sanityErr, ok := err.(*sanity.SanityError[sanity.QueryError]); ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": err.Error(),
				"error":   sanityErr.Err,
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string][]sanity.Project{"projects": projects})
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slug := r.PathValue("project")

	project, err := sanity.GetProject(slug)
	if err != nil {
		// check for sanity type error, and return the struct if it is
		if sanityErr, ok := err.(*sanity.SanityError[sanity.QueryError]); ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{
				"message": err.Error(),
				"error":   sanityErr.Err,
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]sanity.Project{"project": *project})
}
