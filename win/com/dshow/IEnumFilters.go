package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumfilters
type IEnumFilters interface {
	com.IUnknown

	// ⚠️ You must defer IEnumFilters.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-clone
	Clone() IEnumFilters

	// Calls Skip() until the end of the enum to retrieve the actual number of
	// filters, then calls Reset().
	Count() int

	// Calls Next() to retrieve all filters, then calls Reset().
	//
	// ⚠️ You must defer IBaseFilter.Release() on each returned object.
	GetAll() []IBaseFilter

	// ⚠️ You must defer IBaseFilter.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-next
	Next() (IBaseFilter, bool)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-reset
	Reset()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-skip
	Skip(numFilters int) bool
}

type _IEnumFilters struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IEnumFilters.Release().
func NewIEnumFilters(base com.IUnknown) IEnumFilters {
	return &_IEnumFilters{IUnknown: base}
}

func (me *_IEnumFilters) Clone() IEnumFilters {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumFilters)(unsafe.Pointer(*me.Ptr())).Clone, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumFilters(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_IEnumFilters) Count() int {
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

func (me *_IEnumFilters) GetAll() []IBaseFilter {
	filters := make([]IBaseFilter, 0, 10) // arbitrary
	for {
		filter, gotOne := me.Next()
		if gotOne {
			filters = append(filters, filter)
		} else {
			me.Reset()
			return filters
		}
	}
}

func (me *_IEnumFilters) Next() (IBaseFilter, bool) {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IEnumFilters)(unsafe.Pointer(*me.Ptr())).Next, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		1, uintptr(unsafe.Pointer(&ppQueried)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIBaseFilter(com.NewIUnknown(ppQueried)), true
	} else if hr == errco.S_FALSE {
		return nil, false
	} else {
		panic(hr)
	}
}

func (me *_IEnumFilters) Reset() {
	syscall.Syscall(
		(*dshowvt.IEnumFilters)(unsafe.Pointer(*me.Ptr())).Reset, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)
}

func (me *_IEnumFilters) Skip(numFilters int) bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumFilters)(unsafe.Pointer(*me.Ptr())).Skip, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(uint32(numFilters)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
