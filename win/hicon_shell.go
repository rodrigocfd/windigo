//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

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
