package wstr

import (
	"fmt"
)

// Formats a number of bytes into KB, MB, GB, TB, PB or EB, rounding to 2
// decimal places.
func FmtBytes(numBytes uint64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
		PB = 1024 * TB
		EB = 1024 * PB
	)

	switch {
	case numBytes < KB:
		return fmt.Sprintf("%d bytes", numBytes)
	case numBytes < MB:
		return fmt.Sprintf("%.2f KB", float64(numBytes)/KB)
	case numBytes < GB:
		return fmt.Sprintf("%.2f MB", float64(numBytes)/MB)
	case numBytes < TB:
		return fmt.Sprintf("%.2f GB", float64(numBytes)/GB)
	case numBytes < PB:
		return fmt.Sprintf("%.2f TB", float64(numBytes)/TB)
	case numBytes < EB:
		return fmt.Sprintf("%.2f PB", float64(numBytes)/PB)
	default:
		return fmt.Sprintf("%.2f EB", float64(numBytes)/EB)
	}
}
