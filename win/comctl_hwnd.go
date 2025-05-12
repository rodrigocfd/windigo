//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [DefSubclassProc] function.
//
// [DefSubclassProc]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-defsubclassproc
func (hWnd HWND) DefSubclassProc(msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {
	ret, _, _ := syscall.SyscallN(_DefSubclassProc.Addr(),
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return ret
}

var _DefSubclassProc = dll.Comctl32.NewProc("DefSubclassProc")

// [RemoveWindowSubclass] function.
//
// [RemoveWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(subclassProc uintptr, idSubclass uint32) error {
	ret, _, _ := syscall.SyscallN(_RemoveWindowSubclass.Addr(),
		uintptr(hWnd), subclassProc, uintptr(idSubclass))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _RemoveWindowSubclass = dll.Comctl32.NewProc("RemoveWindowSubclass")

// [SetWindowSubclass] function.
//
// [SetWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func (hWnd HWND) SetWindowSubclass(subclassProc uintptr, idSubclass uint32, refData unsafe.Pointer) error {
	ret, _, _ := syscall.SyscallN(_SetWindowSubclass.Addr(),
		uintptr(hWnd), subclassProc, uintptr(idSubclass), uintptr(refData))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _SetWindowSubclass = dll.Comctl32.NewProc("SetWindowSubclass")
