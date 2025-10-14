package handlers

import (
	"net/http"

	"github.com/shawnanuptra/new-personal-site/backend/sanity"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := sanity.GetAllProjects()
	if err != nil {
		// check for sanity type error, and return the struct if it is
		if sanityErr, ok := err.(*sanity.SanityError[sanity.QueryError]); ok {
			writeJSONError(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"error":   sanityErr.Err,
			})
			return
		}

		writeJSONError(w, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string][]sanity.Project{"projects": *projects})
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("project")

	project, err := sanity.GetProject(slug)
	if err != nil {
		// check for sanity type error, and return the struct if it is
		if sanityErr, ok := err.(*sanity.SanityError[sanity.QueryError]); ok {
			writeJSONError(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
				"error":   sanityErr.Err,
			})
			return
		}

		writeJSONError(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]sanity.Project{"project": *project})
}
