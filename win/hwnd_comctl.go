//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [DefSubclassProc] function.
//
// [DefSubclassProc]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-defsubclassproc
func (hWnd HWND) DefSubclassProc(
	msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.SyscallN(proc.DefSubclassProc.Addr(),
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return ret
}

// [RemoveWindowSubclass] function.
//
// [RemoveWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(
	subclassProc uintptr, idSubclass uint32) {

	ret, _, err := syscall.SyscallN(proc.RemoveWindowSubclass.Addr(),
		uintptr(hWnd), subclassProc, uintptr(idSubclass))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [SetWindowSubclass] function.
//
// [SetWindowSubclass]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func (hWnd HWND) SetWindowSubclass(
	subclassProc uintptr, idSubclass uint32, refData unsafe.Pointer) {

	ret, _, err := syscall.SyscallN(proc.SetWindowSubclass.Addr(),
		uintptr(hWnd), subclassProc, uintptr(idSubclass), uintptr(refData))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [TaskDialog] function.
//
// [TaskDialog]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialog
func (hWnd HWND) TaskDialog(
	hInstance HINSTANCE,
	windowTitle, mainInstruction, content StrOpt,
	commonButtons co.TDCBF,
	icon co.TD_ICON) co.ID {

	var pnButton int32
	ret, _, _ := syscall.SyscallN(proc.TaskDialog.Addr(),
		uintptr(hWnd), uintptr(hInstance),
		uintptr(windowTitle.Raw()),
		uintptr(mainInstruction.Raw()),
		uintptr(content.Raw()),
		uintptr(commonButtons), uintptr(icon),
		uintptr(unsafe.Pointer(&pnButton)))
	if wErr := errco.ERROR(ret); wErr != errco.S_OK {
		panic(wErr)
	}
	return co.ID(pnButton)
}
