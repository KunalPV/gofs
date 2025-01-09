package search

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
)

// Traverse lists files and directories in a given directory up to a specified depth.
// If depth is -1, it means unlimited depth. maxThreads controls parallelism.
func Traverse(root string, depth int, maxThreads int) ([]string, error) {
	if depth < -1 {
		return nil, fmt.Errorf("invalid depth: %d, depth must be -1 or a non-negative value", depth)
	}
	if maxThreads <= 0 {
		return nil, fmt.Errorf("maxThreads must be greater than 0")
	}

	var mu sync.Mutex
	var files []string
	var wg sync.WaitGroup

	workChan := make(chan string, maxThreads)
	errChan := make(chan error, 1)

	// Worker function
	wg.Add(maxThreads)
	for i := 0; i < maxThreads; i++ {
		go func() {
			defer wg.Done()
			for dir := range workChan {
				err := filepath.WalkDir(dir, func(currentPath string, d fs.DirEntry, err error) error {
					if err != nil {
						// Send the error and stop processing
						select {
						case errChan <- err:
						default:
						}
						return nil
					}

					// Skip the root directory itself
					if currentPath == root {
						return nil
					}

					// Calculate the relative path and depth
					relativePath, err := filepath.Rel(root, currentPath)
					if err != nil {
						select {
						case errChan <- fmt.Errorf("error calculating relative path for %s: %v", currentPath, err):
						default:
						}
						return nil
					}

					currentDepth := strings.Count(relativePath, string(filepath.Separator)) + 1

					// Skip entries beyond the specified depth
					if depth != -1 && currentDepth > depth {
						if d.IsDir() {
							return filepath.SkipDir // Skip entire directories beyond the depth
						}
						return nil
					}

					// Add the current path to the results
					mu.Lock()
					files = append(files, currentPath)
					mu.Unlock()
					return nil
				})

				if err != nil {
					select {
					case errChan <- err:
					default:
					}
				}
			}
		}()
	}

	// Seed the work channel with the root directory
	go func() {
		defer close(workChan)
		workChan <- root
	}()

	// Wait for all workers to finish
	wg.Wait()

	// Check for errors
	select {
	case err := <-errChan:
		return nil, err
	default:
	}

	return files, nil
}
