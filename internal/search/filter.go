package search

import (
	"os"
	"path/filepath"
	"strings"
)

// FilterByType filters files based on type: file, dir, symlink.
func FilterByType(files []string, fileType string) ([]string, error) {
	var filtered []string
	for _, file := range files {
		info, err := os.Lstat(file)
		if err != nil {
			continue
		}
		switch fileType {
		case "file":
			if info.Mode().IsRegular() {
				filtered = append(filtered, file)
			}
		case "dir":
			if info.IsDir() {
				filtered = append(filtered, file)
			}
		case "symlink":
			if info.Mode()&os.ModeSymlink != 0 {
				filtered = append(filtered, file)
			}
		}
	}
	return filtered, nil
}

// FilterByExtension filters files based on the file extension.
func FilterByExtension(files []string, ext string) ([]string, error) {
	var filtered []string
	for _, file := range files {
		if filepath.Ext(file) == ext {
			filtered = append(filtered, file)
		}
	}
	return filtered, nil
}

// FilterByCaseSensitivity applies case sensitivity filtering.
func FilterByCaseSensitivity(files []string, pattern string) ([]string, error) {
	var filtered []string
	for _, file := range files {
		if strings.Contains(file, pattern) {
			filtered = append(filtered, file)
		}
	}
	return filtered, nil
}
