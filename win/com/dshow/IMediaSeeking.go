package dshow

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IMediaSeekingVtbl struct {
	win.IUnknownVtbl
	GetCapabilities      uintptr
	CheckCapabilities    uintptr
	IsFormatSupported    uintptr
	QueryPreferredFormat uintptr
	GetTimeFormat        uintptr
	IsUsingTimeFormat    uintptr
	SetTimeFormat        uintptr
	GetDuration          uintptr
	GetStopPosition      uintptr
	GetCurrentPosition   uintptr
	ConvertTimeFormat    uintptr
	SetPositions         uintptr
	GetPositions         uintptr
	GetAvailable         uintptr
	SetRate              uintptr
	GetRate              uintptr
	GetPreroll           uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediaseeking
type IMediaSeeking struct {
	win.IUnknown // Base IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-checkcapabilities
func (me *IMediaSeeking) CheckCapabilities(
	capabilities dshowco.SEEKING_CAPABILITIES) dshowco.SEEKING_CAPABILITIES {

	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).CheckCapabilities, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&capabilities)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return capabilities
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-converttimeformat
func (me *IMediaSeeking) ConvertTimeFormat(
	targetFormat, sourceFormat *win.GUID, source int64) int64 {

	var target int64
	ret, _, _ := syscall.Syscall6(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).ConvertTimeFormat, 5,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&target)), uintptr(unsafe.Pointer(targetFormat)),
		uintptr(unsafe.Pointer(&source)), uintptr(unsafe.Pointer(sourceFormat)),
		0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return target
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getavailable
func (me *IMediaSeeking) GetAvailable() (earliest, latest time.Duration) {
	var iEarliest, iLatest int64
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetAvailable, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&iEarliest)), uintptr(unsafe.Pointer(&iLatest)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		earliest, latest = _Nanosec100ToDuration(iEarliest), _Nanosec100ToDuration(iLatest)
		return
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getcapabilities
func (me *IMediaSeeking) GetCapabilities() dshowco.SEEKING_CAPABILITIES {
	var capabilities dshowco.SEEKING_CAPABILITIES
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetCapabilities, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&capabilities)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return capabilities
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getcurrentposition
func (me *IMediaSeeking) GetCurrentPosition() time.Duration {
	var pos int64
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetCurrentPosition, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pos)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return _Nanosec100ToDuration(pos)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getduration
func (me *IMediaSeeking) GetDuration() time.Duration {
	var duration int64
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetDuration, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&duration)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return _Nanosec100ToDuration(duration)
	} else {
		panic(hr)
	}
}

// Returns current and stop positions.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getpositions
func (me *IMediaSeeking) GetPositions() (current, stop time.Duration) {
	var iCurrent, iStop int64
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetPositions, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&iCurrent)), uintptr(unsafe.Pointer(&iStop)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		current, stop = _Nanosec100ToDuration(iCurrent), _Nanosec100ToDuration(iStop)
		return
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getpreroll
func (me *IMediaSeeking) GetPreroll() time.Duration {
	var preroll int64
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetPreroll, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&preroll)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return _Nanosec100ToDuration(preroll)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getrate
func (me *IMediaSeeking) GetRate() float64 {
	var rate float64
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetStopPosition, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&rate)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return rate
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getstopposition
func (me *IMediaSeeking) GetStopPosition() time.Duration {
	var stop int64
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetStopPosition, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&stop)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return _Nanosec100ToDuration(stop)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-gettimeformat
func (me *IMediaSeeking) GetTimeFormat() *win.GUID {
	var format win.GUID
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetTimeFormat, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&format)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return &format
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-isformatsupported
func (me *IMediaSeeking) IsFormatSupported(format *win.GUID) bool {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).IsFormatSupported, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(format)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-isusingtimeformat
func (me *IMediaSeeking) IsUsingTimeFormat(format *win.GUID) bool {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).IsUsingTimeFormat, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(format)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-querypreferredformat
func (me *IMediaSeeking) QueryPreferredFormat() *win.GUID {
	var format win.GUID
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).QueryPreferredFormat, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&format)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return &format
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-setpositions
func (me *IMediaSeeking) SetPositions(
	current time.Duration, currentFlags dshowco.SEEKING_FLAGS,
	stop time.Duration, stopFlags dshowco.SEEKING_FLAGS) error {

	iCurrent, iStop := _DurationTo100Nanosec(current), _DurationTo100Nanosec(stop)
	ret, _, _ := syscall.Syscall6(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).SetPositions, 5,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&iCurrent)), uintptr(currentFlags),
		uintptr(unsafe.Pointer(&iStop)), uintptr(stopFlags), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-setrate
func (me *IMediaSeeking) SetRate(rate float64) error {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).SetRate, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(rate), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else if hr == errco.E_INVALIDARG || hr == errco.VFW_E_UNSUPPORTED_AUDIO {
		return hr
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-settimeformat
func (me *IMediaSeeking) SetTimeFormat(format *win.GUID) error {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).SetTimeFormat, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(format)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}
