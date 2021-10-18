package dshow

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/com/oidl"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IMediaFilterVtbl struct {
	oidl.IPersistVtbl
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
	oidl.IPersist // Base IPersist > IUnknown.
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-getstate
func (me *IMediaFilter) GetState(msTimeout int) (dshowco.FILTER_STATE, error) {
	var state dshowco.FILTER_STATE
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
func (me *IMediaFilter) Run(start time.Duration) bool {
	iStart := util.DurationToNano100(start)
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
