package search

import (
	"strings"
)

// ExcludePattern filters out files matching any pattern from the exclude list.
func ExcludePattern(files []string, excludes []string) ([]string, error) {
	var filtered []string
	if len(excludes) == 0 {
		return files, nil
	}
	for _, file := range files {
		excluded := false
		for _, pattern := range excludes {
			matched := strings.Contains(file, pattern)
			if matched {
				excluded = true
				break // Skip the file if it matches any pattern
			}
			if !excluded {
				filtered = append(filtered, file)
			}
		}
	}
	return filtered, nil
}
