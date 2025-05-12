//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to an [icon].
//
// [icon]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hicon
type HICON HANDLE

// [CreateIconIndirect] function.
//
// ⚠️ You must defer [HICON.DestroyIcon].
//
// [CreateIconIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconindirect
func CreateIconIndirect(info *ICONINFO) (HICON, error) {
	ret, _, err := syscall.SyscallN(_CreateIconIndirect.Addr(),
		uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _CreateIconIndirect = dll.User32.NewProc("CreateIconIndirect")

// [CopyIcon] function.
//
// ⚠️ You must defer [HICON.DestroyIcon].
//
// [CopyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyicon
func (hIcon HICON) CopyIcon() (HICON, error) {
	ret, _, err := syscall.SyscallN(_CopyIcon.Addr(),
		uintptr(hIcon))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _CopyIcon = dll.User32.NewProc("CopyIcon")

// [DestroyIcon] function.
//
// [DestroyIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyicon
func (hIcon HICON) DestroyIcon() error {
	ret, _, err := syscall.SyscallN(_DestroyIcon.Addr(),
		uintptr(hIcon))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _DestroyIcon = dll.User32.NewProc("DestroyIcon")

// [GetIconInfo] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject] in HbmMask and HbmColor fields.
//
// [GetIconInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfo
func (hIcon HICON) GetIconInfo() (ICONINFO, error) {
	var ii ICONINFO
	ret, _, err := syscall.SyscallN(_GetIconInfo.Addr(),
		uintptr(hIcon), uintptr(unsafe.Pointer(&ii)))
	if ret == 0 {
		return ICONINFO{}, co.ERROR(err)
	}
	return ii, nil
}

var _GetIconInfo = dll.User32.NewProc("GetIconInfo")

// [GetIconInfoEx] function.
//
// ⚠️ You must defer [HBITMAP.DeleteObject] in HbmMask and HbmColor fields.
//
// [GetIconInfoEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfoexw
func (hIcon HICON) GetIconInfoEx() (ICONINFOEX, error) {
	var ii ICONINFOEX
	ii.SetCbSize()

	ret, _, _ := syscall.SyscallN(_GetIconInfoExW.Addr(),
		uintptr(hIcon), uintptr(unsafe.Pointer(&ii)))
	if ret == 0 {
		return ICONINFOEX{}, co.ERROR_INVALID_PARAMETER
	}
	return ii, nil
}

var _GetIconInfoExW = dll.User32.NewProc("GetIconInfoExW")
