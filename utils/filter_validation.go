package utils

func HasActiveFilters(filterOptions map[string]interface{}) bool {
	for _, value := range filterOptions {
		switch v := value.(type) {
		case string:
			if v != "" { // Non-empty string
				return true
			}
		case bool:
			if v { // True boolean
				return true
			}
		}
	}
	return false
}
