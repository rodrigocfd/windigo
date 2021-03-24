package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
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
