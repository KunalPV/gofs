package cli

import (
	"errors"
	"runtime"

	"github.com/spf13/cobra"
)

type Config struct {
	Root          string
	Pattern       string
	Pathname      string
	Depth         int
	MaxThreads    int
	CaseSensitive bool
	IncludeHidden bool
	IncludeIgnore bool
	GlobPattern   string
	FilterOptions map[string]interface{} // Holds filter-related options
	FormatOptions map[string]interface{} // Holds format-related options
}

// DefineFlags adds flags to the root command
func DefineFlags(cmd *cobra.Command) {
	// Add standard flags
	cmd.Flags().BoolP("help", "h", false, "Display help for gofs")
	cmd.Flags().BoolP("version", "v", false, "Display the version of gofs")

	// Search flag
	cmd.Flags().StringP("glob", "g", "", "Search using a glob pattern (default: empty string)")

	// Traverse flags
	cmd.Flags().IntP("max-depth", "d", -1, "Limit search to a specific directory depth (-1 for no limit)")
	cmd.Flags().IntP("max-threads", "T", runtime.NumCPU(), "Set the maximum number of parallel threads for traversal")
	cmd.Flags().BoolP("case-sensitive", "S", false, "Perform case-sensitive searches")
	cmd.Flags().BoolP("hidden", "H", false, "Include hidden files in the search")
	cmd.Flags().BoolP("ignore", "I", false, "Include .*ignore files like .gitignore")

	// Filter flags
	cmd.Flags().StringP("extension", "e", "", "Filter results by file extensions")
	cmd.Flags().StringP("file-type", "t", "", "Filter results by file type (file, dir, symlink)")
	cmd.Flags().StringP("exclude", "x", "", "Exclude files/directories matching a glob pattern")

	// Format flags
	cmd.Flags().BoolP("absolute-path", "A", false, "Display resuults as absolute paths")
	cmd.Flags().BoolP("long-list", "l", false, "Display results in long list format")
	cmd.Flags().BoolP("hyper-link", "L", false, "Display results as hyperlinks")
}

// ParseFlags parses the flags and returns a Config struct
func ParseFlags(cmd *cobra.Command, args []string) Config {
	var pattern string
	globPattern, _ := cmd.Flags().GetString("glob")

	if len(args) > 0 {
		pattern = args[0]
	} else if globPattern != "" {
		pattern = globPattern // Use the glob pattern if no positional pattern argument is provided
	} else {
		pattern = "." // If neither pattern nor glob is specified
	}

	pathname := "."
	if len(args) > 1 {
		pathname = args[1]
	}

	depth, _ := cmd.Flags().GetInt("max-depth")
	maxThreads, _ := cmd.Flags().GetInt("max-threads")
	caseSensitive, _ := cmd.Flags().GetBool("case-sensitive")
	includeHidden, _ := cmd.Flags().GetBool("hidden")
	includeIgnore, _ := cmd.Flags().GetBool("ignore")
	extension, _ := cmd.Flags().GetString("extension")
	fileType, _ := cmd.Flags().GetString("file-type")
	exclude, _ := cmd.Flags().GetString("exclude")
	absolutePath, _ := cmd.Flags().GetBool("absolute-path")
	longList, _ := cmd.Flags().GetBool("long-list")
	hyperlink, _ := cmd.Flags().GetBool("hyper-link")

	// Construct FilterOptions as a map
	filterOptions := map[string]interface{}{
		"Extension": extension,
		"FileType":  fileType,
		"Exclude":   exclude,
		"AbsPath":   absolutePath,
	}

	formatOptions := map[string]interface{}{
		"AbsolutePath": absolutePath,
		"LongList":     longList,
		"Hyperlink":    hyperlink,
	}

	return Config{
		Root:          ".",
		Pattern:       pattern,
		Pathname:      pathname,
		Depth:         depth,
		MaxThreads:    maxThreads,
		CaseSensitive: caseSensitive,
		IncludeHidden: includeHidden,
		IncludeIgnore: includeIgnore,
		GlobPattern:   globPattern,
		FilterOptions: filterOptions,
		FormatOptions: formatOptions,
	}
}

// ValidateFlags ensures valid flag combinations
func ValidateFlags(cmd *cobra.Command, args []string) error {
	if len(args) > 2 {
		return errors.New("too many arguments; expected <pattern> and [pathname]")
	}

	return nil
}
