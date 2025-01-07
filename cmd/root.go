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
	Args: func(cmd *cobra.Command, args []string) error {
		// Check if --version or --help flags are set
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			return nil // Allow zero arguments if --version is set
		}
		if len(args) < 1 {
			return fmt.Errorf("requires at least 1 arg(s), only received 0")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Handle --version flag
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("gofs version %s\n", version)
			os.Exit(0)
		}

		// Get flags
		depth, _ := cmd.Flags().GetInt("depth")
		regexFlag, _ := cmd.Flags().GetBool("regex")
		globFlag, _ := cmd.Flags().GetBool("glob")
		excludePatterns, _ := cmd.Flags().GetStringSlice("exclude")
		fileType, _ := cmd.Flags().GetBool("filetype")
		extension, _ := cmd.Flags().GetBool("extension")
		caseSensitive, _ := cmd.Flags().GetBool("casesensitive")

		// Check conflicting flags
		if regexFlag && globFlag {
			fmt.Println("Error: --regex and --glob cannot be used together.")
			os.Exit(1)
		}
		if fileType && extension {
			fmt.Println("Error: --filetype and --extension cannot be used together.")
			os.Exit(1)
		}
		if (regexFlag || globFlag) && caseSensitive {
			fmt.Println("Error: --casesensitive cannot be used with --regex or --glob.")
			os.Exit(1)
		}
		if (regexFlag || globFlag) && fileType {
			fmt.Println("Error: --filetype cannot be used with --regex or --glob.")
			os.Exit(1)
		}
		if (regexFlag || globFlag) && extension {
			fmt.Println("Error: --extension cannot be used with --regex or --glob.")
			os.Exit(1)
		}

		options := search.FilterOptions{
			RegexPattern:    regexFlag,
			GlobPattern:     globFlag,
			CaseSensitive:   caseSensitive,
			ExcludePatterns: excludePatterns,
			FileType:        fileType,
			Extension:       extension,
		}

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

		results, err := search.Search(pattern, path, depth, options)

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
	rootCmd.Flags().StringSliceP("exclude", "x", []string{}, "Exclude files or directories matching a glob pattern")
	rootCmd.Flags().BoolP("filetype", "t", false, "Filter results by file type (e.g., file, dir, symlink)")
	rootCmd.Flags().BoolP("extension", "e", false, "Filter results by file extension")
	rootCmd.Flags().BoolP("casesensitive", "S", false, "Enable case-sensitive search")

	rootCmd.SetUsageTemplate(`Usage:
  gofs [Flag] <pattern> [pathname]

Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}
	`)
}
