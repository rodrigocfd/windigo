//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [IPropertyStore] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IPropertyStore]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nn-propsys-ipropertystore
type IPropertyStore struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IPropertyStore) IID() co.IID {
	return co.IID_IPropertyStore
}

// [Commit] method.
//
// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nf-propsys-ipropertystore-commit
func (me *IPropertyStore) Commit() error {
	ret, _, _ := syscall.SyscallN(
		(*_IPropertyStoreVt)(unsafe.Pointer(*me.Ppvt())).Commit,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// Returns all [co.PKEY] values by calling [IPropertyStore.GetCount] and
// [IPropertyStore.GetAt].
func (me *IPropertyStore) Enum() ([]co.PKEY, error) {
	count, hr := me.GetCount()
	if hr != nil {
		return nil, hr
	}

	pkeys := make([]co.PKEY, 0, count)
	for i := 0; i < count; i++ {
		pkey, hr := me.GetAt(i)
		if hr != nil {
			return nil, hr
		}
		pkeys = append(pkeys, pkey)
	}
	return pkeys, nil
}

// [GetAt] method.
//
// [GetAt]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nf-propsys-ipropertystore-getat
func (me *IPropertyStore) GetAt(index int) (co.PKEY, error) {
	var guidPkey GUID
	ret, _, _ := syscall.SyscallN(
		(*_IPropertyStoreVt)(unsafe.Pointer(*me.Ppvt())).GetAt,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&guidPkey)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.PKEY(guidPkey.String()), nil
	} else {
		return co.PKEY(""), hr
	}
}

// [GetCount] method.
//
// [GetCount]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nf-propsys-ipropertystore-getcount
func (me *IPropertyStore) GetCount() (int, error) {
	var cProps uint32
	ret, _, _ := syscall.SyscallN(
		(*_IPropertyStoreVt)(unsafe.Pointer(*me.Ppvt())).GetCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&cProps)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(cProps), nil
	} else {
		return 0, hr
	}
}

type _IPropertyStoreVt struct {
	_IUnknownVt
	GetCount uintptr
	GetAt    uintptr
	GetValue uintptr
	SetValue uintptr
	Commit   uintptr
}
