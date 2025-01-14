package utils

import (
	"fmt"
)

// ValidateDepth checks if the depth is valid and converts -1 to user-friendly terminology (0 or 1-based).
func ValidateDepth(depth int) (int, error) {
	if depth < -1 {
		return -1, fmt.Errorf("invalid depth: %d, must be -1 (unlimited) or a non-negative value", depth)
	}
	// Return depth as-is (-1 for unlimited) but 0/1+ for user-friendly usage.
	return depth, nil
}
