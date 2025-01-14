package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// CheckHiddenFlag determines if the hidden flag is set.
func CheckHiddenFlag() bool {
	for _, arg := range os.Args {
		if arg == "-H" || arg == "--hidden" {
			return true
		}
	}
	return false
}

// IsHidden determines if a file or directory is hidden.
// A hidden file starts with a dot (.) in its name but excludes the current directory (.)
func IsHidden(path string) bool {
	base := filepath.Base(path) // Extracts the file or directory name
	// Exclude the root directory itself (.)
	return base != "." && strings.HasPrefix(base, ".")
}
