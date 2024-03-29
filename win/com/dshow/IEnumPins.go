package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IEnumPinsVtbl struct {
	win.IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumpins
type IEnumPins struct {
	win.IUnknown // Base IUnknown.
}

// ⚠️ You must defer IEnumPins.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-clone
func (me *IEnumPins) Clone() IEnumPins {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IEnumPinsVtbl)(unsafe.Pointer(*me.Ppv)).Clone, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IEnumPins{
			win.IUnknown{Ppv: ppQueried},
		}
	} else {
		panic(hr)
	}
}

// Calls Skip() until the end of the enum to retrieve the actual number of pins,
// then calls Reset().
func (me *IEnumPins) Count() int {
	count := int(0)
	for {
		gotOne := me.Skip(1)
		if gotOne {
			count++
		} else {
			me.Reset()
			return count
		}
	}
}

// Calls Next() to retrieve all pins, then calls Reset().
//
// ⚠️ You must defer IPin.Release() on each pin.
func (me *IEnumPins) GetAll() []IPin {
	pins := make([]IPin, 0, 10) // arbitrary
	for {
		pin, gotOne := me.Next()
		if gotOne {
			pins = append(pins, pin)
		} else {
			me.Reset()
			return pins
		}
	}
}

// ⚠️ You must defer IPin.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-next
func (me *IEnumPins) Next() (IPin, bool) {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall6(
		(*_IEnumPinsVtbl)(unsafe.Pointer(*me.Ppv)).Next, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		1, uintptr(unsafe.Pointer(&ppQueried)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IPin{
			win.IUnknown{Ppv: ppQueried},
		}, true
	} else if hr == errco.S_FALSE {
		return IPin{}, false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-reset
func (me *IEnumPins) Reset() {
	syscall.Syscall(
		(*_IEnumPinsVtbl)(unsafe.Pointer(*me.Ppv)).Reset, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-skip
func (me *IEnumPins) Skip(cPins int) bool {
	ret, _, _ := syscall.Syscall(
		(*_IEnumPinsVtbl)(unsafe.Pointer(*me.Ppv)).Skip, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(uint32(cPins)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
