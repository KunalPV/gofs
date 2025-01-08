package cmd

import "fmt"

// printResults outputs the search results
func printResults(results []string) {
	if len(results) == 0 {
		fmt.Println("No files found.")
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
