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

// [IMediaFilter] COM interface.
//
// [IMediaFilter]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
type IMediaFilter interface {
	com.IPersist

	// [GetState] COM method.
	//
	// [GetState]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-getstate
	GetState(msTimeout win.NumInf) (dshowco.FILTER_STATE, error)

	// [Pause] COM method.
	//
	// [Pause]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-pause
	Pause() bool

	// [Run] COM method.
	//
	// [Run]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-run
	Run(start time.Duration) bool

	// [Stop] COM method.
	//
	// [Stop]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-stop
	Stop() bool
}

type _IMediaFilter struct{ com.IPersist }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMediaFilter.Release().
func NewIMediaFilter(base com.IUnknown) IMediaFilter {
	return &_IMediaFilter{IPersist: com.NewIPersist(base)}
}

func (me *_IMediaFilter) GetState(msTimeout win.NumInf) (dshowco.FILTER_STATE, error) {
	var state dshowco.FILTER_STATE
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).GetState,
		uintptr(unsafe.Pointer(me.Ptr())),
		msTimeout.Raw(), uintptr(unsafe.Pointer(&state)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return state, nil
	} else if hr == errco.VFW_S_STATE_INTERMEDIATE {
		return state, hr
	} else {
		panic(hr)
	}
}

func (me *_IMediaFilter) Pause() bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Pause,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IMediaFilter) Run(start time.Duration) bool {
	iStart := util.DurationToNano100(start)
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Run,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(iStart))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IMediaFilter) Stop() bool {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Stop,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
