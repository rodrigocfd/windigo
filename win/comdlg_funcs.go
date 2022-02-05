package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms646912(v=vs.85)
func ChooseColor(cc *CHOOSECOLOR) bool {
	ret, _, _ := syscall.Syscall(proc.ChooseColor.Addr(), 1,
		uintptr(unsafe.Pointer(cc)), 0, 0)
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-commdlgextendederror
func CommDlgExtendedError() errco.CDERR {
	ret, _, _ := syscall.Syscall(proc.CommDlgExtendedError.Addr(), 0,
		0, 0, 0)
	return errco.CDERR(ret)
}
