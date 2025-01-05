package search

import (
	"path/filepath"
	"strings"
)

// function to traverse the directory and files matching the pattern
func Search(pattern, path string) ([]string, error) {
	// Get all files from the directory
	files, err := Traverse(path, false)
	if err != nil {
		return nil, err
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
