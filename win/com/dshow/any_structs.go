package dshow

import (
	"time"
)

// Converts 100 nanoseconds to time.Duration.
func _Nanosec100ToDuration(nanosec100 int64) time.Duration {
	return time.Duration(nanosec100 / 10000 * int64(time.Millisecond))
}

// Converts time.Duration to 100 nanoseconds.
func _DurationTo100Nanosec(duration time.Duration) int64 {
	return int64(duration) * 10000 / int64(time.Millisecond)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/ns-evr-mfvideonormalizedrect
type MFVideoNormalizedRect struct {
	Left, Top, Right, Bottom float32
}
