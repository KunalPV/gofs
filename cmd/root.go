package cmd

import (
	"github.com/spf13/cobra"
)

// Root command for the CLI
var rootCmd = &cobra.Command{
	Use:     "gofs [Flag] <pattern> [pathname]",
	Short:   "gofs is a lightweight CLI tool for searching files.",
	Long:    `A program to find files and directories in your filesystem.`,
	PreRunE: validateFlags, // Validation logic
	RunE:    executeSearch, // Main execution logic
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err) // Handles errors and exits gracefully
	}
}

func init() {
	defineFlags() // Initialize and define flags
}
