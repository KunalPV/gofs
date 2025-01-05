package main

import (
	"flag"
	"fmt"
	"gofs/internal/search"
	"os"
)

const version = "1.0.0"

func main() {
	help := flag.Bool("h", false, "Show help message")
	versionFlag := flag.Bool("v", false, "Show the version of the utility")
	path := flag.String("path", ".", "Pathname to search (default: current directory)")

	// Handle long flags manually for now - Implement COBRA later
	for i, arg := range os.Args {
		if arg == "--help" {
			os.Args[i] = "-h"
		}
		if arg == "--version" {
			os.Args[i] = "-v"
		}
	}

	// Override the default usage function
	flag.Usage = func() {
		fmt.Println("Usage: gofs [options] <pattern> [pathname]")
		fmt.Println("\nOptions:")
		fmt.Println("  -h, --help       Show help message")
		fmt.Println("  -v, --version    Show version of the utility")
		fmt.Println("\nPositional Arguments:")
		fmt.Println("  <pattern>       Pattern to search for (required)")
		fmt.Println("  [pathname]      Pathname to search (optional)")
	}

	flag.Parse()

	// Handle --help and -h
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Handle --version and -v
	if *versionFlag {
		fmt.Printf("gofs version %s\n", version)
		os.Exit(0)
	}

	// Handle no arguments: List all files and directories
	args := flag.Args()
	if len(args) == 0 {
		results, err := search.Traverse(*path, true)
		if err != nil {
			fmt.Printf("Error listing files: %v\n", err)
			os.Exit(1)
		}

		if len(results) == 0 {
			fmt.Println("No files found.")
		} else {
			fmt.Println("Files and directories:")
			for _, result := range results {
				fmt.Println(result)
			}
		}
		os.Exit(0)
	}

	// Extract directory and filename
	pattern := args[0] // First argument is always the filename
	if len(args) > 1 {
		*path = args[1] // Second argument, if provided, overrides --dir
	}

	results, err := search.Search(pattern, *path)
	if err != nil {
		fmt.Printf("Search failed: %v\n", err)
		os.Exit(1)
	}

	if len(results) == 0 {
		fmt.Println("No files found.")
	} else {
		fmt.Println("Found files.")
		for _, result := range results {
			fmt.Println(result)
		}
	}
}
