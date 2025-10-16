package handlers

import (
	"net/http"

	"github.com/shawnanuptra/new-personal-site/backend/sanity"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := sanity.GetAllProjects()
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
