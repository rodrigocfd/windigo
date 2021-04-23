package dshow

import (
	"time"

	"github.com/rodrigocfd/windigo/win"
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/ns-strmif-am_media_type
type AM_MEDIA_TYPE struct {
	Majortype            win.GUID
	Subtype              win.GUID
	BFixedSizeSamples    win.BOOL
	BTemporalCompression win.BOOL
	LSampleSize          uint32
	Formattype           win.GUID
	IUnknown             win.IUnknown
	CbFormat             uint32
	PbFormat             *byte
}
