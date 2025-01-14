package search

import (
	"fmt"
	"gofs/utils"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

// SearchWithThreads performs parallel search on traversal results.
func SearchWithThreads(pattern string, traversalResults []string, validThreads int, isGlob bool) ([]string, error) {

	// Compile regex if the pattern is not a glob
	var re *regexp.Regexp
	var err error

	if !isGlob {
		re, err = regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid regex pattern: %v", err)
		}
	} else if !utils.IsValidGlob(pattern) {
		return nil, fmt.Errorf("invalid glob pattern: %s", pattern)
	}

	// Channels for parallel processing
	workChan := make(chan string, len(traversalResults))
	resultsChan := make(chan string, len(traversalResults))
	errChan := make(chan error, 1)

	// Populate the work channel
	go func() {
		defer close(workChan)
		for _, file := range traversalResults {
			workChan <- file
		}
	}()

	// Worker function
	var wg sync.WaitGroup
	wg.Add(validThreads)

	for i := 0; i < validThreads; i++ {
		go func() {
			defer wg.Done()
			for file := range workChan {
				if matchFileOrDir(file, pattern, re, isGlob) {
					resultsChan <- file
				}
			}
		}()
	}

	// Close resultsChan after workers finish
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect results with de-duplication
	resultsSet := make(map[string]struct{})
	var results []string
	for result := range resultsChan {
		if _, exists := resultsSet[result]; !exists {
			resultsSet[result] = struct{}{}
			results = append(results, result)
		}
	}

	// Check for errors
	select {
	case err := <-errChan:
		return nil, err
	default:
	}

	return results, nil
}

// matchFileOrDir checks if a file or directory matches the pattern.
func matchFileOrDir(file string, pattern string, re *regexp.Regexp, isGlob bool) bool {

	if stat, err := os.Stat(file); err == nil {
		if stat.IsDir() {
			// Match directory name
			if isGlob {
				matched, _ := filepath.Match(pattern, filepath.Base(file))
				return matched
			}
			return re.MatchString(filepath.Base(file))
		}

		// Match file name
		if isGlob {
			matched, _ := filepath.Match(pattern, filepath.Base(file))
			return matched
		} else {
			return re.MatchString(file) || re.MatchString(filepath.Base(file))
		}
	}
	return false
}
