/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type OPENFILENAME struct {
	LStructSize       uint32
	HwndOwner         HWND
	HInstance         HINSTANCE
	LpstrFilter       uintptr // LPCWSTR
	LpstrCustomFilter uintptr // LPWSTR
	NMaxCustFilter    uint32
	NFilterIndex      uint32
	LpstrFile         uintptr // LPWSTR
	NMaxFile          uint32
	LpstrFileTitle    uintptr // LPWSTR
	NMaxFileTitle     uint32
	LpstrInitialDir   uintptr // LPCWSTR
	LpstrTitle        uintptr // LPCWSTR
	Flags             co.OFN
	NFileOffset       uint16
	NFileExtension    uint16
	LpstrDefExt       uintptr // LPCWSTR
	LCustData         LPARAM
	LpfnHook          uintptr // LPOFNHOOKPROC
	LpTemplateName    uintptr // LPCWSTR
	PvReserved        uintptr
	DwReserved        uint32
	FlagsEx           co.OFN_EX
}

func (ofn *OPENFILENAME) GetOpenFileName() bool {
	ofn.LStructSize = uint32(unsafe.Sizeof(*ofn)) // safety
	ret, _, _ := syscall.Syscall(proc.GetOpenFileName.Addr(), 1,
		uintptr(unsafe.Pointer(ofn)), 0, 0)

	if ret == 0 {
		ret, _, _ := syscall.Syscall(proc.CommDlgExtendedError.Addr(), 0,
			0, 0, 0)
		if ret != 0 {
			panic(fmt.Sprintf("GetOpenFileName failed: %d.", ret))
		} else {
			return false // user cancelled
		}
	}
	return true // user clicked OK
}

func (ofn *OPENFILENAME) GetSaveFileName() bool {
	ofn.LStructSize = uint32(unsafe.Sizeof(*ofn)) // safety
	ret, _, _ := syscall.Syscall(proc.GetSaveFileName.Addr(), 1,
		uintptr(unsafe.Pointer(ofn)), 0, 0)

	if ret == 0 {
		ret, _, _ := syscall.Syscall(proc.CommDlgExtendedError.Addr(), 0,
			0, 0, 0)
		if ret != 0 {
			panic(fmt.Sprintf("GetSaveFileName failed: %d.", ret))
		} else {
			return false // user cancelled
		}
	}
	return true // user clicked OK
}
