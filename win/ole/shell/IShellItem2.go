//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IShellItem2] COM interface.
//
// / Usually created with [SHCreateItemFromParsingName].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	ish, _ := shell.SHCreateItemFromParsingName[shell.IShellItem2](
//		rel, "C:\\Temp\\foo.txt")
//
// [IShellItem2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem2
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
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
