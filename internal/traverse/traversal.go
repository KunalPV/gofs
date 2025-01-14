package traverse

import (
	"context"
	"fmt"
	"gofs/utils"
	"sync"
)

// TraverseAndValidate performs directory traversal and validates the pathname.
// Returns validated paths and an error if the pathname is invalid.
func TraverseAndValidate(root string, pathname string, depth, maxThreads int) ([]string, error) {
	// Validate depth
	validDepth, err := utils.ValidateDepth(depth)
	if err != nil {
		return nil, err
	}

	// Validate maxThreads
	validThreads, err := utils.ValidateMaxThreads(maxThreads)
	if err != nil {
		return nil, err
	}

	// Check for hidden and ignore flags
	hidden := utils.CheckHiddenFlag()
	ignore := utils.CheckIgnoreFlag()

	// Create a context to manage cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel for traversal results
	results := make(chan string, maxThreads)

	// Perform traversal
	var allPaths []string
	var wg sync.WaitGroup

	// Collect results from traversal
	wg.Add(1)
	go func() {
		defer wg.Done()
		for path := range results {
			allPaths = append(allPaths, path)
		}
	}()

	// Call the traversal logic
	if err := TraverseAndStream(ctx, root, validDepth, results, validThreads, hidden, ignore); err != nil {
		cancel()
		wg.Wait()
		return nil, fmt.Errorf("error during traversal: %v", err)
	}

	// Wait for collection to finish
	wg.Wait()

	// Validate the pathname
	validPaths, found := utils.ValidatePathname(allPaths, pathname)
	if !found {
		return nil, fmt.Errorf("pathname doesn't exist in the current directory: %s", pathname)
	}

	return validPaths, nil
}
