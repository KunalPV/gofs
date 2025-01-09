package parallel

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

// ExecuteWithThreads executes tasks in parallel, distributing the workload across the specified number of threads.
// - task: A function that generates the list of work items.
// - process: A function that processes each work item.
// - maxThreads: The maximum number of threads to use for processing.
func ExecuteWithThreads(task func() ([]string, error), process func(string) (string, error), maxThreads int) ([]string, error) {
	// Fetch work items (e.g., files or paths to process)
	if maxThreads <= 0 {
		maxThreads = runtime.NumCPU() // Use all available CPUs by default
	}

	items, err := task()
	if err != nil {
		return nil, fmt.Errorf("failed to generate work items: %w", err)
	}

	// De-duplicate work items
	workSet := make(map[string]struct{})
	for _, item := range items {
		workSet[item] = struct{}{}
	}

	// Channels for work distribution, results, and errors
	workChan := make(chan string, len(workSet))
	resultsChan := make(chan string, len(workSet))
	errChan := make(chan error, 1)

	// Populate the work channel
	go func() {
		defer close(workChan)
		for item := range workSet {
			workChan <- item
		}
	}()

	// Worker function
	var wg sync.WaitGroup
	wg.Add(maxThreads)

	for i := 0; i < maxThreads; i++ {
		go func() {
			defer wg.Done()
			for item := range workChan {
				result, err := process(item)
				if err != nil {
					// Non-blocking send to errChan to report the first error
					select {
					case errChan <- err:
					default: // Ignore subsequent errors
					}
					return
				}

				// Split multi-line results into individual entries
				if result != "" {
					for _, line := range strings.Split(result, "\n") {
						trimmedLine := strings.TrimSpace(line)
						if trimmedLine != "" {
							resultsChan <- trimmedLine
						}
					}
				}
			}
		}()
	}

	// Close the results channel after workers finish
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

	// Check for errors from errChan
	select {
	case err := <-errChan:
		return nil, err
	default:
	}

	return results, nil
}
