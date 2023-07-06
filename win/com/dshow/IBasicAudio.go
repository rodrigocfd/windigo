//go:build windows

package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/autom"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IBasicAudio] COM interface.
//
// [IBasicAudio]: https://learn.microsoft.com/en-us/windows/win32/api/control/nn-control-ibasicaudio
type IBasicAudio interface {
	autom.IDispatch

	// [GetBalance] COM method.
	//
	// [GetBalance]: https://learn.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-get_balance
	GetBalance() int

	// [GetVolume] COM method.
	//
	// [GetVolume]: https://learn.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-get_volume
	GetVolume() int

	// [PutBalance] COM method.
	//
	// [PutBalance]: https://learn.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-put_balance
	PutBalance(balance int)

	// [PutVolume] COM method.
	//
	// [PutVolume]: https://learn.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-put_volume
	PutVolume(volume int)
}

type _IBasicAudio struct{ autom.IDispatch }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IBasicAudio.Release().
//
// # Example
//
//	var gb dshow.IGraphBuilder // initialized somewhere
//
//	ba := dshow.NewIBasicAudio(
//		gb.QueryInterface(dshowco.IID_IBasicAudio),
//	)
//	defer ba.Release()
func NewIBasicAudio(base com.IUnknown) IBasicAudio {
	return &_IBasicAudio{IDispatch: autom.NewIDispatch(base)}
}

func (me *_IBasicAudio) GetBalance() int {
	var balance int32
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IBasicAudio)(unsafe.Pointer(*me.Ptr())).GetBalance,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&balance)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(balance)
	} else {
		panic(hr)
	}
}

func (me *_IBasicAudio) GetVolume() int {
	var volume int32
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IBasicAudio)(unsafe.Pointer(*me.Ptr())).GetVolume,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&volume)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(volume)
	} else {
		panic(hr)
	}
}

func (me *_IBasicAudio) PutBalance(balance int) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IBasicAudio)(unsafe.Pointer(*me.Ptr())).PutBalance,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(int32(balance)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IBasicAudio) PutVolume(volume int) {
	ret, _, _ := syscall.SyscallN(
		(*dshowvt.IBasicAudio)(unsafe.Pointer(*me.Ptr())).PutVolume,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(int32(volume)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
