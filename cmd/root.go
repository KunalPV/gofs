package cmd

import (
	"fmt"
	"gofs/internal/cli"
	"gofs/internal/filter"
	"gofs/internal/search"
	"gofs/internal/traverse"
	"gofs/utils"

	"github.com/spf13/cobra"
)

// Root command for the CLI
var rootCmd = &cobra.Command{
	Use:     "gofs <pattern> [pathname]",
	Short:   "gofs is a lightweight CLI tool for searching files.",
	Long:    `A program to find files and directories in your filesystem.`,
	PreRunE: cli.PrioritizeHelpAndVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Step 1: Validate command
		err := utils.ValidateCommand(cmd, args)
		if err != nil {
			return err // Command validation errors are returned to Cobra
		}

		// Step 2: Parse flags and arguments into a Config struct
		config := cli.ParseFlags(cmd, args)

		// Step 3: Perform traversal and pathname validation
		traversalResults, err := traverse.TraverseAndValidate(config.Root, config.Pathname, config.Depth, config.MaxThreads)
		if err != nil {
			return err // Handle traversal or pathname validation errors
		}

		// Step 4: Perform pattern check and validation
		effectivePattern, err := utils.HandlePattern(config.Pattern, config.GlobPattern)
		if err != nil {
			return fmt.Errorf("error determining pattern: %v", err)
		}

		// Step 5: Perform search (regex/common-string or glob) on traversalResults
		var searchResults []string
		searchResults, err = search.SearchPattern(effectivePattern, traversalResults, config.MaxThreads, config.GlobPattern != "")
		if err != nil {
			return fmt.Errorf("error during search: %v", err)
		}

		// Step 6: Apply filters if any active FilterOptions are provided
		if utils.HasActiveFilters(config.FilterOptions) {
			searchResults, err = filter.FilterResults(searchResults, config.FilterOptions)
			if err != nil {
				return fmt.Errorf("error during filtering: %v", err)
			}
		}

		// Step x: Print the results
		cli.PrintResults(searchResults)

		return nil
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err) // Handles errors and exits gracefully
	}
}

func init() {
	cli.DefineFlags(rootCmd)
}
