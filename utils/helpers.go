// utils/helpers.go
package utils

import (
	"fmt"
	"math"
)

func Truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}

func FormatTime(seconds float64) string {
	minutes := math.Floor(seconds / 60)
	remainingSeconds := math.Round(math.Mod(seconds, 60))
	return fmt.Sprintf("%02d:%02d", int(minutes), int(remainingSeconds))
}
