package api

import (
	"syscall"
	"unsafe"
	"winffi/api/proc"
)

type HDC HANDLE

func (hdc HDC) EnumDisplayMonitors(rcClip *RECT) []HMONITOR {
	hMons := []HMONITOR{}
	syscall.Syscall6(proc.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(rcClip)),
		syscall.NewCallback(
			func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lp LPARAM) uintptr {
				hMons = append(hMons, hMon)
				return uintptr(1)
			}), 0, 0, 0)
	return hMons
}
