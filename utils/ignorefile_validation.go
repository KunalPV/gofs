package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// CheckIgnoreFlag determines if the ignore flag is set.
func CheckIgnoreFlag() bool {
	for _, arg := range os.Args {
		if arg == "-I" || arg == "--ignore" {
			return true
		}
	}
	return false
}

var (
	ignorePatternsCache     = make(map[string][]string) // Cache for preprocessed patterns
	ignorePatternsCacheLock sync.Mutex                  // Mutex for thread-safe access
)

// IsIgnored checks if a file or directory should be ignored based on patterns in ignore files.
func IsIgnored(path string) bool {
	// Find all .ignore files
	ignoreFiles, err := filepath.Glob(".*ignore")
	if err != nil || len(ignoreFiles) == 0 {
		return false // If no ignore files are found, nothing is ignored
	}

	for _, ignoreFile := range ignoreFiles {
		// Retrieve preprocessed patterns from cache or process them
		patterns := getCachedOrProcessPatterns(ignoreFile)

		for _, pattern := range patterns {
			if matchPattern(path, pattern) {
				return true
			}
		}
	}

	return false
}

// getCachedOrProcessPatterns retrieves cached patterns or processes the ignore file to generate them.
func getCachedOrProcessPatterns(ignoreFile string) []string {
	ignorePatternsCacheLock.Lock()
	defer ignorePatternsCacheLock.Unlock()

	// Check if patterns are already cached
	if cachedPatterns, exists := ignorePatternsCache[ignoreFile]; exists {
		return cachedPatterns
	}

	// Process the ignore file and cache the result
	patterns, err := parseIgnoreFile(ignoreFile)
	if err != nil {
		return nil
	}

	preprocessedPatterns := preprocessPatterns(patterns)
	ignorePatternsCache[ignoreFile] = preprocessedPatterns

	return preprocessedPatterns
}

// parseIgnoreFile reads patterns from a given ignore file.
func parseIgnoreFile(ignoreFile string) ([]string, error) {
	file, err := os.Open(ignoreFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		patterns = append(patterns, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return patterns, nil
}

// preprocessPatterns processes and normalizes all patterns from an ignore file.
func preprocessPatterns(patterns []string) []string {
	var processed []string
	for _, pattern := range patterns {
		// Check if the pattern starts with '*'
		if len(pattern) > 0 && pattern[0] == '*' {
			processed = append(processed, pattern) // Directly add wildcard patterns
		} else {
			// Normalize directory patterns or other non-wildcard patterns
			parsedPattern := parseIgnorePattern(pattern)
			if parsedPattern != "" {
				processed = append(processed, parsedPattern)
			}
		}
	}
	return processed
}

// parseIgnorePattern normalizes an ignore pattern to its root directory or file.
func parseIgnorePattern(pattern string) string {
	// Clean the pattern to remove leading/trailing slashes
	normalizedPattern := strings.Trim(pattern, "/")
	normalizedPattern = filepath.Clean(normalizedPattern)

	// Check if the pattern represents a directory (ends with "/*" or a trailing "/")
	if strings.HasSuffix(pattern, "/*") || strings.HasSuffix(pattern, string(filepath.Separator)) {
		// Extract the directory portion
		parts := strings.Split(normalizedPattern, string(filepath.Separator))
		if len(parts) > 0 {
			return parts[0]
		}
	}

	return normalizedPattern
}

// matchPattern checks if a path matches a pattern from an ignore file.
func matchPattern(path, pattern string) bool {
	// Normalize the path and pattern
	normalizedPath := filepath.Clean(path)
	normalizedPattern := filepath.Clean(pattern)

	// Match directory patterns
	if !strings.Contains(normalizedPattern, "*") {
		// Check if the path starts with the directory pattern
		if strings.HasPrefix(normalizedPath, normalizedPattern) {
			return true
		}
	}

	// Match patterns with wildcards
	matched, err := filepath.Match(normalizedPattern, filepath.Base(normalizedPath))
	if err != nil {
		return false
	}

	return matched
}
