package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumfilters
type IEnumFilters struct{ win.IUnknown }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer IEnumFilters.Release().
func NewIEnumFilters(ptr win.IUnknownPtr) IEnumFilters {
	return IEnumFilters{
		IUnknown: win.NewIUnknown(ptr),
	}
}

// ⚠️ You must defer IEnumFilters.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-clone
func (me *IEnumFilters) Clone() IEnumFilters {
	var ppQueried win.IUnknownPtr
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumFilters)(unsafe.Pointer(*me.Ptr())).Clone, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumFilters(ppQueried)
	} else {
		panic(hr)
	}
}

// Calls Skip() until the end of the enum to retrieve the actual number of
// filters, then calls Reset().
func (me *IEnumFilters) Count() int {
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

// Calls Next() to retrieve all filters, then calls Reset().
//
// ⚠️ You must defer IBaseFilter.Release() on each filter.
func (me *IEnumFilters) GetAll() []IBaseFilter {
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

// ⚠️ You must defer IBaseFilter.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-next
func (me *IEnumFilters) Next() (IBaseFilter, bool) {
	var ppQueried win.IUnknownPtr
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IEnumFilters)(unsafe.Pointer(*me.Ptr())).Next, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		1, uintptr(unsafe.Pointer(&ppQueried)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIBaseFilter(ppQueried), true
	} else if hr == errco.S_FALSE {
		return IBaseFilter{}, false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-reset
func (me *IEnumFilters) Reset() {
	syscall.Syscall(
		(*dshowvt.IEnumFilters)(unsafe.Pointer(*me.Ptr())).Reset, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-skip
func (me *IEnumFilters) Skip(numFilters int) bool {
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
