package sanity_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/shawnanuptra/new-personal-site/backend/sanity"
)

const (
	mockProjectID  = "projectId"
	mockAPIVersion = "v2025-01-01"
)

func TestQueryURL(t *testing.T) {
	tests := []struct {
		name  string
		query string
	}{
		{
			name:  "Simple query",
			query: "*[_type=='project']",
		},
		{
			name:  "Query with spaces and punctuations",
			query: "*[_type=='blog' && slug.current=='a-new-beginning'][0]{title, markdownContent, publishedAt, series, entry}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			sanityProjectID := os.Getenv("SANITY_PROJECT_ID")
			sanityAPIVersion := os.Getenv("SANITY_API_VERSION")
			sanityBaseURL := os.Getenv("SANITY_BASE_URL")

			os.Setenv("SANITY_PROJECT_ID", mockProjectID)
			os.Setenv("SANITY_API_VERSION", mockAPIVersion)
			os.Setenv("SANITY_BASE_URL", "https://%s.api.sanity.io/%s/data/query/production")

			// teardown
			t.Cleanup(func() {
				os.Setenv("SANITY_PROJECT_ID", sanityProjectID)
				os.Setenv("SANITY_API_VERSION", sanityAPIVersion)
				os.Setenv("SANITY_BASE_URL", sanityBaseURL)
			})

			// assert
			actual := sanity.QueryURL(tt.query)

			// escape
			expected := fmt.Sprintf("https://projectId.api.sanity.io/v2025-01-01/data/query/production?query=%s", url.QueryEscape(tt.query))

			// Assert
			if actual != expected {
				t.Fatalf("Expected != actual. Expected: '%s', Actual: '%s'", expected, actual)
			}
		})
	}
}

func TestGetAllProjects(t *testing.T) {
	tests := []struct {
		name         string
		query        string
		expectsError bool
	}{
		{
			name:         "sanity Error",
			query:        "invalid-query",
			expectsError: true,
		},
		{
			name:         "decoding error",
			query:        "decoding-error",
			expectsError: true,
		},
		{
			name:         "correct",
			query:        "*[_type='projects']",
			expectsError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch tt.query {
				// test returns 400
				case "invalid-query":
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(sanity.SanityError[sanity.QueryError]{
						Err: sanity.QueryError{
							Description: "Sanity Error: Invalid Query",
							Query:       tt.query,
							Type:        "queryParseError",
						},
					})

				case "decoding-error":
					w.WriteHeader(http.StatusOK)
					// return broken json, so can trigger Decoding error
					w.Write([]byte(`{"result":[`))

				case "*[_type='projects']":
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(sanity.Response[[]sanity.Project]{
						Ms:       12,
						SyncTags: []string{"syncTag1", "syncTag2"},
						Result:   []sanity.Project{{Title: "Hello World"}},
					})
				}
			}))
			defer server.Close()

			// setup
			sanityProjectID := os.Getenv("SANITY_PROJECT_ID")
			sanityAPIVersion := os.Getenv("SANITY_API_VERSION")
			sanityBaseURL := os.Getenv("SANITY_BASE_URL")

			os.Setenv("SANITY_PROJECT_ID", mockProjectID)
			os.Setenv("SANITY_API_VERSION", mockAPIVersion)
			os.Setenv("SANITY_BASE_URL", fmt.Sprintf("%s/data/query/production/%%s/%%s", server.URL))

			// teardown
			t.Cleanup(func() {
				os.Setenv("SANITY_PROJECT_ID", sanityProjectID)
				os.Setenv("SANITY_API_VERSION", sanityAPIVersion)
				os.Setenv("SANITY_BASE_URL", sanityBaseURL)
			})

			projects, err := sanity.GetAllProjects()

			switch tt.query {
			case "invalid-query":
				if tt.expectsError && err == nil {
					t.Errorf("Expected to receive error, but error is nil")
				}

				// check error type - if it's NOT QueryError, Fail
				if _, ok := err.(*sanity.SanityError[sanity.QueryError]); !ok {
					t.Errorf("Expected error to be type SanityError[QueryError], instead error is:%v", err)
				}

			case "decoding-error":
				if tt.expectsError && err == nil {
					t.Errorf("Expected to receive error, but error is nil")
				}

				// check error type - if it's QueryError, Fail
				if _, ok := err.(*sanity.SanityError[sanity.QueryError]); ok {
					t.Errorf("Expected error to be type SanityError[QueryError], instead error is:%v", err)
				}

			case "*[_type='projects']":
				if err != nil {
					t.Errorf("Expected error is nil, but received error:%v", err)
				}
				if len(*projects) != 1 {
					t.Errorf("Expected length of projects returned to be 1, got:%v", len(*projects))
				}

				if (*projects)[0].Title != "Hello World" {
					t.Errorf("Expected project title to be 'Hello World', got:%v", (*projects)[0].Title)
				}
			}
		})
	}
}

