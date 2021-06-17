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

// Extracts all icons: big and small.
//
// ⚠️ You must defer DestroyIcon() on each icon returned in both slices.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-extracticonexw
func ExtractIconEx(lpszFile string) ([]HICON, []HICON) {
	retrieveIdx := -1
	ret, _, err := syscall.Syscall6(proc.ExtractIconEx.Addr(), 5,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpszFile))),
		uintptr(retrieveIdx), 0, 0, 0, 0)
	if ret == _UINT_MAX {
		panic(errco.ERROR(err))
	}

	numIcons := int(ret)
	largeIcons := make([]HICON, numIcons)
	smallIcons := make([]HICON, numIcons)

	ret, _, err = syscall.Syscall6(proc.ExtractIconEx.Addr(), 5,
		uintptr(unsafe.Pointer(Str.ToUint16Ptr(lpszFile))), 0,
		uintptr(unsafe.Pointer(&largeIcons[0])),
		uintptr(unsafe.Pointer(&smallIcons[0])), uintptr(numIcons), 0)
	if ret == _UINT_MAX {
		panic(errco.ERROR(err))
	}

	return largeIcons, smallIcons
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

// ⚠️ You must defer DeleteObject() in HbmMask and HbmColor fields.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfo
func (hIcon HICON) GetIconInfo(piconinfo *ICONINFO) {
	ret, _, err := syscall.Syscall(proc.GetIconInfo.Addr(), 2,
		uintptr(hIcon), uintptr(unsafe.Pointer(piconinfo)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ You must defer DeleteObject() in HbmMask and HbmColor fields.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-geticoninfoexw
func (hIcon HICON) GetIconInfoEx(piconinfo *ICONINFOEX) {
	piconinfo.CbSize = uint32(unsafe.Sizeof(*piconinfo)) // safety
	ret, _, err := syscall.Syscall(proc.GetIconInfoEx.Addr(), 2,
		uintptr(hIcon), uintptr(unsafe.Pointer(piconinfo)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
