package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1hwndrendertarget
type ID2D1HwndRenderTarget interface {
	ID2D1RenderTarget

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-gethwnd
	GetHwnd() win.HWND

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-resize(constd2d1_size_u)
	Resize(pixelSize SIZE_U)
}

type _ID2D1HwndRenderTarget struct{ ID2D1RenderTarget }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1HwndRenderTarget.Release().
func NewID2D1HwndRenderTarget(base com.IUnknown) ID2D1HwndRenderTarget {
	return &_ID2D1HwndRenderTarget{ID2D1RenderTarget: NewID2D1RenderTarget(base)}
}

func (me *_ID2D1HwndRenderTarget) GetHwnd() win.HWND {
	ret, _, _ := syscall.Syscall(
		(*d2d1vt.ID2D1HwndRenderTarget)(unsafe.Pointer(*me.Ptr())).GetHwnd, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	return win.HWND(ret)
}

func (me *_ID2D1HwndRenderTarget) Resize(pixelSize SIZE_U) {
	ret, _, _ := syscall.Syscall(
		(*d2d1vt.ID2D1HwndRenderTarget)(unsafe.Pointer(*me.Ptr())).Resize, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pixelSize)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
