//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [IShellItem2] COM interface.
//
// Implements [OleObj] and [OleResource].
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
//	var item *win.IShellItem2
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
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
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	var item2 *win.IShellItem2
//	_ = item.QueryInterface(rel, &item2)
//
// [IShellItem2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem2
type IShellItem2 struct{ IShellItem }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItem2) IID() co.IID {
	return co.IID_IShellItem2
}

// [GetBool] method.
//
// [GetBool]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getbool
func (me *IShellItem2) GetBool(pkey co.PKEY) (bool, error) {
	guidPkey := PropertykeyFrom(pkey)
	var bVal int32 // BOOL

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetBool,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&bVal)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return bVal != 0, nil
	} else {
		return false, hr
	}
}

// [GetCLSID] method.
//
// [GetCLSID]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getclsid
func (me *IShellItem2) GetCLSID(pkey co.PKEY) (co.CLSID, error) {
	guidPkey := PropertykeyFrom(pkey)
	var guidClsid GUID

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetCLSID,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&guidClsid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.CLSID(guidClsid.String()), nil
	} else {
		return co.CLSID(""), hr
	}
}

// [GetFileTime] method.
//
// [GetFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getfiletime
func (me *IShellItem2) GetFileTime(pkey co.PKEY) (time.Time, error) {
	guidPkey := PropertykeyFrom(pkey)
	var ft FILETIME

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetFileTime,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&ft)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ft.ToTime(), nil
	} else {
		return time.Time{}, hr
	}
}

// [GetInt32] method.
//
// [GetInt32]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getint32
func (me *IShellItem2) GetInt32(pkey co.PKEY) (int32, error) {
	guidPkey := PropertykeyFrom(pkey)
	var i int32

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetInt32,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&i)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return i, nil
	} else {
		return 0, hr
	}
}

// [GetPropertyStore] method.
//
// [GetPropertyStore]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getpropertystore
func (me *IShellItem2) GetPropertyStore(releaser *OleReleaser, flags co.GPS) (*IPropertyStore, error) {
	var ppvtQueried **_IUnknownVt
	guid := GuidFrom(co.IID_IPropertyStore)

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetPropertyStore,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(flags),
		uintptr(unsafe.Pointer(&guid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IPropertyStore{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
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
//	var item *win.IShellItem2
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp\\foo.txt", &item)
//
//	ty, _ := item.GetString(co.PKEY_ItemTypeText)
//	println(ty)
//
// [GetString]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getstring
func (me *IShellItem2) GetString(pkey co.PKEY) (string, error) {
	guidPkey := PropertykeyFrom(pkey)
	var psz uintptr

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetString,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&psz)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer HTASKMEM(psz).CoTaskMemFree()
		name := wstr.DecodePtr((*uint16)(unsafe.Pointer(psz)))
		return name, nil
	} else {
		return "", hr
	}
}

// [GetUInt32] method.
//
// [GetUInt32]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getuint32
func (me *IShellItem2) GetUInt32(pkey co.PKEY) (uint32, error) {
	guidPkey := PropertykeyFrom(pkey)
	var ui uint32

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetUInt32,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&ui)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ui, nil
	} else {
		return 0, hr
	}
}

// [GetUInt64] method.
//
// [GetUInt64]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getuint64
func (me *IShellItem2) GetUInt64(pkey co.PKEY) (uint64, error) {
	guidPkey := PropertykeyFrom(pkey)
	var ull uint64

	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetUInt64,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&guidPkey)),
		uintptr(unsafe.Pointer(&ull)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ull, nil
	} else {
		return 0, hr
	}
}

// [Update] method.
//
// [Update]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-update
func (me *IShellItem2) Update(bc *IBindCtx) error {
	ret, _, _ := syscall.SyscallN(
		(*_IShellItem2Vt)(unsafe.Pointer(*me.Ppvt())).GetUInt64,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(bc.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

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
