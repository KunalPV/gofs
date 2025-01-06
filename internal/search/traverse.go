package search

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// Traverse lists files and directories in a given directory up to a specified depth.
// If depth is -1, it means unlimited depth.
func Traverse(root string, depth int) ([]string, error) {
	var files []string

	// Use WalkDir for recursive traversal
	err := filepath.WalkDir(root, func(currentPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate the depth of the current file/directory
		relativePath, err := filepath.Rel(root, currentPath)
		if err != nil {
			return err
		}

		// Skip the root directory itself (".")
		if relativePath == "." {
			return nil
		}

		// Calculate the current depth
		currentDepth := strings.Count(relativePath, string(filepath.Separator)) + 1

		// Skip files/folders beyond the specified depth
		if depth != -1 && currentDepth > depth {
			if d.IsDir() {
				return filepath.SkipDir // Skip directories beyond the depth
			}
			return nil
		}

		// Add valid file or directory paths to the results
		files = append(files, currentPath)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
