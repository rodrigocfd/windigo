package dshow

import (
	"time"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
)

// Converts time.Duration to 100 nanoseconds.
func _DurationTo100Nanosec(duration time.Duration) int64 {
	return int64(duration) * 10_000 / int64(time.Millisecond)
}

// Converts 100 nanoseconds to time.Duration.
func _Nanosec100ToDuration(nanosec100 int64) time.Duration {
	return time.Duration(nanosec100 / 10_000 * int64(time.Millisecond))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/ns-strmif-am_media_type
type AM_MEDIA_TYPE struct {
	Majortype            win.GUID
	Subtype              win.GUID
	bFixedSizeSamples    win.BOOL
	bTemporalCompression win.BOOL
	LSampleSize          uint32
	Formattype           win.GUID
	IUnknown             win.IUnknown
	CbFormat             uint32
	PbFormat             *byte
}

func (amt *AM_MEDIA_TYPE) BFixedSizeSamples() bool { return amt.bFixedSizeSamples != 0 }
func (amt *AM_MEDIA_TYPE) SetBFixedSizeSamples(val bool) {
	amt.bFixedSizeSamples = win.BOOL(util.BoolToUintptr(val))
}

func (amt *AM_MEDIA_TYPE) BTemporalCompression() bool { return amt.bTemporalCompression != 0 }
func (amt *AM_MEDIA_TYPE) SetBTemporalCompression(val bool) {
	amt.bTemporalCompression = win.BOOL(util.BoolToUintptr(val))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/ns-strmif-filter_info
type FILTER_INFO struct {
	AchName [128]uint16
	PGraph  IFilterGraph
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/ns-evr-mfvideonormalizedrect
type MFVideoNormalizedRect struct {
	Left, Top, Right, Bottom float32
}
