package search

import (
	"path/filepath"
	"regexp"
)

// RegexFilter filters a list of file paths based on a regex pattern.
func RegexFilter(files []string, pattern string) ([]string, error) {
	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	// Filter files based on regex match
	var results []string
	for _, file := range files {
		if re.MatchString(file) || re.MatchString(filepath.Base(file)) {
			results = append(results, file)
		}
	}

	return results, nil
}

// GlobFilter filters a list of file paths based on a glob pattern.
func GlobFilter(files []string, pattern string) ([]string, error) {
	var results []string

	for _, file := range files {
		matched, err := filepath.Match(pattern, filepath.Base(file))
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, file)
		}
	}

	return results, nil
}
