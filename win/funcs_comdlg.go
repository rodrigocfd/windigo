//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [ChooseColor] function.
//
// [ChooseColor]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms646912(v=vs.85)
func ChooseColor(cc *CHOOSECOLOR) bool {
	ret, _, _ := syscall.SyscallN(proc.ChooseColor.Addr(),
		uintptr(unsafe.Pointer(cc)))
	if ret == 0 {
		dlgErr := CommDlgExtendedError()
		if dlgErr == errco.CDERR_OK {
			return false
		} else {
			panic(dlgErr)
		}
	}
	return true
}

// [CommDlgExtendedError] function.
//
// [CommDlgExtendedError]: https://learn.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-commdlgextendederror
func CommDlgExtendedError() errco.CDERR {
	ret, _, _ := syscall.SyscallN(proc.CommDlgExtendedError.Addr())
	return errco.CDERR(ret)
}
