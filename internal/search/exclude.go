package search

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ExcludePattern filters out files matching any pattern from the exclude list.
func ExcludePattern(files []string, excludes []string) ([]string, error) {
	var filtered []string
	for _, file := range files {
		excluded := false
		for _, pattern := range excludes {
			// Use filepath.Match for glob-like pattern matching
			matched, err := filepath.Match(pattern, file)
			if err != nil {
				return nil, fmt.Errorf("invalid exclude pattern: %s", pattern)
			}
			if matched || strings.HasPrefix(file, filepath.Clean(pattern)) {
				excluded = true
				break
			}
		}
		if !excluded {
			filtered = append(filtered, file)
		}
	}
	return filtered, nil
}
