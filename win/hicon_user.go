//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to an icon.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hicon
type HICON HANDLE

// ⚠️ You must defer HICON.DestroyIcon().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createiconindirect
func CreateIconIndirect(info *ICONINFO) HICON {
	ret, _, err := syscall.SyscallN(proc.CreateIconIndirect.Addr(),
		uintptr(unsafe.Pointer(info)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}

// ⚠️ You must defer HICON.DestroyIcon().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyicon
func (hIcon HICON) CopyIcon() HICON {
	ret, _, err := syscall.SyscallN(proc.CopyIcon.Addr(),
		uintptr(hIcon))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyicon
func (hIcon HICON) DestroyIcon() {
	ret, _, err := syscall.SyscallN(proc.DestroyIcon.Addr(),
		uintptr(hIcon))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HBITMAP.DeleteObject() in HbmMask and HbmColor fields.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfo
func (hIcon HICON) GetIconInfo(iconInfo *ICONINFO) {
	ret, _, err := syscall.SyscallN(proc.GetIconInfo.Addr(),
		uintptr(hIcon), uintptr(unsafe.Pointer(iconInfo)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HBITMAP.DeleteObject() in HbmMask and HbmColor fields.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfoexw
func (hIcon HICON) GetIconInfoEx(iconInfoEx *ICONINFOEX) {
	iconInfoEx.SetCbSize() // safety
	ret, _, err := syscall.SyscallN(proc.GetIconInfoEx.Addr(),
		uintptr(hIcon), uintptr(unsafe.Pointer(iconInfoEx)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
