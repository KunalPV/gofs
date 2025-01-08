package search

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// Traverse lists files and directories in a given directory up to a specified depth.
// If depth is -1, it means unlimited depth.
func Traverse(root string, depth int) ([]string, error) {
	if depth < -1 {
		return nil, fmt.Errorf("invalid depth: %d, depth must be -1 or a non-negative value", depth)
	}

	var files []string

	err := filepath.WalkDir(root, func(currentPath string, d fs.DirEntry, err error) error {
		if err != nil {
			// Return any traversal errors
			return fmt.Errorf("error accessing %s: %v", currentPath, err)
		}

		// Skip the root directory itself
		if currentPath == root {
			return nil
		}

		// Calculate the relative path and depth
		relativePath, err := filepath.Rel(root, currentPath)
		if err != nil {
			return fmt.Errorf("error calculating relative path for %s: %v", currentPath, err)
		}

		currentDepth := strings.Count(relativePath, string(filepath.Separator)) + 1

		// Skip entries beyond the specified depth
		if depth != -1 && currentDepth > depth {
			if d.IsDir() {
				return filepath.SkipDir // Skip entire directories beyond the depth
			}
			return nil
		}

		// Add the current path to the results
		files = append(files, currentPath)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error during traversal: %v", err)
	}

	return files, nil
}
