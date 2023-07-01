//go:build windows

package win

import (
	"sync"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [DrawIcon] function.
//
// [DrawIcon]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawicon
func (hdc HDC) DrawIcon(x, y int32, hIcon HICON) {
	ret, _, err := syscall.SyscallN(proc.DrawIcon.Addr(),
		uintptr(hdc), uintptr(x), uintptr(y), uintptr(hIcon))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [DrawIconEx] function.
//
// [DrawIconEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-drawiconex
func (hdc HDC) DrawIconEx(
	pos POINT, hIcon HICON, size SIZE, frameIndex uint32,
	hbrFlickerFree HBRUSH, diFlags co.DI) {

	ret, _, err := syscall.SyscallN(proc.DrawIconEx.Addr(),
		uintptr(hdc), uintptr(pos.X), uintptr(pos.Y), uintptr(hIcon),
		uintptr(size.Cx), uintptr(size.Cy), uintptr(frameIndex),
		uintptr(hbrFlickerFree), uintptr(diFlags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [EnumDisplayMonitors] function.
//
// To continue enumeration, the callback function must return true; to stop
// enumeration, it must return false.
//
// [EnumDisplayMonitors]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func (hdc HDC) EnumDisplayMonitors(
	rcClip *RECT,
	callback func(hMon HMONITOR, hdcMon HDC, rcMon *RECT) bool) error {

	pPack := &_EnumMonitorsPack{f: callback}
	_globalEnumMonitorsMutex.Lock()
	if _globalEnumMonitorsFuncs == nil { // the set was not initialized yet?
		_globalEnumMonitorsFuncs = make(map[*_EnumMonitorsPack]struct{}, 1)
	}
	_globalEnumMonitorsFuncs[pPack] = struct{}{} // store pointer in the set
	_globalEnumMonitorsMutex.Unlock()

	ret, _, err := syscall.SyscallN(proc.EnumDisplayMonitors.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(rcClip)),
		_globalEnumMonitorsCallback, uintptr(unsafe.Pointer(pPack)))

	_globalEnumMonitorsMutex.Lock()
	delete(_globalEnumMonitorsFuncs, pPack) // remove from the set
	_globalEnumMonitorsMutex.Unlock()

	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

type _EnumMonitorsPack struct {
	f func(hMon HMONITOR, hdcMon HDC, rcMon *RECT) bool
}

var (
	_globalEnumMonitorsFuncs    map[*_EnumMonitorsPack]struct{} // keeps pointers from being collected by GC
	_globalEnumMonitorsMutex    = sync.Mutex{}
	_globalEnumMonitorsCallback = syscall.NewCallback(
		func(hMon HMONITOR, hdcMon HDC, rcMon *RECT, lParam LPARAM) uintptr {
			pPack := (*_EnumMonitorsPack)(unsafe.Pointer(lParam))
			return util.BoolToUintptr(pPack.f(hMon, hdcMon, rcMon))
		})
)

// [FrameRect] function.
//
// [FrameRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-framerect
func (hdc HDC) FrameRect(rc *RECT, hBrush HBRUSH) {
	ret, _, err := syscall.SyscallN(proc.FrameRect.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(rc)), uintptr(hBrush))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [InvertRect] function.
//
// [InvertRect]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-invertrect
func (hdc HDC) InvertRect(rc *RECT) {
	ret, _, err := syscall.SyscallN(proc.InvertRect.Addr(),
		uintptr(hdc), uintptr(unsafe.Pointer(rc)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// [PaintDesktop] function.
//
// [PaintDesktop]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-paintdesktop
func (hdc HDC) PaintDesktop() {
	ret, _, err := syscall.SyscallN(proc.PaintDesktop.Addr(),
		uintptr(hdc))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}
