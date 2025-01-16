package formats

import (
	"os"
	"path/filepath"
)

func AbsPathFormat(results []string) []string {
	var absolutePaths []string
	for _, file := range results {
		absPath, err := filepath.Abs(file)
		if err != nil {
			continue
		} else {
			info, err := os.Stat(file)
			if err != nil {
				continue
			}
			if info.IsDir() {
				absPath += string(filepath.Separator)
			}
			absolutePaths = append(absolutePaths, absPath)
		}
	}
	return absolutePaths
}
