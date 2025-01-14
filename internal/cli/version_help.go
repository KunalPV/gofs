package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

// PrioritizeHelpAndVersion ensures -h and -v are handled correctly
func PrioritizeHelpAndVersion(cmd *cobra.Command, args []string) error {
	// Check if the -h or -v flags are present
	showHelp, _ := cmd.Flags().GetBool("help")
	showVersion, _ := cmd.Flags().GetBool("version")

	if showHelp {
		cmd.Help() // Display help and exit
		os.Exit(0)
	}

	if showVersion {
		fmt.Println("gofs version", version) // Display version and exit
		os.Exit(0)
	}

	return nil
}
