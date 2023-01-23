//go:build windows

package dshow

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediaseeking
type IMediaSeeking interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-checkcapabilities
	CheckCapabilities(
		capabilities dshowco.SEEKING_CAPABILITIES) dshowco.SEEKING_CAPABILITIES

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-converttimeformat
	ConvertTimeFormat(targetFormat, sourceFormat *win.GUID, source int64) int64

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getavailable
	GetAvailable() (earliest, latest time.Duration)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getcapabilities
	GetCapabilities() dshowco.SEEKING_CAPABILITIES

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getcurrentposition
	GetCurrentPosition() time.Duration

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getduration
	GetDuration() time.Duration

	// Returns current and stop positions.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getpositions
	GetPositions() (current, stop time.Duration)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getpreroll
	GetPreroll() time.Duration

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getrate
	GetRate() float64

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-getstopposition
	GetStopPosition() time.Duration

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-gettimeformat
	GetTimeFormat() *win.GUID

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-isformatsupported
	IsFormatSupported(format *win.GUID) bool

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-isusingtimeformat
	IsUsingTimeFormat(format *win.GUID) bool

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-querypreferredformat
	QueryPreferredFormat() *win.GUID

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-setpositions
	SetPositions(
		current time.Duration, currentFlags dshowco.SEEKING_FLAGS,
		stop time.Duration, stopFlags dshowco.SEEKING_FLAGS) error

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-setrate
	SetRate(rate float64) error

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-settimeformat
	SetTimeFormat(format *win.GUID) error
}

type _IMediaSeeking struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMediaSeeking.Release().
//
// Example:
//
//	var gb dshow.IGraphBuilder // initialized somewhere
//
//	ms := dshow.NewIMediaSeeking(
//		gb.QueryInterface(dshowco.IID_IMediaSeeking),
//	)
//	defer ms.Release()
func NewIMediaSeeking(base com.IUnknown) IMediaSeeking {
	return &_IMediaSeeking{IUnknown: base}
}

func (me *_IMediaSeeking) CheckCapabilities(
	capabilities dshowco.SEEKING_CAPABILITIES) dshowco.SEEKING_CAPABILITIES {

	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).CheckCapabilities,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&capabilities)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return capabilities
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) ConvertTimeFormat(
	targetFormat, sourceFormat *win.GUID, source int64) int64 {

	var target int64
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).ConvertTimeFormat,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&target)), uintptr(unsafe.Pointer(targetFormat)),
		uintptr(unsafe.Pointer(&source)), uintptr(unsafe.Pointer(sourceFormat)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return target
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetAvailable() (earliest, latest time.Duration) {
	var iEarliest, iLatest int64
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetAvailable,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&iEarliest)), uintptr(unsafe.Pointer(&iLatest)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		earliest, latest = util.Nano100ToDuration(iEarliest), util.Nano100ToDuration(iLatest)
		return
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetCapabilities() dshowco.SEEKING_CAPABILITIES {
	var capabilities dshowco.SEEKING_CAPABILITIES
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetCapabilities,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&capabilities)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return capabilities
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetCurrentPosition() time.Duration {
	var pos int64
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetCurrentPosition,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pos)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return util.Nano100ToDuration(pos)
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetDuration() time.Duration {
	var duration int64
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetDuration,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&duration)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return util.Nano100ToDuration(duration)
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetPositions() (current, stop time.Duration) {
	var iCurrent, iStop int64
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetPositions,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&iCurrent)), uintptr(unsafe.Pointer(&iStop)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		current, stop = util.Nano100ToDuration(iCurrent), util.Nano100ToDuration(iStop)
		return
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetPreroll() time.Duration {
	var preRoll int64
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetPreroll,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&preRoll)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return util.Nano100ToDuration(preRoll)
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetRate() float64 {
	var rate float64
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetStopPosition,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&rate)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return rate
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetStopPosition() time.Duration {
	var stop int64
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetStopPosition,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&stop)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return util.Nano100ToDuration(stop)
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) GetTimeFormat() *win.GUID {
	format := &win.GUID{}
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).GetTimeFormat,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(format)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return format
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) IsFormatSupported(format *win.GUID) bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).IsFormatSupported,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(format)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) IsUsingTimeFormat(format *win.GUID) bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).IsUsingTimeFormat,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(format)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) QueryPreferredFormat() *win.GUID {
	format := win.GUID{}
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).QueryPreferredFormat,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&format)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return &format
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) SetPositions(
	current time.Duration, currentFlags dshowco.SEEKING_FLAGS,
	stop time.Duration, stopFlags dshowco.SEEKING_FLAGS) error {

	iCurrent, iStop := util.DurationToNano100(current), util.DurationToNano100(stop)
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).SetPositions,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&iCurrent)), uintptr(currentFlags),
		uintptr(unsafe.Pointer(&iStop)), uintptr(stopFlags))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}

func (me *_IMediaSeeking) SetRate(rate float64) error {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).SetRate,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(rate))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else if hr == errco.E_INVALIDARG || hr == errco.VFW_E_UNSUPPORTED_AUDIO {
		return hr
	} else {
		panic(hr)
	}
}

func (me *_IMediaSeeking) SetTimeFormat(format *win.GUID) error {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaSeeking)(unsafe.Pointer(*me.Ptr())).SetTimeFormat,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(format)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return nil
	} else {
		return hr
	}
}
