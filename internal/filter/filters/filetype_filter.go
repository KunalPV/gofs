package filters

import (
	"fmt"
	"os"
)

func FileTypeFilter(results []string, fileType string) ([]string, error) {
	var filtered []string
	for _, file := range results {
		info, err := os.Stat(file)
		if err != nil {
			continue // Skip invalid paths
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
		default:
			return nil, fmt.Errorf("invalid file type: %s", fileType)
		}
	}
	return filtered, nil
}
