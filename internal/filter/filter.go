package filter

import (
	"fmt"
	"gofs/internal/filter/filters"
)

func FilterResults(searchResults []string, filterOptions map[string]interface{}) ([]string, error) {
	filteredResults := searchResults
	var err error

	// Apply filters one by one
	for key, value := range filterOptions {
		switch key {
		case "Extension":
			if ext, ok := value.(string); ok && ext != "" {
				filteredResults = filters.ExtensionFilter(filteredResults, ext)
			}
		case "FileType":
			if fileType, ok := value.(string); ok && fileType != "" {
				filteredResults, err = filters.FileTypeFilter(filteredResults, fileType)
				if err != nil {
					return nil, fmt.Errorf("error applying file type filter: %v", err)
				}
			}
		case "Exclude":
			if excludePattern, ok := value.(string); ok && excludePattern != "" {
				filteredResults, err = filters.ExcludeFilter(filteredResults, excludePattern)
				if err != nil {
					return nil, fmt.Errorf("error applying exclude filter: %v", err)
				}
			}
		}
	}

	return filteredResults, nil
}
