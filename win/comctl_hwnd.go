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
	ret, _, _ := syscall.SyscallN(dll.Comctl(dll.PROC_DefSubclassProc),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

// [ImageList_DragEnter] function.
//
// [ImageList_DragEnter]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragenter
func (hWnd HWND) ImageListDragEnter(x, y int) error {
	ret, _, _ := syscall.SyscallN(dll.Comctl(dll.PROC_ImageList_DragEnter),
		uintptr(hWnd),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [ImageList_DragLeave] function.
//
// [ImageList_DragLeave]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragleave
func (hWnd HWND) ImageListDragLeave() error {
	ret, _, _ := syscall.SyscallN(dll.Comctl(dll.PROC_ImageList_DragLeave),
		uintptr(hWnd))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [RemoveWindowSubclass] function.
//
// [RemoveWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(subclassProc uintptr, idSubclass uint32) error {
	ret, _, _ := syscall.SyscallN(dll.Comctl(dll.PROC_RemoveWindowSubclass),
		uintptr(hWnd),
		subclassProc,
		uintptr(idSubclass))
	return utl.ZeroAsSysInvalidParm(ret)
}

// [SetWindowSubclass] function.
//
// [SetWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func (hWnd HWND) SetWindowSubclass(
	subclassProc uintptr,
	idSubclass uint32,
	refData unsafe.Pointer,
) error {
	ret, _, _ := syscall.SyscallN(dll.Comctl(dll.PROC_SetWindowSubclass),
		uintptr(hWnd),
		subclassProc,
		uintptr(idSubclass),
		uintptr(refData))
	return utl.ZeroAsSysInvalidParm(ret)
}
