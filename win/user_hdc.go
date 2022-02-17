package win

import (
	"sync"
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
	callback func(hMon HMONITOR, hdcMon HDC, rcMon *RECT) bool) {

	pPack := &_EnumMonitorsPack{f: callback}
	_globalEnumMonitorsMutex.Lock()
	if _globalEnumMonitorsFuncs == nil { // the set was not initialized yet?
		_globalEnumMonitorsFuncs = make(map[*_EnumMonitorsPack]struct{}, 1)
	}
	_globalEnumMonitorsFuncs[pPack] = struct{}{} // store pointer in the set
	_globalEnumMonitorsMutex.Unlock()

	ret, _, err := syscall.Syscall6(proc.EnumDisplayMonitors.Addr(), 4,
		uintptr(hdc), uintptr(unsafe.Pointer(rcClip)),
		_globalEnumMonitorsCallback, uintptr(unsafe.Pointer(pPack)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

type _EnumMonitorsPack struct {
	f func(hMon HMONITOR, hdcMon HDC, rcMon *RECT) bool
}

var (
	_globalEnumMonitorsCallback uintptr = syscall.NewCallback(_EnumMonitorProc)
	_globalEnumMonitorsFuncs    map[*_EnumMonitorsPack]struct{}
	_globalEnumMonitorsMutex    = sync.Mutex{}
)

func _EnumMonitorProc(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lParam LPARAM) uintptr {
	pPack := (*_EnumMonitorsPack)(unsafe.Pointer(lParam))
	retVal := uintptr(0)

	_globalEnumMonitorsMutex.Lock()
	_, isStored := _globalEnumMonitorsFuncs[pPack]
	_globalEnumMonitorsMutex.Unlock()

	if isStored {
		retVal = util.BoolToUintptr(pPack.f(hMon, hdcMon, rcMon))
		if retVal == 0 {
			_globalEnumMonitorsMutex.Lock()
			delete(_globalEnumMonitorsFuncs, pPack) // remove from the set
			_globalEnumMonitorsMutex.Unlock()
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
