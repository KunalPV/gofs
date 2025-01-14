package filters

import (
	"fmt"
	"gofs/utils"
	"path/filepath"
)

func ExcludeFilter(results []string, pattern string) ([]string, error) {
	if !utils.IsValidGlob(pattern) {
		return nil, fmt.Errorf("invalid glob pattern: %s", pattern)
	}

	var filtered []string
	for _, file := range results {
		matched, err := filepath.Match(pattern, filepath.Base(file))
		if err != nil {
			return nil, fmt.Errorf("error matching glob pattern: %v", err)
		}
		if !matched {
			filtered = append(filtered, file)
		}
	}
	return filtered, nil
}
