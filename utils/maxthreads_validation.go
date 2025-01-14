package utils

import (
	"fmt"
	"runtime"
)

// ValidateMaxThreads checks if the number of threads is within a valid range (1 to max CPU cores).
func ValidateMaxThreads(maxThreads int) (int, error) {
	maxCores := runtime.NumCPU()
	if maxThreads < 1 || maxThreads > maxCores {
		return -1, fmt.Errorf("invalid maxThreads: %d, must be between 1 and %d", maxThreads, maxCores)
	}
	return maxThreads, nil
}
