package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IEnumFiltersVtbl struct {
	win.IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumfilters
type IEnumFilters struct {
	win.IUnknown // Base IUnknown.
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-clone
func (me *IEnumFilters) Clone() IEnumFilters {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IEnumFiltersVtbl)(unsafe.Pointer(*me.Ppv)).Clone, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IEnumFilters{
			win.IUnknown{Ppv: ppQueried},
		}
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
// ⚠️ You must defer Release() on each filter.
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

// ⚠️ You must defer Release() if true.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-next
func (me *IEnumFilters) Next() (IBaseFilter, bool) {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall6(
		(*_IEnumFiltersVtbl)(unsafe.Pointer(*me.Ppv)).Next, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		1, uintptr(unsafe.Pointer(&ppQueried)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return IBaseFilter{
			IMediaFilter{
				IPersist{
					win.IUnknown{Ppv: ppQueried},
				},
			},
		}, true
	} else if hr == errco.S_FALSE {
		return IBaseFilter{}, false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-reset
func (me *IEnumFilters) Reset() {
	syscall.Syscall(
		(*_IEnumFiltersVtbl)(unsafe.Pointer(*me.Ppv)).Reset, 1,
		uintptr(unsafe.Pointer(me.Ppv)),
		0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienumfilters-skip
func (me *IEnumFilters) Skip(cFilters int) bool {
	ret, _, _ := syscall.Syscall(
		(*_IEnumFiltersVtbl)(unsafe.Pointer(*me.Ppv)).Skip, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(uint32(cFilters)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
