package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type SanityResponse struct {
	Result []any `json:"result"`
}

func SanityAPI(query string) string {
	baseURL := fmt.Sprintf(
		"https://%s.api.sanity.io/%s/data/query/production",
		os.Getenv("SANITY_PROJECT_ID"),
		os.Getenv("SANITY_API_VERSION"),
	)

	params := url.Values{}
	params.Add("query", query)
	params.Add("returnQuery", "true")

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func AllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := "*[_type=='project'] | order(_updatedAt desc)[0...4] | {title, slug, description, thumbnail}"

	res, err := http.Get(SanityAPI(query))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Server Error. Failed to fetch from Sanity API."})
		return
	}
	defer res.Body.Close()

	var sanityResponse SanityResponse
	if err := json.NewDecoder(res.Body).Decode(&sanityResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	// only return the `result` field
	json.NewEncoder(w).Encode(map[string][]any{"projects": sanityResponse.Result})
}
