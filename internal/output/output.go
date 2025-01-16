package output

import (
	"gofs/internal/output/formats"
)

func FormatResults(results []string, formatOptions map[string]interface{}) []string {
	formatedResults := results

	// Apply filters one by one
	for key, value := range formatOptions {
		switch key {
		case "AbsolutePath":
			if absPath, ok := value.(bool); ok && absPath {
				formatedResults = formats.AbsPathFormat(formatedResults)
			}
		case "LongList":
			if longList, ok := value.(bool); ok && longList {
				formatedResults = formats.LongListFormat(formatedResults)
			}
			// case "Hyperlink":
			// 	if hyperlink, ok := value.(bool); ok && hyperlink {
			// 		formatedResults = formats.HyperlinkFormat(formatedResults)
			// 	}
		}
	}

	return formatedResults
}
