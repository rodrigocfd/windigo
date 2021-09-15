package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to an icon.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hicon
type HICON HANDLE

// Extracts all icons: big and small.
//
// ⚠️ You must defer HICON.DestroyIcon() on each icon returned in both slices.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-extracticonexw
func ExtractIconEx(fileName string) (largeIcons, smallIcons []HICON) {
	lpszFile16 := Str.ToNativeSlice(fileName)
	retrieveIdx := -1
	ret, _, err := syscall.Syscall6(proc.ExtractIconEx.Addr(), 5,
		uintptr(unsafe.Pointer(&lpszFile16[0])), uintptr(retrieveIdx), 0, 0, 0, 0)
	if ret == _UINT_MAX {
		panic(errco.ERROR(err))
	}

	numIcons := int(ret)
	largeIcons = make([]HICON, numIcons)
	smallIcons = make([]HICON, numIcons)

	ret, _, err = syscall.Syscall6(proc.ExtractIconEx.Addr(), 5,
		uintptr(unsafe.Pointer(&lpszFile16[0])), 0,
		uintptr(unsafe.Pointer(&largeIcons[0])),
		uintptr(unsafe.Pointer(&smallIcons[0])),
		uintptr(numIcons), 0)
	runtime.KeepAlive(lpszFile16)
	if ret == _UINT_MAX {
		panic(errco.ERROR(err))
	}

	return
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-copyicon
func (hIcon HICON) CopyIcon() HICON {
	ret, _, err := syscall.Syscall(proc.CopyIcon.Addr(), 1,
		uintptr(hIcon), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HICON(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroyicon
func (hIcon HICON) DestroyIcon() {
	ret, _, err := syscall.Syscall(proc.DestroyIcon.Addr(), 1,
		uintptr(hIcon), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HBITMAP.DeleteObject() in HbmMask and HbmColor fields.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfo
func (hIcon HICON) GetIconInfo(iconInfo *ICONINFO) {
	ret, _, err := syscall.Syscall(proc.GetIconInfo.Addr(), 2,
		uintptr(hIcon), uintptr(unsafe.Pointer(iconInfo)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer HBITMAP.DeleteObject() in HbmMask and HbmColor fields.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfoexw
func (hIcon HICON) GetIconInfoEx(iconInfoEx *ICONINFOEX) {
	iconInfoEx.SetCbSize() // safety
	ret, _, err := syscall.Syscall(proc.GetIconInfoEx.Addr(), 2,
		uintptr(hIcon), uintptr(unsafe.Pointer(iconInfoEx)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
