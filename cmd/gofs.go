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
	dir := flag.String("dir", ".", "Directory to search (default: current directory)")

	// Handle long flags manually for now - Implement cobra later
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
		fmt.Println("Usage: gofs [options] <directory> <filename>")
		fmt.Println("\nOptions:")
		fmt.Println("  -h, --help       Show help message")
		fmt.Println("  -v, --version    Show version of the utility")
		fmt.Println("  --dir            Directory to search (default: current directory)")
		fmt.Println("\nPositional Arguments:")
		fmt.Println("  <directory>      Directory to search (overrides --dir)")
		fmt.Println("  <filename>       Pattern to search for (required)")
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

	// Handle no arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: <filename> is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Extract directory and filename
	filename := args[len(args)-1] // last arg is the filename
	if len(args) > 1 {
		*dir = args[len(args)-2] // last second arg is the directory
	}

	results, err := search.Search(*dir, filename)
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
