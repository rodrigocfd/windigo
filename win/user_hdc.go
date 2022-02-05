package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
func (hdc HDC) DrawIcon(x, y int32, hIcon HICON) {
	ret, _, err := syscall.Syscall6(proc.DrawIcon.Addr(), 4,
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(hIcon), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(
	rcClip *RECT,
	enumFunc func(hMon HMONITOR, hdcMon HDC, rcMon *RECT) bool) {

	pPack := &_EnumMonitorPack{f: enumFunc}
	if _globalEnumMonitorFuncs == nil {
		_globalEnumMonitorFuncs = make(map[*_EnumMonitorPack]struct{}, 2)
	}
	_globalEnumMonitorFuncs[pPack] = struct{}{} // store pointer in the set

	ret, _, err := syscall.Syscall6(proc.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(rcClip)),
		_globalEnumMonitorCallback, uintptr(unsafe.Pointer(pPack)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

type _EnumMonitorPack struct {
	f func(hMon HMONITOR, hdcMon HDC, rcMon *RECT) bool
}

var (
	_globalEnumMonitorCallback uintptr = syscall.NewCallback(_EnumMonitorProc)
	_globalEnumMonitorFuncs    map[*_EnumMonitorPack]struct{}
)

func _EnumMonitorProc(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lParam LPARAM) uintptr {
	pPack := (*_EnumMonitorPack)(unsafe.Pointer(lParam))
	retVal := uintptr(0)
	if _, isStored := _globalEnumMonitorFuncs[pPack]; isStored {
		retVal = util.BoolToUintptr(pPack.f(hMon, hdcMon, rcMon))
		if retVal == 0 {
			delete(_globalEnumMonitorFuncs, pPack) // remove from set
		}
	}
	return retVal
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-framerect
func (hdc HDC) FrameRect(rc *RECT, hBrush HBRUSH) {
	ret, _, err := syscall.Syscall(proc.FrameRect.Addr(), 3,
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(rc *RECT) {
	ret, _, err := syscall.Syscall(proc.InvertRect.Addr(), 2,
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
