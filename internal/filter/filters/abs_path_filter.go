package filters

import "path/filepath"

func AbsPathFilter(results []string) []string {
	var absolutePaths []string
	for _, file := range results {
		absPath, err := filepath.Abs(file)
		if err == nil {
			absolutePaths = append(absolutePaths, absPath)
		}
	}
	return absolutePaths
}
