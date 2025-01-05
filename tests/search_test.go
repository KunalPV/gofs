package search_test

import (
	"gofs/internal/search"
	"os"
	"path/filepath"
	"testing"
)

func TestSearch(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	testdataDir := filepath.Join(currentDir, "../testdata")

	t.Run("Search for existing file", func(t *testing.T) {
		results, err := search.Search(testdataDir, "example.txt")

		if err != nil {
			t.Fatalf("Search failed: %v", err)
		}

		if len(results) != 1 {
			t.Fatalf("Expected 1 match, found %d", len(results))
		}
	})

	t.Run("Search for non-existent file", func(t *testing.T) {
		results, err := search.Search(testdataDir, "random.txt")

		if err != nil {
			t.Fatalf("Search failed: %v", err)
		}

		if len(results) != 1 {
			t.Fatalf("Expected 1 match, found %d", len(results))
		}
	})

	t.Run("Search for nested directory", func(t *testing.T) {
		results, err := search.Search(testdataDir, "example4.json")

		if err != nil {
			t.Fatalf("Search failed: %v", err)
		}

		if len(results) != 1 {
			t.Fatalf("Expected 1 match, fount %d", len(results))
		}
	})
}
