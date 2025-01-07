package search

import (
	"fmt"
	"path/filepath"
	"strings"
)

type FilterOptions struct {
	RegexPattern    bool
	GlobPattern     bool
	CaseSensitive   bool
	ExcludePatterns []string
	FileType        bool
	Extension       bool
}

// Search finds files matching a pattern (glob, regex, or substring) in the specified directory.
func Search(pattern, path string, depth int, options FilterOptions) ([]string, error) {
	var results []string
	var err error

	// Get all files from the directory
	files, err := Traverse(path, depth)
	if err != nil {
		fmt.Printf("Error traversing files: %v\n", err)
		return nil, err
	}

	if options.RegexPattern {
		// Perform regex search
		results, err := RegexFilter(files, pattern) // Filter using glob
		if err != nil {
			fmt.Printf("Error filtering files: %v\n", err)
			return nil, err
		}
		results, err = ExcludePattern(results, options.ExcludePatterns) // Exclude files
		if err != nil {
			fmt.Printf("Error removing files: %v\n", err)
			return nil, err
		}
		return results, nil
	} else if options.GlobPattern {
		// Perform glob search
		results, err = GlobFilter(files, pattern) // Filter using regex
		if err != nil {
			fmt.Printf("Error filtering files: %v\n", err)
			return nil, err
		}
		results, err = ExcludePattern(results, options.ExcludePatterns) // Exclude files
		if err != nil {
			fmt.Printf("Error removing files: %v\n", err)
			return nil, err
		}
		return results, nil
	} else if options.CaseSensitive {
		// Perform case-sensitive search
		results, err = FilterByCaseSensitivity(files, pattern) // Filter using case sensitivity
		if err != nil {
			fmt.Printf("Error filtering files: %v\n", err)
			return nil, err
		}
		results, err = ExcludePattern(results, options.ExcludePatterns) // Exclude files
		if err != nil {
			fmt.Printf("Error removing files: %v\n", err)
			return nil, err
		}
		return results, nil
	} else if options.FileType {
		// Perform file type search
		results, err = FilterByType(files, pattern) // Filter by file type
		if err != nil {
			fmt.Printf("Error filtering files: %v\n", err)
			return nil, err
		}
		results, err = ExcludePattern(results, options.ExcludePatterns) // Exclude files
		if err != nil {
			fmt.Printf("Error removing files: %v\n", err)
			return nil, err
		}
		return results, nil
	} else if options.Extension {
		// Perform file extension search
		results, err = FilterByExtension(files, pattern) // Filter by file extension
		if err != nil {
			fmt.Printf("Error filtering files: %v\n", err)
			return nil, err
		}
		results, err = ExcludePattern(results, options.ExcludePatterns) // Exclude files
		if err != nil {
			fmt.Printf("Error removing files: %v\n", err)
			return nil, err
		}
		return results, nil
	} else {
		// Perform normal pattern search
		for _, file := range files {
			if strings.Contains(file, pattern) || strings.Contains(filepath.Base(file), pattern) {
				results = append(results, file)
			}
		}
		results, err = ExcludePattern(results, options.ExcludePatterns) // Exclude files
		if err != nil {
			fmt.Printf("Error removing files: %v\n", err)
			return nil, err
		}
		return results, nil
	}
}