func TestGetOneProject(t *testing.T) {
	tests := []struct {
		name         string
		slug         string
		expectsError bool
	}{
		{
			name:         "sanity Error",
			slug:         "invalid-query",
			expectsError: true,
		},
		{
			name:         "decoding error",
			slug:         "decoding-error",
			expectsError: true,
		},
		{
			name:         "correct",
			slug:         "my-project",
			expectsError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch tt.slug {
				// test returns 400
				case "invalid-query":
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(sanity.SanityError[sanity.QueryError]{
						Err: sanity.QueryError{
							Description: "Sanity Error: Invalid Query",
							Query:       tt.slug,
							Type:        "queryParseError",
						},
					})

				case "decoding-error":
					w.WriteHeader(http.StatusOK)
					// return broken json, so can trigger Decoding error
					w.Write([]byte(`{"result":[`))

				case "my-project":
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(sanity.Response[sanity.Project]{
						Ms:       12,
						SyncTags: []string{"syncTag1", "syncTag2"},
						Result:   sanity.Project{Title: "Hello World"},
					})
				}
			}))
			defer server.Close()

			// setup
			sanityProjectID := os.Getenv("SANITY_PROJECT_ID")
			sanityAPIVersion := os.Getenv("SANITY_API_VERSION")
			sanityBaseURL := os.Getenv("SANITY_BASE_URL")

			os.Setenv("SANITY_PROJECT_ID", mockProjectID)
			os.Setenv("SANITY_API_VERSION", mockAPIVersion)
			os.Setenv("SANITY_BASE_URL", fmt.Sprintf("%s/data/query/production/%%s/%%s", server.URL))

			// teardown
			t.Cleanup(func() {
				os.Setenv("SANITY_PROJECT_ID", sanityProjectID)
				os.Setenv("SANITY_API_VERSION", sanityAPIVersion)
				os.Setenv("SANITY_BASE_URL", sanityBaseURL)
			})

			project, err := sanity.GetProject(tt.slug)

			switch tt.slug {
			case "invalid-query":
				if tt.expectsError && err == nil {
					t.Errorf("Expected to receive error, but error is nil")
				}

				// check error type - if it's NOT QueryError, Fail
				if _, ok := err.(*sanity.SanityError[sanity.QueryError]); !ok {
					t.Errorf("Expected error to be type SanityError[QueryError], instead error is:%v", err)
				}

			case "decoding-error":
				if tt.expectsError && err == nil {
					t.Errorf("Expected to receive error, but error is nil")
				}

				// check error type - if it's QueryError, Fail
				if _, ok := err.(*sanity.SanityError[sanity.QueryError]); ok {
					t.Errorf("Expected error to be type SanityError[QueryError], instead error is:%v", err)
				}

			case "my-project":
				if err != nil {
					t.Errorf("Expected error is nil, but received error:%v", err)
				}

				                if (*project).Title != "Hello World" {
									t.Errorf("Expected project title to be 'Hello World', got:%v", (*project).Title)
								}			}
		})
	}
}
