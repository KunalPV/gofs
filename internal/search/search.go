package search

import (
	"fmt"
	"path/filepath"
	"strings"
)

type FilterOptions struct {
	RegexPattern    bool
	GlobPattern     bool
	CaseSensitive   bool
	ExcludePatterns []string
	FileType        bool
	Extension       bool
}

// Search finds files matching a pattern (glob, regex, or substring) in the specified directory.
func Search(pattern, path string, depth int, options FilterOptions) ([]string, error) {
	// Get all files from the directory
	files, err := Traverse(path, depth)
	if err != nil {
		return nil, fmt.Errorf("error traversing files: %v", err)
	}

	// Filter files based on the provided options
	filteredFiles, err := applyFilters(files, pattern, options)
	if err != nil {
		return nil, fmt.Errorf("error applying filters: %v", err)
	}

	// Exclude files based on exclude patterns
	finalFiles, err := ExcludePattern(filteredFiles, options.ExcludePatterns)
	if err != nil {
		return nil, fmt.Errorf("error excluding files: %v", err)
	}

	return finalFiles, nil
}

// applyFilters applies the necessary filtering logic based on the FilterOptions.
func applyFilters(files []string, pattern string, options FilterOptions) ([]string, error) {
	switch {
	case options.RegexPattern:
		return RegexFilter(files, pattern)
	case options.GlobPattern:
		return GlobFilter(files, pattern)
	case options.CaseSensitive:
		return FilterByCaseSensitivity(files, pattern)
	case options.FileType:
		return FilterByType(files, pattern)
	case options.Extension:
		return FilterByExtension(files, pattern)
	default:
		// Default to normal substring search
		return filterBySubstring(files, pattern), nil
	}
}

// filterBySubstring performs a simple substring search on the files.
func filterBySubstring(files []string, pattern string) []string {
	var results []string
	for _, file := range files {
		if strings.Contains(file, pattern) || strings.Contains(filepath.Base(file), pattern) {
			results = append(results, file)
		}
	}
	return results
}
