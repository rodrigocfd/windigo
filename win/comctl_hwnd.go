package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-defsubclassproc
func (hWnd HWND) DefSubclassProc(
	msg co.WM, wParam WPARAM, lParam LPARAM) uintptr {

	ret, _, _ := syscall.Syscall6(proc.DefSubclassProc.Addr(), 4,
		uintptr(hWnd), uintptr(msg), uintptr(wParam), uintptr(lParam),
		0, 0)
	return ret
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-removewindowsubclass
func (hWnd HWND) RemoveWindowSubclass(
	subclassProc uintptr, idSubclass uint32) {

	ret, _, err := syscall.Syscall(proc.RemoveWindowSubclass.Addr(), 3,
		uintptr(hWnd), subclassProc, uintptr(idSubclass))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func (hWnd HWND) SetWindowSubclass(
	subclassProc uintptr, idSubclass uint32, refData unsafe.Pointer) {

	ret, _, err := syscall.Syscall6(proc.SetWindowSubclass.Addr(), 4,
		uintptr(hWnd), subclassProc, uintptr(idSubclass), uintptr(refData),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-taskdialog
func (hWnd HWND) TaskDialog(
	hInstance HINSTANCE,
	windowTitle, mainInstruction, content StrOpt,
	commonButtons co.TDCBF, icon co.TD_ICON) co.ID {

	var pnButton int32
	ret, _, _ := syscall.Syscall9(proc.TaskDialog.Addr(), 8,
		uintptr(hWnd), uintptr(hInstance),
		uintptr(windowTitle.raw()),
		uintptr(mainInstruction.raw()),
		uintptr(content.raw()),
		uintptr(commonButtons), uintptr(icon),
		uintptr(unsafe.Pointer(&pnButton)), 0)
	if wErr := errco.ERROR(ret); wErr != errco.S_OK {
		panic(wErr)
	}
	return co.ID(pnButton)
}
