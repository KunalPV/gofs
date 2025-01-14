package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ValidateCommand ensures valid pattern, pathname, and flag usage
func ValidateCommand(cmd *cobra.Command, args []string) error {
	// Step 1: Default arguments
	pattern := "."
	pathname := ""

	// Step 2: Parse arguments
	if len(args) > 0 {
		pattern = args[0]
	}
	if len(args) > 1 {
		pathname = args[1]
	}

	// Step 3: Validate pathname is not mistakenly provided in pattern
	if len(args) == 1 && strings.HasSuffix(pattern, "/") {
		return errors.New("pathname has been provided in the pattern parameter")
	}

	// Step 4: Validate pathname if provided
	if len(args) == 2 {
		// Manually check if the pathname exists
		if _, err := os.Stat(pathname); os.IsNotExist(err) {
			return fmt.Errorf("pathname doesn't exist in the current directory: %s", pathname)
		}
	}

	// Step 5: Set pattern and pathname in flags
	cmd.Flags().Set("pattern", pattern)
	cmd.Flags().Set("pathname", pathname)

	// Step 6: Validate flags (e.g., help/version logic)
	if err := validateFlags(cmd); err != nil {
		return err
	}

	return nil
}

// validateFlags checks if flags are used correctly
func validateFlags(cmd *cobra.Command) error {
	// Get the list of flags provided by the user
	flags := cmd.Flags()

	// Validate mutually exclusive flags, if any
	showHelp, _ := flags.GetBool("help")
	showVersion, _ := flags.GetBool("version")

	// Check the order of appearance
	args := os.Args[1:] // Exclude the program name
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			if showHelp {
				cmd.Help() // Display help and exit
				os.Exit(0)
			}
		} else if arg == "-v" || arg == "--version" {
			if showVersion {
				fmt.Println("gofs version 1.0.0") // Display version and exit
				os.Exit(0)
			}
		}
	}

	// Add additional flag-specific validations here if needed
	return nil
}
