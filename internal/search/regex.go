package search

import (
	"fmt"
	"path/filepath"
	"regexp"
)

// RegexFilter filters a list of file paths based on a regex pattern.
func RegexFilter(files []string, pattern string) ([]string, error) {
	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %v", err)
	}

	// Filter files based on regex match
	results := make([]string, 0, len(files))
	for _, file := range files {
		if re.MatchString(file) || re.MatchString(filepath.Base(file)) {
			results = append(results, file)
		}
	}

	return results, nil
}

// GlobFilter filters a list of file paths based on a glob pattern.
func GlobFilter(files []string, pattern string) ([]string, error) {
	// Ensure the glob pattern is valid
	if !isValidGlob(pattern) {
		return nil, fmt.Errorf("invalid glob pattern: %s", pattern)
	}

	results := make([]string, 0, len(files))
	for _, file := range files {
		matched, err := filepath.Match(pattern, filepath.Base(file))

		if err != nil {
			return nil, fmt.Errorf("error matching file '%s' with pattern '%s': %v", file, pattern, err)
		}

		if matched {
			results = append(results, file)
		}
	}

	return results, nil
}

// isValidGlob validates the glob pattern to prevent runtime issues.
func isValidGlob(pattern string) bool {
	// A minimal validation for patterns (e.g., empty strings are not valid).
	// Expand this function if stricter validation is needed.
	return len(pattern) > 0
}
