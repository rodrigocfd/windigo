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
		uintptr(hWnd),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret
}

var _DefSubclassProc = dll.Comctl32.NewProc("DefSubclassProc")

// [ImageList_DragEnter] function.
//
// [ImageList_DragEnter]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragenter
func (hWnd HWND) ImageListDragEnter(x, y int) error {
	ret, _, _ := syscall.SyscallN(_ImageList_DragEnter.Addr(),
		uintptr(hWnd),
		uintptr(int32(x)),
		uintptr(int32(y)))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_DragEnter = dll.Comctl32.NewProc("ImageList_DragEnter")

// [ImageList_DragLeave] function.
//
// [ImageList_DragLeave]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_dragleave
func (hWnd HWND) ImageListDragLeave() error {
	ret, _, _ := syscall.SyscallN(_ImageList_DragLeave.Addr(),
		uintptr(hWnd))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _ImageList_DragLeave = dll.Comctl32.NewProc("ImageList_DragLeave")

// [RemoveWindowSubclass] function.
//
// [RemoveWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(subclassProc uintptr, idSubclass uint32) error {
	ret, _, _ := syscall.SyscallN(_RemoveWindowSubclass.Addr(),
		uintptr(hWnd),
		subclassProc,
		uintptr(idSubclass))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _RemoveWindowSubclass = dll.Comctl32.NewProc("RemoveWindowSubclass")

// [SetWindowSubclass] function.
//
// [SetWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func (hWnd HWND) SetWindowSubclass(
	subclassProc uintptr,
	idSubclass uint32,
	refData unsafe.Pointer,
) error {
	ret, _, _ := syscall.SyscallN(_SetWindowSubclass.Addr(),
		uintptr(hWnd),
		subclassProc,
		uintptr(idSubclass),
		uintptr(refData))
	return utl.ZeroAsSysInvalidParm(ret)
}

var _SetWindowSubclass = dll.Comctl32.NewProc("SetWindowSubclass")
