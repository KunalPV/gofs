package cmd

import (
	"fmt"
	"gofs/internal/search"
	"os"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

// Root command for the CLI
var rootCmd = &cobra.Command{
	Use:   "gofs [options] <pattern> [pathname]",
	Short: "gofs is a lightweight CLI tool for searching files.",
	Long: `gofs is a fast and lightweight CLI tool implemented in Go for searching files 
in directories. It supports pattern matching, directory traversal, and more.`,
	Args: cobra.MaximumNArgs(2), // Allow 0, 1, or 2 arguments
	Run: func(cmd *cobra.Command, args []string) {
		// Handle --version flag
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("gofs version %s\n", version)
			os.Exit(0)
		}

		// Handle no arguments: List top-level files and directories
		if len(args) == 0 {
			path, _ := cmd.Flags().GetString("path")
			results, err := search.Traverse(path, true) // Top-level listing only
			if err != nil {
				fmt.Printf("Error listing files: %v\n", err)
				return
			}

			if len(results) == 0 {
				fmt.Println("No files or directories found.")
			} else {
				for _, result := range results {
					fmt.Println(result)
				}
			}
			return
		}

		// Handle arguments for searching
		pattern := args[0]
		path := "."
		if len(args) == 2 {
			path = args[1]
		}

		results, err := search.Search(pattern, path)
		if err != nil {
			fmt.Printf("Search failed: %v\n", err)
			return
		}

		if len(results) == 0 {
			fmt.Println("No files found.")
		} else {
			for _, result := range results {
				fmt.Println(result)
			}
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("path", "p", ".", "Pathname to search (default: current directory)")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of the utility")
}
