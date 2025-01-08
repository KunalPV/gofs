package utils

import (
	"fmt"
	"path/filepath"
)

// ConvertToAbsolutePaths converts a list of file paths to absolute paths.
func ConvertToAbsolutePaths(files []string) ([]string, error) {
	var absPaths []string
	for _, file := range files {
		absPath, err := filepath.Abs(file)
		if err != nil {
			return nil, fmt.Errorf("could not convert %s to absolute path: %v", file, err)
		}
		absPaths = append(absPaths, absPath)
	}
	return absPaths, nil
}
