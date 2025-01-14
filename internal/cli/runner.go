package cli

import (
	"fmt"
)

// PrintResults outputs the search results to the console
func PrintResults(results []string) {
	for _, result := range results {
		fmt.Println(result)
	}
}
