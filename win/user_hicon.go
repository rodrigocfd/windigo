//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to an [icon].
//
// [icon]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hicon
type HICON HANDLE

// [CreateIconFromResourceEx] function.
//
// This function creates [HICON] only. The [HCURSOR] variation is
// [CreateCursorFromResourceEx].
//
// ⚠️ You must defer [HICON.DestroyIcon].
//
// [CreateIconFromResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateIconFromResourceEx(
	resBits []byte,
	fmtVersion uint32,
	cxDesired, cyDesired int,
	flags co.LR,
) (HICON, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_CreateIconFromResourceEx, "CreateIconFromResourceEx"),
		uintptr(unsafe.Pointer(unsafe.SliceData(resBits))),
		uintptr(uint32(len(resBits))),
		1,
		uintptr(fmtVersion),
		uintptr(int32(cxDesired)),
		uintptr(int32(cyDesired)),
		uintptr(flags))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _user_CreateIconFromResourceEx *syscall.Proc

// [CreateIconIndirect] function.
//
// ⚠️ You must defer [HICON.DestroyIcon].
//
// [CreateIconIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconindirect
func CreateIconIndirect(info *ICONINFO) (HICON, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_CreateIconIndirect, "CreateIconIndirect"),
		uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _user_CreateIconIndirect *syscall.Proc

// [CopyIcon] function.
//
// ⚠️ You must defer [HICON.DestroyIcon].
//
// [CopyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyicon
func (hIcon HICON) CopyIcon() (HICON, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_CopyIcon, "CopyIcon"),
		uintptr(hIcon))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _user_CopyIcon *syscall.Proc

// [DestroyIcon] function.
//
// [DestroyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyicon
func (hIcon HICON) DestroyIcon() error {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_DestroyIcon, "DestroyIcon"),
		uintptr(hIcon))
	return utl.ZeroAsGetLastError(ret, err)
}

var _user_DestroyIcon *syscall.Proc

// [GetIconInfo] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject] in HbmMask and HbmColor fields.
//
// [GetIconInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfo
func (hIcon HICON) GetIconInfo() (ICONINFO, error) {
	var ii ICONINFO
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_GetIconInfo, "GetIconInfo"),
		uintptr(hIcon),
		uintptr(unsafe.Pointer(&ii)))
	if ret == 0 {
		return ICONINFO{}, co.ERROR(err)
	}
	return ii, nil
}

var _user_GetIconInfo *syscall.Proc

// [GetIconInfoEx] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject] in HbmMask and HbmColor fields.
//
// [GetIconInfoEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfoexw
func (hIcon HICON) GetIconInfoEx() (ICONINFOEX, error) {
	var ii ICONINFOEX
	ii.SetCbSize()

	ret, _, _ := syscall.SyscallN(
		dll.User.Load(&_user_GetIconInfoExW, "GetIconInfoExW"),
		uintptr(hIcon),
		uintptr(unsafe.Pointer(&ii)))
	if ret == 0 {
		return ICONINFOEX{}, co.ERROR_INVALID_PARAMETER
	}
	return ii, nil
}

var _user_GetIconInfoExW *syscall.Proc
