//go:build windows

package dshow

import (
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
)

// [AM_MEDIA_TYPE] struct.
//
// [AM_MEDIA_TYPE]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/ns-strmif-am_media_type
type AM_MEDIA_TYPE struct {
	Majortype            win.GUID
	Subtype              win.GUID
	bFixedSizeSamples    int32 // BOOL
	bTemporalCompression int32 // BOOL
	LSampleSize          uint32
	Formattype           win.GUID
	IUnknown             com.IUnknown
	CbFormat             uint32
	PbFormat             *byte
}

func (mt *AM_MEDIA_TYPE) BFixedSizeSamples() bool       { return mt.bFixedSizeSamples != 0 }
func (mt *AM_MEDIA_TYPE) SetBFixedSizeSamples(val bool) { mt.bFixedSizeSamples = util.BoolToInt32(val) }

func (mt *AM_MEDIA_TYPE) BTemporalCompression() bool { return mt.bTemporalCompression != 0 }
func (mt *AM_MEDIA_TYPE) SetBTemporalCompression(val bool) {
	mt.bTemporalCompression = util.BoolToInt32(val)
}

// [FILTER_INFO] struct.
//
// [FILTER_INFO]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/ns-strmif-filter_info
type FILTER_INFO struct {
	achName [128]uint16
	PGraph  IFilterGraph
}

func (fi *FILTER_INFO) AchName() string { return win.Str.FromNativeSlice(fi.achName[:]) }
func (fi *FILTER_INFO) SetAchName(val string) {
	copy(fi.achName[:], win.Str.ToNativeSlice(win.Str.Substr(val, 0, len(fi.achName)-1)))
}

// [MFVideoNormalizedRect] struct.
//
// [MFVideoNormalizedRect]: https://learn.microsoft.com/en-us/windows/win32/api/evr/ns-evr-mfvideonormalizedrect
type MFVideoNormalizedRect struct {
	Left, Top, Right, Bottom float32
}
