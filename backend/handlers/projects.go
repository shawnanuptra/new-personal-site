package handlers

import (
	"net/http"
	"strconv"

	"github.com/shawnanuptra/new-personal-site/backend/sanity"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	countStr := r.URL.Query().Get("count")

	count := 4 // defaults to 4
	if countStr != "" {
		c, err := strconv.Atoi(countStr)
		if err != nil {
			writeJSONError(w, http.StatusBadRequest, "Count should be an integer")
			return
		}

		count = c
	}

	projects, err := sanity.GetProjects(count)
	if err != nil {
		HandleSanityError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string][]sanity.Project{"projects": *projects})
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("project")

	project, err := sanity.GetProject(slug)
	if err != nil {
		HandleSanityError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]sanity.Project{"project": *project})
}
