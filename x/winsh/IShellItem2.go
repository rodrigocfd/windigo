//go:build windows

package winsh

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/wstr"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IShellItem2] COM interface.
//
// Usually created with [SHCreateItemFromParsingName].
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *winsh.IShellItem2
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
// It can also be queried from an [IShellItem] object:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	var item2 *winsh.IShellItem2
//	_ = item.QueryInterface(rel, &item2)
//
// [IShellItem2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem2
type IShellItem2 struct{ IShellItem }

type _IShellItem2Vt struct {
	_IShellItemVt
	GetPropertyStore                 uintptr
	GetPropertyStoreWithCreateObject uintptr
	GetPropertyStoreForKeys          uintptr
	GetPropertyDescriptionList       uintptr
	Update                           uintptr
	GetProperty                      uintptr
	GetCLSID                         uintptr
	GetFileTime                      uintptr
	GetInt32                         uintptr
	GetString                        uintptr
	GetUInt32                        uintptr
	GetUInt64                        uintptr
	GetBool                          uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItem2) IID() *co.IID {
	return &cosh.IID_IShellItem2
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IShellItem2) AddRef(releaser *win.OleReleaser) *IShellItem2 {
	return utl.OleNewFromAddRef[*IShellItem2](me, releaser)
}

// [GetBool] method.
//
// [GetBool]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getbool
func (me *IShellItem2) GetBool(pKey *cosh.PROPERTYKEY) (bool, error) {
	var bVal win.BOOL
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetBool,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pKey)),
		uintptr(unsafe.Pointer(&bVal)))
	return utl.HresultToBoolError(int32(bVal), ret)
}

// [GetCLSID] method.
//
// [GetCLSID]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getclsid
func (me *IShellItem2) GetCLSID(pKey *cosh.PROPERTYKEY) (co.CLSID, error) {
	var clsid co.CLSID
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetCLSID,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pKey)),
		uintptr(unsafe.Pointer(&clsid)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return co.CLSID{}, hr
	}
	return clsid, nil
}

// [GetFileTime] method.
//
// [GetFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getfiletime
func (me *IShellItem2) GetFileTime(pKey *cosh.PROPERTYKEY) (time.Time, error) {
	var ft win.FILETIME
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetFileTime,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pKey)),
		uintptr(unsafe.Pointer(&ft)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return time.Time{}, hr
	}
	return ft.ToTime(), nil
}

// [GetInt32] method.
//
// [GetInt32]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getint32
func (me *IShellItem2) GetInt32(pKey *cosh.PROPERTYKEY) (int32, error) {
	var i int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetInt32,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pKey)),
		uintptr(unsafe.Pointer(&i)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return i, nil
}

// [GetPropertyStore] method.
//
// [GetPropertyStore]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getpropertystore
func (me *IShellItem2) GetPropertyStore(releaser *win.OleReleaser, flags cosh.GPS) (*IPropertyStore, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetPropertyStore,
		me.Ppvt(),
		uintptr(flags),
		uintptr(unsafe.Pointer(&cosh.IID_IPropertyStore)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IPropertyStore](ret, ppvtQueried, releaser)
}

// [GetString] method.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *winsh.IShellItem2
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	ty, _ := item.GetString(cosh.PKEY_ItemTypeText)
//	println(ty)
//
// [GetString]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getstring
func (me *IShellItem2) GetString(pKey *cosh.PROPERTYKEY) (string, error) {
	var psz uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetString,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pKey)),
		uintptr(unsafe.Pointer(&psz)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	defer win.HTASKMEM(psz).CoTaskMemFree()
	name := wstr.DecodePtr((*uint16)(unsafe.Pointer(psz)))
	return name, nil
}

// [GetUInt32] method.
//
// [GetUInt32]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getuint32
func (me *IShellItem2) GetUInt32(pKey *cosh.PROPERTYKEY) (uint32, error) {
	var ui uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetUInt32,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pKey)),
		uintptr(unsafe.Pointer(&ui)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return ui, nil
}

// [GetUInt64] method.
//
// [GetUInt64]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getuint64
func (me *IShellItem2) GetUInt64(pKey *cosh.PROPERTYKEY) (uint64, error) {
	var ull uint64
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetUInt64,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pKey)),
		uintptr(unsafe.Pointer(&ull)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return ull, nil
}

// [Update] method.
//
// [Update]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-update
func (me *IShellItem2) Update(bc *win.IBindCtx) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItem2Vt](me.Ppvt()).GetUInt64,
		me.Ppvt(),
		bc.Ppvt())
	return utl.HresultToError(ret)
}
