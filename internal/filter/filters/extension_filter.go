package filters

import "strings"

// extensionFilter filters results by file extension.
func ExtensionFilter(results []string, ext string) []string {
	var filtered []string
	for _, file := range results {
		if strings.HasSuffix(file, "."+ext) {
			filtered = append(filtered, file)
		}
	}
	return filtered
}
