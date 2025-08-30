//go:build windows

package win

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [RegisterDragDrop] function.
//
// Paired with [HWND.RevokeDragDrop].
//
// [RegisterDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-registerdragdrop
func (hWnd HWND) RegisterDragDrop(dropTarget *IDropTarget) error {
	exStyle, _ := hWnd.ExStyle()
	if (exStyle & co.WS_EX_ACCEPTFILES) != 0 {
		return errors.New("do not use WS_EX_ACCEPTFILES with RegisterDragDrop")
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_RegisterDragDrop, "RegisterDragDrop"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(dropTarget.Ppvt())))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		if hr == co.HRESULT_E_OUTOFMEMORY {
			return errors.New("RegisterDragDrop failed, did you call OleInitialize?")
		}
		return hr
	}
	return nil
}

var _RegisterDragDrop *syscall.Proc

// [RevokeDragDrop] function.
//
// Paired with [HWND.RegisterDragDrop].
//
// [RevokeDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-revokedragdrop
func (hWnd HWND) RevokeDragDrop() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_RevokeDragDrop, "RevokeDragDrop"),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

var _RevokeDragDrop *syscall.Proc
