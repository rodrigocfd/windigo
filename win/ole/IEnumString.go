//go:build windows

package ole

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IEnumString] COM interface.
//
// Implements [ComObj] and [ComResource].
//
// [IEnumString]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ienumstring
type IEnumString struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumString) IID() co.IID {
	return co.IID_IEnumString
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-clone
func (me *IEnumString) Clone(releaser *Releaser) (*IEnumString, error) {
	var ppvtQueried **IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IEnumStringVt)(unsafe.Pointer(*me.Ppvt())).Clone,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		var pObj *IEnumString
		utl.ComCreateObj(&pObj, unsafe.Pointer(ppvtQueried))
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// Returns all objects by calling [IEnumString.Next].
func (me *IEnumString) Enum() ([]string, error) {
	strs := make([]string, 0)
	var s string
	var hr error

	for {
		s, hr = me.Next()
		if hr != nil { // actual error
			return nil, hr
		} else if s == "" { // no more items to fetch
			return strs, nil
		} else { // item fetched
			strs = append(strs, s)
		}
	}
}

// [Next] method.
//
// [Next]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-next
func (me *IEnumString) Next() (string, error) {
	var pv uintptr
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		(*_IEnumStringVt)(unsafe.Pointer(*me.Ppvt())).Next,
		uintptr(unsafe.Pointer(me.Ppvt())),
		1,
		uintptr(unsafe.Pointer(&pv)),
		uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer HTASKMEM(pv).CoTaskMemFree()
		name := wstr.WstrPtrToStr((*uint16)(unsafe.Pointer(pv)))
		return name, nil
	} else if hr == co.HRESULT_S_FALSE {
		return "", nil
	} else {
		return "", hr
	}
}

// [Reset] method.
//
// [Reset]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-reset
func (me *IEnumString) Reset() error {
	ret, _, _ := syscall.SyscallN(
		(*_IEnumStringVt)(unsafe.Pointer(*me.Ppvt())).Reset,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [Skip] method.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumstring-skip
func (me *IEnumString) Skip(count uint) error {
	ret, _, _ := syscall.SyscallN(
		(*_IEnumStringVt)(unsafe.Pointer(*me.Ppvt())).Skip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(count)))
	return utl.ErrorAsHResult(ret)
}

type _IEnumStringVt struct {
	IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}
