package traverse

import (
	"context"
	"gofs/utils"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
)

// TraverseAndStream traverses the directory tree starting from root, up to a specified depth.
// Streams files and directories to a results channel for further processing.
func TraverseAndStream(ctx context.Context, root string, depth int, results chan<- string, maxThreads int, hidden bool, ignore bool) error {

	var wg sync.WaitGroup
	workChan := make(chan string, maxThreads)
	errChan := make(chan error, 1)

	// Worker function to process directories
	wg.Add(maxThreads)
	for i := 0; i < maxThreads; i++ {
		go func(workerID int) {
			defer wg.Done()
			for dir := range workChan {

				// Stream the directory itself
				// Stream the directory itself, but skip "."
				if dir != "." {
					select {
					case results <- dir:
					case <-ctx.Done():
						return
					}
				}

				err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
					if err != nil {
						// Send error to errChan and stop further processing
						select {
						case errChan <- err:
						default:
						}
						return filepath.SkipDir
					}

					// Skip hidden files if hidden flag is not set
					if !hidden && utils.IsHidden(path) {
						if d.IsDir() {
							return filepath.SkipDir
						}
						return nil
					}

					// Skip ignored files (implement .ignore file parsing logic)
					if !ignore && utils.IsIgnored(path) {
						return nil
					}

					// Skip the root directory itself (already streamed)
					if path == root {
						return nil
					}

					// Calculate depth
					relativePath, err := filepath.Rel(root, path)
					if err != nil {
						return nil
					}
					currentDepth := strings.Count(relativePath, string(filepath.Separator))

					// Skip entries beyond the specified depth
					if depth != -1 && currentDepth > depth {
						if d.IsDir() {
							return filepath.SkipDir
						}
						return nil
					}

					// Stream the result immediately
					select {
					case results <- path:
					case <-ctx.Done(): // Stop if context is canceled
						return context.Canceled
					}

					return nil
				})
				if err != nil {
					select {
					case errChan <- err:
					default:
					}
				}
			}
		}(i)
	}

	// Seed the work channel with the root directory
	go func() {
		defer close(workChan)
		workChan <- root
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
		close(errChan)
	}()

	// Check for errors
	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return context.Canceled
	}

	return nil
}
