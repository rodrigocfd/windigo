//go:build windows

package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1co"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [ID2D1RenderTarget] COM interface.
//
// [ID2D1RenderTarget]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1rendertarget
type ID2D1RenderTarget interface {
	ID2D1Resource

	// [BeginDraw] COM method.
	//
	// ⚠️ You must defer ID2D1RenderTarget.EndDraw().
	//
	// [BeginDraw]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-begindraw
	BeginDraw()

	// [Clear] COM method.
	//
	// [Clear]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-clear(constd2d1_color_f)
	Clear(clearColor *COLOR_F)

	// [CreateLayer] COM method.
	//
	// ⚠️ You must defer ID2D1Layer.Release().
	//
	// [CreateLayer]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createlayer(d2d1_size_f_id2d1layer)
	CreateLayer(size SIZE_F) ID2D1Layer

	// [CreateMesh] COM method.
	//
	// ⚠️ You must defer ID2D1Mesh.Release().
	//
	// [CreateMesh]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createmesh
	CreateMesh() ID2D1Mesh

	// [EndDraw] COM method.
	//
	// [EndDraw]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-enddraw
	EndDraw() (tag1, tag2 uint64)

	// [Flush] COM method.
	//
	// [Flush]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-flush
	Flush() (tag1, tag2 uint64)

	// [GetAntialiasMode] COM method.
	//
	// [GetAntialiasMode]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getantialiasmode
	GetAntialiasMode() d2d1co.ANTIALIAS_MODE

	// [GetDpi] COM method.
	//
	// [GetDpi]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getdpi
	GetDpi() (float32, float32)

	// [GetMaximumBitmapSize] COM method.
	//
	// [GetMaximumBitmapSize]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getmaximumbitmapsize
	GetMaximumBitmapSize() uint32

	// [GetPixelFormat] COM method.
	//
	// [GetPixelFormat]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getpixelformat
	GetPixelFormat() PIXEL_FORMAT

	// [GetPixelSize] COM method.
	//
	// [GetPixelSize]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getpixelsize
	GetPixelSize() SIZE_U

	// [GetSize] COM method.
	//
	// [GetSize]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getsize
	GetSize() SIZE_F

	// [IsSupported] COM method.
	//
	// [IsSupported]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-issupported(constd2d1_render_target_properties)
	IsSupported(renderTargetProperties *RENDER_TARGET_PROPERTIES) bool

	// [SetAntialiasMode] COM method.
	//
	// [SetAntialiasMode]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-setantialiasmode
	SetAntialiasMode(antialiasMode d2d1co.ANTIALIAS_MODE)

	// [SetDpi] COM method.
	//
	// [SetDpi]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-setdpi
	SetDpi(dpiX, dpiY float32)
}

type _ID2D1RenderTarget struct{ ID2D1Resource }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1RenderTarget.Release().
func NewID2D1RenderTarget(base com.IUnknown) ID2D1RenderTarget {
	return &_ID2D1RenderTarget{ID2D1Resource: NewID2D1Resource(base)}
}

func (me *_ID2D1RenderTarget) BeginDraw() {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).BeginDraw,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ID2D1RenderTarget) Clear(clearColor *COLOR_F) {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).Clear,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(clearColor)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_ID2D1RenderTarget) CreateLayer(size SIZE_F) ID2D1Layer {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).CreateLayer,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(util.Make64(uint32(size.Width), uint32(size.Height))),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewID2D1Layer(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_ID2D1RenderTarget) CreateMesh() ID2D1Mesh {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).CreateMesh,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewID2D1Mesh(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_ID2D1RenderTarget) EndDraw() (tag1, tag2 uint64) {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).EndDraw,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	} else {
		return
	}
}

func (me *_ID2D1RenderTarget) Flush() (tag1, tag2 uint64) {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).Flush,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	} else {
		return
	}
}

func (me *_ID2D1RenderTarget) GetAntialiasMode() d2d1co.ANTIALIAS_MODE {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).GetAntialiasMode,
		uintptr(unsafe.Pointer(me.Ptr())))

	return d2d1co.ANTIALIAS_MODE(ret)
}

func (me *_ID2D1RenderTarget) GetDpi() (float32, float32) {
	var dpiX, dpiY float32
	syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).GetDpi,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)))

	return dpiX, dpiY
}

func (me *_ID2D1RenderTarget) GetMaximumBitmapSize() uint32 {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).GetMaximumBitmapSize,
		uintptr(unsafe.Pointer(me.Ptr())))

	return uint32(ret)
}

func (me *_ID2D1RenderTarget) GetPixelFormat() PIXEL_FORMAT {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).GetPixelSize,
		uintptr(unsafe.Pointer(me.Ptr())))

	lo, hi := util.Break64(uint64(ret))
	return PIXEL_FORMAT{
		Format:    d2d1co.DXGI_FORMAT(lo),
		AlphaMode: d2d1co.ALPHA_MODE(hi),
	}
}

func (me *_ID2D1RenderTarget) GetPixelSize() SIZE_U {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).GetPixelSize,
		uintptr(unsafe.Pointer(me.Ptr())))

	lo, hi := util.Break64(uint64(ret))
	return SIZE_U{Width: lo, Height: hi}
}

func (me *_ID2D1RenderTarget) GetSize() SIZE_F {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).GetSize,
		uintptr(unsafe.Pointer(me.Ptr())))

	lo, hi := util.Break64(uint64(ret))
	return SIZE_F{Width: float32(lo), Height: float32(hi)}
}

func (me *_ID2D1RenderTarget) IsSupported(
	renderTargetProperties *RENDER_TARGET_PROPERTIES) bool {

	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).IsSupported,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(renderTargetProperties)))

	return ret != 0
}

func (me *_ID2D1RenderTarget) SetAntialiasMode(
	antialiasMode d2d1co.ANTIALIAS_MODE) {

	syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).SetAntialiasMode,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(antialiasMode))
}

func (me *_ID2D1RenderTarget) SetDpi(dpiX, dpiY float32) {
	syscall.SyscallN(
		(*d2d1vt.ID2D1RenderTarget)(unsafe.Pointer(*me.Ptr())).SetDpi,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(dpiX), uintptr(dpiY))
}
