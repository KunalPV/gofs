package utils

import (
	"fmt"
	"runtime"
)

// ValidateThreads ensures the thread count is within a valid range.
func ValidateThreads(maxThreads int) (int, error) {
	if maxThreads < 1 {
		return 0, fmt.Errorf("invalid thread count: %d. Defaulting to %d threads", maxThreads, runtime.NumCPU())
	}
	maxCores := runtime.NumCPU()
	if maxThreads > maxCores {
		return 0, fmt.Errorf("thread count exceeds available cores (%d). Defaulting to %d threads", maxCores, maxCores)
	}
	return maxThreads, nil
}

// GetValidatedThreads retrieves and validates the thread count from the command flags.
func GetValidatedThreads(threadFlag int) (int, error) {
	if threadFlag == 0 {
		threadFlag = runtime.NumCPU() // Default to the number of CPU cores
	}
	return ValidateThreads(threadFlag)
}
