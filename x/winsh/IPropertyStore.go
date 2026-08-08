//go:build windows

package winsh

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IPropertyStore] COM interface.
//
// Implements [OleResource].
//
// [IPropertyStore]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nn-propsys-ipropertystore
type IPropertyStore struct{ win.IUnknown }

type _IPropertyStoreVt struct {
	utl.IUnknownVt
	GetCount uintptr
	GetAt    uintptr
	GetValue uintptr
	SetValue uintptr
	Commit   uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IPropertyStore) IID() *co.IID {
	return &cosh.IID_IPropertyStore
}

// [Commit] method.
//
// [Commit]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nf-propsys-ipropertystore-commit
func (me *IPropertyStore) Commit() error {
	return utl.OleCallWithoutParms(me,
		utl.Vt[_IPropertyStoreVt](me.Ppvt()).Commit)
}

// Returns all [co.PROPERTYKEY] values by calling [IPropertyStore.GetCount] and
// [IPropertyStore.GetAt].
func (me *IPropertyStore) Enum() ([]cosh.PROPERTYKEY, error) {
	count, hr := me.GetCount()
	if hr != nil {
		return nil, hr
	}

	pkeys := make([]cosh.PROPERTYKEY, 0, count)
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
func (me *IPropertyStore) GetAt(index int) (cosh.PROPERTYKEY, error) {
	var key cosh.PROPERTYKEY
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPropertyStoreVt](me.Ppvt()).GetAt,
		me.Ppvt(),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&key)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return key, nil
	} else {
		return cosh.PROPERTYKEY{}, hr
	}
}

// [GetCount] method.
//
// [GetCount]: https://learn.microsoft.com/en-us/windows/win32/api/propsys/nf-propsys-ipropertystore-getcount
func (me *IPropertyStore) GetCount() (int, error) {
	var cProps uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPropertyStoreVt](me.Ppvt()).GetCount,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&cProps)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(cProps), nil
	} else {
		return 0, hr
	}
}
