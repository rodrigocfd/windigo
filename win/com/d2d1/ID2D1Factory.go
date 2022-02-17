package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1co"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1factory
type ID2D1Factory struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1Factory.Release().
func NewID2D1Factory(base com.IUnknown) ID2D1Factory {
	return ID2D1Factory{IUnknown: base}
}

// Creates a new ID2D1Factory.
//
// ⚠️ You must defer ID2D1Factory.Release().
func D2D1CreateFactory(
	factoryType d2d1co.FACTORY_TYPE,
	debugLevel d2d1co.DEBUG_LEVEL) ID2D1Factory {

	options := FACTORY_OPTIONS{
		DebugLevel: debugLevel,
	}

	var ppvQueried com.IUnknown
	ret, _, _ := syscall.Syscall6(proc.D2D1CreateFactory.Addr(), 4,
		uintptr(factoryType),
		uintptr(unsafe.Pointer(win.GuidFromIid(d2d1co.IID_ID2D1Factory))),
		uintptr(unsafe.Pointer(&options)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewID2D1Factory(ppvQueried)
	} else {
		panic(hr)
	}
}

// ⚠️ You must defer ID2D1HwndRenderTarget.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createhwndrendertarget(constd2d1_render_target_properties_constd2d1_hwnd_render_target_properties_id2d1hwndrendertarget)
func (me *ID2D1Factory) CreateHwndRenderTarget(
	targetProps *RENDER_TARGET_PROPERTIES,
	hwndTargetProps *HWND_RENDER_TARGET_PROPERTIES) ID2D1HwndRenderTarget {

	var ppvQueried com.IUnknown
	ret, _, _ := syscall.Syscall6(
		(*d2d1vt.ID2D1Factory)(unsafe.Pointer(*me.Ptr())).CreateHwndRenderTarget, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(targetProps)),
		uintptr(unsafe.Pointer(hwndTargetProps)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewID2D1HwndRenderTarget(ppvQueried)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-reloadsystemmetrics
func (me *ID2D1Factory) ReloadSystemMetrics() {
	ret, _, _ := syscall.Syscall(
		(*d2d1vt.ID2D1Factory)(unsafe.Pointer(*me.Ptr())).ReloadSystemMetrics, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
