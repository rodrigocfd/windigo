package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumpins
type IEnumPins struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IEnumPins.Release().
func NewIEnumPins(base win.IUnknown) IEnumPins {
	return IEnumPins{IUnknown: base}
}

// ⚠️ You must defer IEnumPins.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-clone
func (me *IEnumPins) Clone() IEnumPins {
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumPins)(unsafe.Pointer(*me.Ptr())).Clone, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumPins(ppQueried)
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
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IEnumPins)(unsafe.Pointer(*me.Ptr())).Next, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		1, uintptr(unsafe.Pointer(&ppQueried)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(ppQueried), true
	} else if hr == errco.S_FALSE {
		return IPin{}, false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-reset
func (me *IEnumPins) Reset() {
	syscall.Syscall(
		(*dshowvt.IEnumPins)(unsafe.Pointer(*me.Ptr())).Reset, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-skip
func (me *IEnumPins) Skip(numPins int) bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumPins)(unsafe.Pointer(*me.Ptr())).Skip, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(uint32(numPins)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
