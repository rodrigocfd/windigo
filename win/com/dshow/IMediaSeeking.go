package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
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
func (me *IMediaSeeking) CheckCapabilities(capabilities co.AM_SEEKING) co.AM_SEEKING {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).CheckCapabilities, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&capabilities)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return capabilities
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-converttimeformat
func (me *IMediaSeeking) ConvertTimeFormat(
	targetFormat, sourceFormat *win.GUID, source int64) int64 {

	target := int64(0)
	ret, _, _ := syscall.Syscall6(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).ConvertTimeFormat, 5,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&target)), uintptr(unsafe.Pointer(targetFormat)),
		uintptr(unsafe.Pointer(&source)), uintptr(unsafe.Pointer(sourceFormat)),
		0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return target
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getavailable
func (me *IMediaSeeking) GetAvailable() (earliest, latest int64) {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetAvailable, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&earliest)), uintptr(unsafe.Pointer(&latest)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getcapabilities
func (me *IMediaSeeking) GetCapabilities() co.AM_SEEKING {
	capabilities := co.AM_SEEKING(0)
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetCapabilities, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&capabilities)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return capabilities
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getcurrentposition
func (me *IMediaSeeking) GetCurrentPosition() int64 {
	pos := int64(0)
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetCurrentPosition, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pos)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return pos
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getduration
func (me *IMediaSeeking) GetDuration() int64 {
	duration := int64(0)
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetDuration, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&duration)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return duration
}

// Returns current and stop positions.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getpositions
func (me *IMediaSeeking) GetPositions() (current, stop int64) {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetPositions, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&current)), uintptr(unsafe.Pointer(&stop)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getpreroll
func (me *IMediaSeeking) GetPreroll() int64 {
	preroll := int64(0)
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetPreroll, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&preroll)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return preroll
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getrate
func (me *IMediaSeeking) GetRate() float64 {
	rate := float64(0)
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetStopPosition, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&rate)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return rate
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getstopposition
func (me *IMediaSeeking) GetStopPosition() int64 {
	pos := int64(0)
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetStopPosition, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&pos)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return pos
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-gettimeformat
func (me *IMediaSeeking) GetTimeFormat() *win.GUID {
	format := win.GUID{}
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).GetTimeFormat, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&format)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return &format
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-isformatsupported
func (me *IMediaSeeking) IsFormatSupported(format *win.GUID) bool {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).IsFormatSupported, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(format)), 0)

	lerr := err.ERROR(ret)
	if lerr == err.S_OK {
		return true
	} else if lerr == err.S_FALSE {
		return false
	}
	panic(lerr)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-isusingtimeformat
func (me *IMediaSeeking) IsUsingTimeFormat(format *win.GUID) bool {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).IsUsingTimeFormat, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(format)), 0)

	lerr := err.ERROR(ret)
	if lerr == err.S_OK {
		return true
	} else if lerr == err.S_FALSE {
		return false
	}
	panic(lerr)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-querypreferredformat
func (me *IMediaSeeking) QueryPreferredFormat() *win.GUID {
	format := win.GUID{}
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).QueryPreferredFormat, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&format)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return &format
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-setpositions
func (me *IMediaSeeking) SetPositions(
	current int64, currentFlags co.AM_SEEKING,
	stop int64, stopFlags co.AM_SEEKING) error {

	ret, _, _ := syscall.Syscall6(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).SetPositions, 5,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&current)), uintptr(currentFlags),
		uintptr(unsafe.Pointer(&stop)), uintptr(stopFlags), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return lerr
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-setrate
func (me *IMediaSeeking) SetRate(rate float64) error {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).SetRate, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(rate), 0)

	if lerr := err.ERROR(ret); lerr == err.E_INVALIDARG || lerr == err.VFW_E_UNSUPPORTED_AUDIO {
		return lerr
	} else if lerr != err.S_OK {
		panic(lerr)
	}

	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-settimeformat
func (me *IMediaSeeking) SetTimeFormat(format *win.GUID) error {
	ret, _, _ := syscall.Syscall(
		(*_IMediaSeekingVtbl)(unsafe.Pointer(*me.Ppv)).SetTimeFormat, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(format)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return lerr
	}
	return nil
}
