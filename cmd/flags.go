package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func defineFlags() {
	rootCmd.PersistentFlags().StringP("pathname", "p", ".", "Pathname to search (default: current directory)")

	rootCmd.Flags().BoolP("version", "v", false, "Display the version of the utility")

	rootCmd.Flags().BoolP("regex", "r", false, "Use regex pattern for searching")
	rootCmd.Flags().BoolP("glob", "g", false, "Use glob pattern for searching")

	rootCmd.Flags().IntP("max-depth", "d", -1, "Limit the depth of directory traversal (-1 for unlimited depth)")
	rootCmd.Flags().StringSliceP("exclude", "x", []string{}, "Exclude files or directories matching a glob pattern")

	rootCmd.Flags().BoolP("file-type", "t", false, "Filter results by file type (e.g., file, dir, symlink)")
	rootCmd.Flags().BoolP("extension", "e", false, "Filter results by file extension")
	rootCmd.Flags().BoolP("case-sensitive", "S", false, "Enable case-sensitive search")
	rootCmd.Flags().BoolP("abs-path", "A", false, "Show absolute paths in the results")

	rootCmd.Flags().IntP("threads", "T", runtime.NumCPU(), "Set the number of threads for parallel execution (default: number of CPU cores)")

	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	rootCmd.SetUsageTemplate(`Usage:
  gofs [Flag] <pattern> [pathname]

Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}
	`)
}

// validateFlags checks for conflicts and ensures proper flag usage
func validateFlags(cmd *cobra.Command, args []string) error {
	regexFlag, _ := cmd.Flags().GetBool("regex")
	globFlag, _ := cmd.Flags().GetBool("glob")
	fileType, _ := cmd.Flags().GetBool("file-type")
	extension, _ := cmd.Flags().GetBool("extension")
	caseSensitive, _ := cmd.Flags().GetBool("case-sensitive")

	if regexFlag && globFlag {
		return fmt.Errorf("--regex and --glob cannot be used together")
	}
	if fileType && extension {
		return fmt.Errorf("--file-type and --extension cannot be used together")
	}
	if (regexFlag || globFlag) && (caseSensitive || fileType || extension) {
		return fmt.Errorf("--regex and --glob are incompatible with --case-sensitive, --file-type, and --extension")
	}
	return nil
}
