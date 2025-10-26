// Package sanity
package sanity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func QueryURL(query string) string {
	baseURL := fmt.Sprintf(
		os.Getenv("SANITY_BASE_URL"),
		os.Getenv("SANITY_PROJECT_ID"),
		os.Getenv("SANITY_API_VERSION"),
	)

	params := url.Values{}
	params.Add("query", query)

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func executeQuery[T any](query string) (*T, error) {
	res, err := client.Get(QueryURL(query))
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

	var sanityResponse Response[T]
	if err := json.NewDecoder(res.Body).Decode(&sanityResponse); err != nil {
		return nil, err
	}

	return &sanityResponse.Result, nil
}

func GetProjects(count int) (*[]Project, error) {
	query := fmt.Sprintf("*[_type=='project'] | order(_updatedAt desc)[0...%v] | {title, 'slug':slug.current, description, 'thumbnailUrl':thumbnail.asset->url}", count)
	return executeQuery[[]Project](query)
}

func GetProject(slug string) (*Project, error) {
	query := fmt.Sprintf("*[_type=='project' && slug.current=='%s'][0]{title, markdownContent, publishedAt, series, entry}", slug)
	return executeQuery[Project](query)
}

func GetBlogs(count int) (*[]Blog, error) {
	query := fmt.Sprintf("*[_type=='blog'] | order(_updatedAt desc)[0...%v] | {title, 'slug':slug.current, description}", count)
	return executeQuery[[]Blog](query)
}

func GetBlog(slug string) (*Blog, error) {
	query := fmt.Sprintf("*[_type=='blog' && slug.current=='%s'][0]{title, markdownContent, publishedAt, series, entry}", slug)
	return executeQuery[Blog](query)
}
