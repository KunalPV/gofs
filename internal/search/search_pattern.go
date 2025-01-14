package search

import (
	"fmt"
	"gofs/utils"
)

// SearchPattern orchestrates the search logic, validating threads and leveraging parallel search.
func SearchPattern(pattern string, traversalResults []string, maxThreads int, isGlob bool) ([]string, error) {
	// Validate maxThreads
	validThreads, err := utils.ValidateMaxThreads(maxThreads)
	if err != nil {
		return nil, fmt.Errorf("error validating maxThreads: %v", err)
	}

	// If the pattern is ".", return traversal results directly
	if pattern == "." {
		return traversalResults, nil
	}

	// Execute the search using parallel threads
	searchResults, err := SearchWithThreads(pattern, traversalResults, validThreads, isGlob)
	if err != nil {
		return nil, fmt.Errorf("error during search: %v", err)
	}

	// Return the search results
	return searchResults, nil
}
