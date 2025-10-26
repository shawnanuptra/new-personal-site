package handlers

import (
	"net/http"

	"github.com/shawnanuptra/new-personal-site/backend/sanity"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	count, err := getCount(r)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
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

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	count, err := getCount(r)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
	}

	projects, err := sanity.GetProjects(count)
	if err != nil {
		HandleSanityError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string][]sanity.Project{"projects": *projects})
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("project")

	project, err := sanity.GetProject(slug)
	if err != nil {
		HandleSanityError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]sanity.Project{"project": *project})
}
