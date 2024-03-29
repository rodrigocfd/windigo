package dshow

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IMediaFilterVtbl struct {
	win.IPersistVtbl
	Stop          uintptr
	Pause         uintptr
	Run           uintptr
	GetState      uintptr
	SetSyncSource uintptr
	GetSyncSource uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
type IMediaFilter struct {
	win.IPersist // Base IPersist > IUnknown.
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-getstate
func (me *IMediaFilter) GetState(msTimeout int) (dshowco.FILTER_STATE, error) {
	state := dshowco.FILTER_STATE(0)
	ret, _, _ := syscall.Syscall(
		(*_IMediaFilterVtbl)(unsafe.Pointer(*me.Ppv)).GetState, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(int32(msTimeout)), uintptr(unsafe.Pointer(&state)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return state, nil
	} else if hr == errco.VFW_S_STATE_INTERMEDIATE {
		return state, hr
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-pause
func (me *IMediaFilter) Pause() bool {
	ret, _, _ := syscall.Syscall(
		(*_IMediaFilterVtbl)(unsafe.Pointer(*me.Ppv)).Pause, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-run
func (me *IMediaFilter) Run(tStart time.Duration) bool {
	iStart := _DurationTo100Nanosec(tStart)
	ret, _, _ := syscall.Syscall(
		(*_IMediaFilterVtbl)(unsafe.Pointer(*me.Ppv)).Run, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(iStart), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-stop
func (me *IMediaFilter) Stop() bool {
	ret, _, _ := syscall.Syscall(
		(*_IMediaFilterVtbl)(unsafe.Pointer(*me.Ppv)).Stop, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
