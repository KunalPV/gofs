package cmd

import (
	"fmt"
	"gofs/internal/parallel"
	"gofs/internal/search"
	"gofs/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

// executeSearch runs the search logic
func executeSearch(cmd *cobra.Command, args []string) error {
	// Handle --version flag early
	versionFlag, _ := cmd.Flags().GetBool("version")
	if versionFlag {
		fmt.Printf("gofs version %s\n", version)
		os.Exit(0) // Exit early to prevent further execution
	}

	// Retrieve and validate thread count
	threadFlag, _ := cmd.Flags().GetInt("threads")
	maxThreads, err := utils.GetValidatedThreads(threadFlag)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	// Retrieve flags
	depth, _ := cmd.Flags().GetInt("max-depth")
	path, _ := cmd.Flags().GetString("pathname")
	regexFlag, _ := cmd.Flags().GetBool("regex")
	globFlag, _ := cmd.Flags().GetBool("glob")
	excludePatterns, _ := cmd.Flags().GetStringSlice("exclude")
	fileType, _ := cmd.Flags().GetBool("file-type")
	extension, _ := cmd.Flags().GetBool("extension")
	caseSensitive, _ := cmd.Flags().GetBool("case-sensitive")
	abspath, _ := cmd.Flags().GetBool("abs-path")

	// Validate conflicting flags
	if regexFlag && globFlag {
		return fmt.Errorf("--regex and --glob cannot be used together")
	}
	if fileType && extension {
		return fmt.Errorf("--file-type and --extension cannot be used together")
	}

	options := search.FilterOptions{
		RegexPattern:    regexFlag,
		GlobPattern:     globFlag,
		CaseSensitive:   caseSensitive,
		ExcludePatterns: excludePatterns,
		FileType:        fileType,
		Extension:       extension,
		AbsPath:         abspath,
	}

	// Handle `gofs .`: List all files and directories recursively
	if len(args) == 1 && args[0] == "." {
		results, err := search.Traverse(path, depth, maxThreads)
		if err != nil {
			return fmt.Errorf("error listing files: %v", err)
		}
		printResults(results)
		return nil
	}

	// Ensure there's a valid pattern argument
	if len(args) == 0 {
		return fmt.Errorf("missing required <pattern> argument")
	}

	// Extract pattern and optional pathname
	pattern := args[0]
	if len(args) == 2 {
		path = args[1]
	}

	// Perform traversal and search
	allFiles, err := search.Traverse(path, depth, maxThreads)
	if err != nil {
		return fmt.Errorf("error during traversal: %v", err)
	}

	// Perform parallel search
	results, err := parallel.ExecuteWithThreads(
		func() ([]string, error) { return allFiles, nil }, // Task to provide files for search
		func(file string) (string, error) { // Task to process each file
			matches, err := search.Search(pattern, path, []string{file}, options)
			if err != nil {
				return "", err
			}
			return strings.Join(matches, "\n"), nil
		},
		maxThreads,
	)

	if err != nil {
		return fmt.Errorf("search failed: %v", err)
	}

	// Remove duplicates from results
	uniqueResults := make(map[string]struct{})
	for _, result := range results {
		uniqueResults[result] = struct{}{}
	}

	finalResults := []string{}
	for result := range uniqueResults {
		finalResults = append(finalResults, result)
	}

	printResults(finalResults)
	return nil
}
