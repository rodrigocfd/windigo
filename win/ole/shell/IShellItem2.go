//go:build windows

package shell

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IShellItem2] COM interface.
//
// Usually created with [SHCreateItemFromParsingName].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	ish, _ := shell.SHCreateItemFromParsingName[shell.IShellItem2](
//		rel, "C:\\Temp\\foo.txt")
//
// It can also be queried from an [IShellItem] object:
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	item, _ := shell.SHCreateItemFromParsingName[shell.IShellItem](
//		rel, "C:\\Temp\\foo.txt")
//
//	item2, _ := ole.QueryInterface[shell.IShellItem2](
//		&item.IUnknown, rel)
//
// [IShellItem2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem2
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
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
func (me *IShellItem2) GetBool(key *PROPERTYKEY) (bool, error) {
	var bVal int32 // BOOL
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItem2)(unsafe.Pointer(*me.Ppvt())).GetBool,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(&bVal)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return bVal != 0, nil
	} else {
		return false, hr
	}
}

// [GetCLSID] method.
//
// [GetCLSID]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getclsid
func (me *IShellItem2) GetCLSID(key *PROPERTYKEY) (win.GUID, error) {
	var clsid win.GUID
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItem2)(unsafe.Pointer(*me.Ppvt())).GetCLSID,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(&clsid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return clsid, nil
	} else {
		return win.GUID{}, hr
	}
}

// [GetFileTime] method.
//
// [GetFileTime]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getfiletime
func (me *IShellItem2) GetFileTime(key *PROPERTYKEY) (time.Time, error) {
	var ft win.FILETIME
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItem2)(unsafe.Pointer(*me.Ppvt())).GetFileTime,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(&ft)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ft.ToTime(), nil
	} else {
		return time.Time{}, hr
	}
}

// [GetInt32] method.
//
// [GetInt32]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getint32
func (me *IShellItem2) GetInt32(key *PROPERTYKEY) (int32, error) {
	var i int32
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItem2)(unsafe.Pointer(*me.Ppvt())).GetInt32,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(&i)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return i, nil
	} else {
		return 0, hr
	}
}

// [GetString] method.
//
// [GetString]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getstring
func (me *IShellItem2) GetString(key *PROPERTYKEY) (string, error) {
	var psz uintptr
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItem2)(unsafe.Pointer(*me.Ppvt())).GetString,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(&psz)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		defer ole.HTASKMEM(psz).CoTaskMemFree()
		name := wstr.WstrPtrToStr((*uint16)(unsafe.Pointer(psz)))
		return name, nil
	} else {
		return "", hr
	}
}

// [GetUInt32] method.
//
// [GetUInt32]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getuint32
func (me *IShellItem2) GetUInt32(key *PROPERTYKEY) (uint32, error) {
	var ui uint32
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItem2)(unsafe.Pointer(*me.Ppvt())).GetUInt32,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(&ui)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ui, nil
	} else {
		return 0, hr
	}
}

// [GetUInt64] method.
//
// [GetUInt64]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-getuint64
func (me *IShellItem2) GetUInt64(key *PROPERTYKEY) (uint64, error) {
	var ull uint64
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItem2)(unsafe.Pointer(*me.Ppvt())).GetUInt64,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(&ull)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return ull, nil
	} else {
		return 0, hr
	}
}

// [Update] method.
//
// [Update]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem2-update
func (me *IShellItem2) Update(bc *ole.IBindCtx) error {
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItem2)(unsafe.Pointer(*me.Ppvt())).GetUInt64,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(bc.Ppvt())))
	return wutil.ErrorAsHResult(ret)
}
