//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
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
	szDesired SIZE,
	flags co.LR,
) (HICON, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_CreateIconFromResourceEx, "CreateIconFromResourceEx"),
		uintptr(unsafe.Pointer(&resBits[0])),
		uintptr(uint32(len(resBits))),
		1,
		uintptr(fmtVersion),
		uintptr(szDesired.Cx),
		uintptr(szDesired.Cy),
		uintptr(flags))
	if ret == 0 {
		return HICON(0), co.ERROR(err)
	}
	return HICON(ret), nil
}

var _user_CreateIconFromResourceEx *syscall.Proc

// Calls [SHGetFileInfo] to retrieve the icon correspondent to the given file
// extension in Windows Explorer.
//
// This function is fully implemented in Shell package, but a minimal subset is
// available here to avoid a circular dependency, since it's required by
// HIMAGELIST and ui.
//
// Panics if size is different from 16 or 32. Only 16x16 and 32x32 icons can be
// retrieved.
//
// ⚠️ You must defer [HICON.DestroyIcon].
//
// [SHGetFileInfo]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shgetfileinfow
func LoadIconOfFileExt(fileExtension string, size int) (HICON, error) {
	if size != 16 && size != 32 {
		panic("Size of icons from file extension must be 16 or 32.")
	}

	var pathBuf [20]uint16
	wstr.EncodeToBuf(pathBuf[:], "*."+fileExtension)

	type SHFILEINFO struct {
		HIcon         HICON
		IIcon         int32
		DwAttributes  uint32
		szDisplayName [utl.MAX_PATH]uint16
		szTypeName    [80]uint16
	}
	var sfi SHFILEINFO

	const (
		SHGFI_ICON              = 0x0000_0100
		SHGFI_LARGEICON         = 0x0000_0000
		SHGFI_SMALLICON         = 0x0000_0001
		SHGFI_USEFILEATTRIBUTES = 0x0000_0010
	)
	shgfi := SHGFI_USEFILEATTRIBUTES | SHGFI_ICON
	if size == 16 {
		shgfi |= SHGFI_SMALLICON
	} else {
		shgfi |= SHGFI_LARGEICON
	}

	ret, _, _ := syscall.SyscallN(
		dll.Shell.Load(&utl.Shell_SHGetFileInfoW, "SHGetFileInfoW"), // note: syscall.Proc from utl
		uintptr(unsafe.Pointer(&pathBuf[0])),
		uintptr(co.FILE_ATTRIBUTE_NORMAL),
		uintptr(unsafe.Pointer(&sfi)),
		unsafe.Sizeof(sfi),
		uintptr(shgfi))
	if ret == 0 {
		return HICON(0), co.ERROR_UNIDENTIFIED_ERROR
	}
	return sfi.HIcon, nil
}

// [CreateIconIndirect] function.
//
// ⚠️ You must defer [HICON.DestroyIcon].
//
// [CreateIconIndirect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconindirect
func CreateIconIndirect(pInfo *ICONINFO) (HICON, error) {
	ret, _, err := syscall.SyscallN(
		dll.User.Load(&_user_CreateIconIndirect, "CreateIconIndirect"),
		uintptr(unsafe.Pointer(pInfo)))
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
		return ICONINFOEX{}, co.ERROR_UNIDENTIFIED_ERROR
	}
	return ii, nil
}

var _user_GetIconInfoExW *syscall.Proc
