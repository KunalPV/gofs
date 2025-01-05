package search

import (
	"io/fs"
	"path/filepath"
)

// function to traverse the directory and files matching the pattern
func Search(directory string, pattern string) ([]string, error) {
	var results []string

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Name() == pattern {
			results = append(results, path)
		}

		return nil
	})

	return results, err
}
