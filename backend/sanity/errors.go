package sanity

import "fmt"

type SanityError[T any] struct {
	Err T `json:"error"`
}

/*
 * Query Responses
 * ref:https://www.sanity.io/docs/http-reference/query
 */

type QueryError struct {
	Description string `json:"description"`
	Query       string `json:"query"` // the query (e.g. "*[_type=='project']")
	Type        string `json:"type"`  // error type (e.g. "queryParseError")
}

func (e *SanityError[T]) Error() string {
	// Error message for QueryError
	if queryErr, ok := any(e.Err).(QueryError); ok {
		return fmt.Sprintf("Sanity API error: %s, Query: %s", queryErr.Type, queryErr.Query)
	}

	// Generic error message
	return fmt.Sprintf("Sanity API error - 400: %v", e.Err)
}
