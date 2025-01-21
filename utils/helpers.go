// utils/helpers.go
package utils

import (
	"fmt"
	"math"
)

// Truncate a string to a maximum length
func Truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}

// Format time in seconds to a human-readable format
func FormatTime(seconds float64) string {
	minutes := math.Floor(seconds / 60)
	remainingSeconds := math.Round(math.Mod(seconds, 60))
	return fmt.Sprintf("%02d:%02d", int(minutes), int(remainingSeconds))
}
