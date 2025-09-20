// Package sanity
package sanity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func QueryURL(query string) string {
	baseURL := fmt.Sprintf(
		"https://%s.api.sanity.io/%s/data/query/production",
		os.Getenv("SANITY_PROJECT_ID"),
		os.Getenv("SANITY_API_VERSION"),
	)

	params := url.Values{}
	params.Add("query", query)

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func GetAllProjects() (projects []Project, err error) {
	query := "*[_type=='project'] | order(_updatedAt desc)[0...4] | {title, 'slug':slug.current, description, 'thumbnailUrl':thumbnail.asset->url}"

	// ref:https://www.sanity.io/docs/http-reference/query
	// hit the api: only return 200 or 400
	res, err := http.Get(QueryURL(query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 400 {
		var sanityError SanityError[QueryError]
		if err := json.NewDecoder(res.Body).Decode(&sanityError); err != nil {
			return nil, err
		}
		return nil, &sanityError
	}

	// decode response
	var sanityResponse Response[[]Project]
	if err := json.NewDecoder(res.Body).Decode(&sanityResponse); err != nil {
		return nil, err
	}

	return sanityResponse.Result, nil
}

func GetProject(slug string) (project *Project, err error) {
	query := fmt.Sprintf("*[_type=='blog' && slug.current=='%s'][0]{title, markdownContent, publishedAt, series, entry}", slug)

	res, err := http.Get(QueryURL(query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 400 {
		var sanityError SanityError[QueryError]
		if err := json.NewDecoder(res.Body).Decode(&sanityError); err != nil {
			return nil, err
		}
		return nil, &sanityError
	}

	// decode response
	var sanityResponse Response[Project]
	if err := json.NewDecoder(res.Body).Decode(&sanityResponse); err != nil {
		return nil, err
	}

	return &sanityResponse.Result, nil
}
