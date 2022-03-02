package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumpins
type IEnumPins interface {
	com.IUnknown

	// ⚠️ You must defer IEnumPins.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-clone
	Clone() IEnumPins

	// Calls Skip() until the end of the enum to retrieve the actual number of
	// pins, then calls Reset().
	Count() int

	// Calls Next() to retrieve all pins, then calls Reset().
	//
	// ⚠️ You must defer IPin.Release() on each returned object.
	GetAll() []IPin

	// ⚠️ You must defer IPin.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-next
	Next() (IPin, bool)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-reset
	Reset()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumpins-skip
	Skip(numPins int) bool
}

type _IEnumPins struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IEnumPins.Release().
func NewIEnumPins(base com.IUnknown) IEnumPins {
	return &_IEnumPins{IUnknown: base}
}

func (me *_IEnumPins) Clone() IEnumPins {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumPins)(unsafe.Pointer(*me.Ptr())).Clone, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumPins(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_IEnumPins) Count() int {
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

func (me *_IEnumPins) GetAll() []IPin {
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

func (me *_IEnumPins) Next() (IPin, bool) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IEnumPins)(unsafe.Pointer(*me.Ptr())).Next, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		1, uintptr(unsafe.Pointer(&ppQueried)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIPin(com.NewIUnknown(ppQueried)), true
	} else if hr == errco.S_FALSE {
		return nil, false
	} else {
		panic(hr)
	}
}

func (me *_IEnumPins) Reset() {
	syscall.Syscall(
		(*dshowvt.IEnumPins)(unsafe.Pointer(*me.Ptr())).Reset, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)
}

func (me *_IEnumPins) Skip(numPins int) bool {
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
