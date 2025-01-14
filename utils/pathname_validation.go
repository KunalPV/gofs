package utils

import (
	"os"
	"strings"
)

// ValidatePathname filters traversed paths based on the provided pathname.
// Returns a list of valid paths and a boolean indicating if any matches were found.
func ValidatePathname(traversedPaths []string, pathname string) ([]string, bool) {
	var validPaths []string

	// Normalize the pathname (remove trailing slashes)
	isRoot := pathname == "."
	normalizedPathname := strings.TrimSuffix(pathname, string(os.PathSeparator))

	for _, path := range traversedPaths {
		// Check if the path exists
		info, err := os.Stat(path)
		if err != nil {
			continue // Skip invalid paths
		}

		// Handle directories
		if info.IsDir() {
			// Match directory name or substring with the provided pathname
			if isRoot || strings.Contains(path, normalizedPathname) {
				validPaths = append(validPaths, path+string(os.PathSeparator))
			}
			continue
		}

		// Handle files
		if strings.Contains(path, normalizedPathname) {
			validPaths = append(validPaths, path)
		}
	}

	return validPaths, len(validPaths) > 0
}
