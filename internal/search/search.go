package search

import (
	"path/filepath"
	"strings"
)

// Search finds files matching a pattern (glob, regex, or substring) in the specified directory.
func Search(pattern, path string, depth int) ([]string, error) {
	// Get all files from the directory
	files, err := Traverse(path, depth)
	if err != nil {
		return nil, err
	}

	// Check if the pattern is a glob (contains wildcard characters like `*` or `?`)
	if strings.ContainsAny(pattern, "*?[]") {
		return GlobFilter(files, pattern)
	}

	// Check if the pattern is a regex (starts with "re:")
	if strings.HasPrefix(pattern, "re:") {
		regexPattern := strings.TrimPrefix(pattern, "re:")
		return RegexFilter(files, regexPattern)
	}

	// Filter files by the pattern
	var results []string
	for _, file := range files {
		if strings.Contains(file, pattern) || strings.Contains(filepath.Base(file), pattern) {
			results = append(results, file)
		}
	}

	return results, nil
}
