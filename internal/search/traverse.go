package search

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Traverse lists files and directories in a given directory.
// If topLevelOnly is true, it only lists items in the top-level directory (non-recursive).
func Traverse(path string, topLevelOnly bool) ([]string, error) {
	var files []string

	if topLevelOnly {
		// Use os.ReadDir for top-level items only
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			files = append(files, filepath.Join(path, entry.Name()))
		}
	} else {
		// Use WalkDir for recursive traversal
		err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() {
				files = append(files, path)
			}

			return nil
		})

		// Check for errors from WalkDir
		if err != nil {
			return nil, err
		}
	}

	return files, nil
}
