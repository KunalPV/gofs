package search

import (
	"fmt"
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
				var matched bool
				var err error

				// Check if it's a directory
				if stat, statErr := os.Stat(file); statErr == nil && stat.IsDir() {
					if isGlob {
						// Match the directory name using the glob pattern
						matched, err = filepath.Match(pattern, filepath.Base(file))
						if err != nil {
							select {
							case errChan <- fmt.Errorf("error matching directory with glob pattern: %v", err):
							default:
							}
							continue
						}
					} else {
						// Match the directory name using the regex
						matched = re.MatchString(filepath.Base(file))
					}

					if matched {
						// Add the directory to results with a trailing slash
						resultsChan <- file
						continue
					}
				}

				// Match files
				if isGlob {
					// Match the file name using the glob pattern
					matched, err = filepath.Match(pattern, filepath.Base(file))
					if err != nil {
						select {
						case errChan <- fmt.Errorf("error matching glob pattern: %v", err):
						default:
						}
						continue
					}
				} else {
					// Match the file name using the regex
					matched = re.MatchString(file) || re.MatchString(filepath.Base(file))
				}

				if matched {
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
