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
	Use:   "gofs [Flag] <pattern> [pathname]",
	Short: "gofs is a lightweight CLI tool for searching files.",
	Long:  `A program to find files and directories in your filesystem`,
	Args:  cobra.MinimumNArgs(1), // Require at least 1 argument (pattern)
	Run: func(cmd *cobra.Command, args []string) {
		// Handle --version flag
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("gofs version %s\n", version)
			os.Exit(0)
		}

		// Get depth flag
		depth, _ := cmd.Flags().GetInt("depth")

		// Handle `gofs .`: List all files and directories recursively
		if len(args) == 1 && args[0] == "." {
			path, _ := cmd.Flags().GetString("pathname")
			results, err := search.Traverse(path, depth)
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

		// Extract pattern and optional pathname
		pattern := args[0]
		path := "."
		if len(args) == 2 {
			path = args[1]
		}

		// Check for flags
		regexFlag, _ := cmd.Flags().GetBool("regex")
		globFlag, _ := cmd.Flags().GetBool("glob")
		var results []string
		var err error

		if globFlag {
			// Perform glob search
			allFiles, err := search.Traverse(path, depth) // Get all files recursively
			if err != nil {
				fmt.Printf("Error traversing files: %v\n", err)
				return
			}

			results, err = search.GlobFilter(allFiles, pattern) // Filter using glob
			if err != nil {
				fmt.Printf("Error filtering files: %v\n", err)
				return
			}
		} else if regexFlag {
			// Perform regex search
			allFiles, err := search.Traverse(path, depth) // Get all files recursively
			if err != nil {
				fmt.Printf("Error traversing files: %v\n", err)
				return
			}

			results, err = search.RegexFilter(allFiles, pattern) // Filter using regex
			if err != nil {
				fmt.Printf("Error filtering files: %v\n", err)
				return
			}
		} else {
			// Perform normal substring search
			results, err = search.Search(pattern, path, depth)
		}

		// Handle errors
		if err != nil {
			fmt.Printf("Search failed: %v\n", err)
			return
		}

		// Output results
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
	rootCmd.PersistentFlags().StringP("pathname", "p", ".", "Pathname to search (default: current directory)")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of the utility")
	rootCmd.Flags().BoolP("regex", "r", false, "Use regex pattern for searching")
	rootCmd.Flags().BoolP("glob", "g", false, "Use glob pattern for searching")
	rootCmd.Flags().IntP("depth", "d", -1, "Limit the depth of directory traversal (-1 for unlimited depth)")

	rootCmd.SetUsageTemplate(`Usage:
  gofs [Flag] <pattern> [pathname]

Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}
	`)
}
