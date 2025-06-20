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
	ret, _, _ := syscall.SyscallN(
		dll.Comctl(&_DefSubclassProc, "DefSubclassProc"),
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

var _DefSubclassProc *syscall.Proc

// [ImageList_DragEnter] function.
//
// [ImageList_DragEnter]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragenter
func (hWnd HWND) ImageListDragEnter(x, y int) error {
	ret, _, _ := syscall.SyscallN(
		dll.Comctl(&_ImageList_DragEnter, "ImageList_DragEnter"),
		uintptr(hWnd),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_DragEnter *syscall.Proc

// [ImageList_DragLeave] function.
//
// [ImageList_DragLeave]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragleave
func (hWnd HWND) ImageListDragLeave() error {
	ret, _, _ := syscall.SyscallN(
		dll.Comctl(&_ImageList_DragLeave, "ImageList_DragLeave"),
		uintptr(hWnd))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_DragLeave *syscall.Proc

// [RemoveWindowSubclass] function.
//
// [RemoveWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(subclassProc uintptr, idSubclass uint32) error {
	ret, _, _ := syscall.SyscallN(
		dll.Comctl(&_RemoveWindowSubclass, "RemoveWindowSubclass"),
		uintptr(hWnd),
		subclassProc,
		uintptr(idSubclass))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _RemoveWindowSubclass *syscall.Proc

// [SetWindowSubclass] function.
//
// [SetWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func (hWnd HWND) SetWindowSubclass(
	subclassProc uintptr,
	idSubclass uint32,
	refData unsafe.Pointer,
) error {
	ret, _, _ := syscall.SyscallN(
		dll.Comctl(&_SetWindowSubclass, "SetWindowSubclass"),
		uintptr(hWnd),
		subclassProc,
		uintptr(idSubclass),
		uintptr(refData))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _SetWindowSubclass *syscall.Proc
