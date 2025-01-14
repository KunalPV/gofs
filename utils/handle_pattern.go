package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
)

// HandlePattern determines the effective pattern based on the provided glob or regex pattern.
func HandlePattern(pattern string, globPattern string) (string, error) {
	// Case 1: Special case for default pattern "."
	if pattern == "." {
		return pattern, nil
	}

	// Case 2: Handle glob pattern
	if globPattern != "" {
		// Validate the glob pattern
		if !isValidGlob(globPattern) {
			return "", fmt.Errorf("invalid glob pattern: %s", globPattern)
		}
		// Return the globPattern as the effective pattern
		return globPattern, nil
	}

	// Case 3: Validate the pattern as a regex
	_, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("invalid regex pattern: %v", err)
	}

	// Return the pattern as the effective regex/common string
	return pattern, nil
}

// isValidGlob validates a glob pattern for correctness.
func isValidGlob(pattern string) bool {
	// Minimal validation for glob patterns
	if pattern == "" {
		return false
	}

	// Use filepath.Match to validate the pattern
	matches, err := filepath.Glob(pattern)
	return err == nil && len(matches) >= 0
}
