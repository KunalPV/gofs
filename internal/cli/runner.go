package cli

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ANSI color codes following fd's color scheme
const (
	colorReset   = "\033[0m"  // Reset color
	colorCode    = "\033[32m" // Green for directories
	colorData    = "\033[34m" // Blue for data files
	colorDir     = "\033[36m" // Cyan for code files
	colorText    = "\033[33m" // Yellow for text files
	colorDefault = "\033[37m" // White for other files
	colorHidden  = "\033[90m" // Grey for hidden files
	colorExec    = "\033[31m" // Red for executables and scripts
)

// getColorForFileType determines the color based on file extension
func getColorForFileType(file string) string {
	ext := strings.ToLower(filepath.Ext(file))

	switch ext {
	case ".go", ".py", ".cpp", ".c", ".java", ".js", ".ts", ".rs":
		return colorCode // Code files
	case ".json", ".csv", ".xml", ".yaml", ".yml":
		return colorData // Data files
	case ".sh", ".bat", ".ps1":
		return colorExec // Executables and scripts
	case ".md", ".txt", ".log":
		return colorText // Text files
	default:
		return colorDefault // Default color for other files
	}
}

// PrintResults automatically handles long and short formats
func PrintResults(results []string) {
	for _, result := range results {
		// Check if the result contains a timestamp pattern to identify long list output
		if strings.Contains(result, "IST") {
			// Split metadata and pathname
			infoEndIndex := strings.LastIndex(result, "IST") + len("IST")
			info := result[:infoEndIndex]
			pathname := result[infoEndIndex+1:]

			// Print the uncolored metadata and colored path
			fmt.Print(info + " ")
			printColoredPathname(pathname)
		} else {
			// Treat it as a short list
			printColoredPathname(result)
		}
		fmt.Println()
	}
}

// printColoredPathname applies color only to the pathname components
func printColoredPathname(pathname string) {
	parts := strings.Split(pathname, string(filepath.Separator))
	for i, part := range parts {
		if part == "" {
			continue // Ignore empty parts for better formatting
		}

		if i == len(parts)-1 { // Last part (the file itself)
			fmt.Printf("%s%s%s", getColorForFileType(part), part, colorReset)
		} else { // Intermediate directories
			fmt.Printf("%s%s%s%s", colorDir, part, colorReset, string(filepath.Separator))
		}
	}
}
